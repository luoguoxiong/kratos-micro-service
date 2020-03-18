package dao

import (
	"context"
	"time"

	client "kratosmicoservice/pkg"

	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/bilibili/kratos/pkg/time"

	"github.com/google/wire"
)

// Provider ..
var Provider = wire.NewSet(New, NewDB, NewRedis, NewMC)

// Dao dao.
type Dao struct {
	db           *sql.DB
	redis        *redis.Redis
	mc           *memcache.Memcache
	goodsService interface{}
	cache        *fanout.Fanout
	demoExpire   int32
}

// New new a dao and return.
func New(r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(r, mc, db)
}
func newDao(r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d *Dao, cf func(), err error) {
	var cfg struct {
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	var (
		service struct {
			Servers map[string]*client.ServerConfig
		}
	)
	_ = paladin.Get("dao.toml").UnmarshalTOML(&service)
	d = &Dao{
		db:           db,
		redis:        r,
		mc:           mc,
		goodsService: client.NewCommonServiceClient(client.NewServerConf(service.Servers["goods-service"])),
		cache:        fanout.New("cache"),
		demoExpire:   int32(time.Duration(cfg.DemoExpire) / time.Second),
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return nil
}

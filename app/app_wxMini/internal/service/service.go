package service

import (
	"context"

	pb "kratosmicoservice/app/app_wxMini/api"
	"kratosmicoservice/app/app_wxMini/internal/dao"
	goodsService "kratosmicoservice/service/service_goods/api"

	"github.com/bilibili/kratos/pkg/conf/paladin"

	"github.com/google/wire"
)

// Provider ...
var Provider = wire.NewSet(New, wire.Bind(new(pb.WxMiniServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// GoodsDetail bm
func (s *Service) GoodsDetail(ctx context.Context, req *goodsService.GoodsReq) (reply *goodsService.GoodsRes, err error) {
	s.dao.
		reply = &goodsService.GoodsRes{
		GoodsName: "hello " + req.Id,
	}
	return
}

// Close close the resource.
func (s *Service) Close() {
}

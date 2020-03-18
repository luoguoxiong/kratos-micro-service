package pkg

import (
	"context"
	goods_service_v1 "kratosmicoservice/service/service_goods/api"

	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	xtime "github.com/bilibili/kratos/pkg/time"
)

// GRPCConf ...
type GRPCConf struct {
	Server                *warden.ClientConfig
	Addr                  string
	MaxReceiveMessageSize int
}

// ServerConfig ss
type ServerConfig struct {
	Addr                  string
	Timeout               xtime.Duration
	MaxReceiveMessageSize int
}

// NewCommonServiceClient ...
func NewCommonServiceClient(cfg *GRPCConf) goods_service_v1.GoodsClient {
	cc, err := warden.NewClient(cfg.Server).Dial(context.Background(), cfg.Addr)
	if err != nil {
		panic(err)
	}
	return goods_service_v1.NewGoodsClient(cc)
}

// NewServerConf ...
func NewServerConf(config *ServerConfig) *GRPCConf {
	wConf := &warden.ClientConfig{
		Timeout: config.Timeout,
	}
	return &GRPCConf{
		Addr:                  config.Addr,
		Server:                wConf,
		MaxReceiveMessageSize: config.MaxReceiveMessageSize,
	}

}

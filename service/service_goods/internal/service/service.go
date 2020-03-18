package service

import (
	"context"
	"fmt"

	pb "kratosmicoservice/service/service_goods/api"
	"kratosmicoservice/service/service_goods/internal/dao"

	"github.com/bilibili/kratos/pkg/conf/paladin"

	"github.com/google/wire"
)

// Provider ...
var Provider = wire.NewSet(New, wire.Bind(new(pb.GoodsServer), new(*Service)))

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

// GoodsDetail 获取商品详情
func (s *Service) GoodsDetail(ctx context.Context, req *pb.GoodsReq) (reply *pb.GoodsRes, err error) {
	reply = &pb.GoodsRes{
		GoodsName: `查找商品id为` + req.Id + "的详情！",
	}
	return
}

// Close 服务关闭后释放资源？？
func (s *Service) Close() {
	fmt.Println("服务已关闭！")
}

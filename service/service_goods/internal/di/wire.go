// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"kratosmicoservice/service/service_goods/internal/dao"
	"kratosmicoservice/service/service_goods/internal/server/grpc"
	"kratosmicoservice/service/service_goods/internal/server/http"
	"kratosmicoservice/service/service_goods/internal/service"

	"github.com/google/wire"
)

// InitApp go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}

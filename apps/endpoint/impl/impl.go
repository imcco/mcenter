package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

type impl struct {
	col *mongo.Collection
	log *zerolog.Logger
	endpoint.UnimplementedRPCServer
	ioc.ObjectImpl

	svc service.MetaService
}

func (i *impl) Init() error {
	i.col = ioc_mongo.DB().Collection(i.Name())
	i.log = log.Sub(i.Name())

	i.svc = ioc.Controller().Get(service.AppName).(service.MetaService)
	return nil
}

func (i *impl) Name() string {
	return endpoint.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	endpoint.RegisterRPCServer(server, i)
}

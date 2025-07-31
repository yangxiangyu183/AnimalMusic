package svc

import (
	"Music_api/internal/config"
	"Music_api/internal/middleware"
	"Music_service/musicserviceclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	AuthTemplate rest.Middleware
	MusicService musicserviceclient.MusicService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		AuthTemplate: middleware.NewAuthTemplateMiddleware().Handle,
		MusicService: musicserviceclient.NewMusicService(zrpc.MustNewClient(c.MusicService)),
	}
}

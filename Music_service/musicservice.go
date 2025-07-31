package main

import (
	"Music_api/appconfig"
	"Music_service/inits"
	"flag"
	"fmt"

	"Music_service/internal/config"
	"Music_service/internal/server"
	"Music_service/internal/svc"
	"Music_service/music_service"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/musicservice.yaml", "the config file")

func main() {
	flag.Parse()
	appconfig.GetViperData()
	inits.InitAll()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		music_service.RegisterMusicServiceServer(grpcServer, server.NewMusicServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

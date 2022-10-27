package app

import (
	"context"
	"github.com/Skijetler/alphinium/auth/internal/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type App struct {
	ctx        context.Context
	conf       *config.Config
	log        *logrus.Logger
	grpcServer *grpc.Server
	httpServer *http.ServeMux
}

func New(conf *config.Config, logger *logrus.Logger, grpcSrv *grpc.Server, httpSrv *http.ServeMux) *App {
	return &App{
		ctx:        context.Background(),
		conf:       conf,
		log:        logger,
		grpcServer: grpcSrv,
		httpServer: httpSrv,
	}
}

func (a *App) Run() error {
	eg, ctx := errgroup.WithContext(a.ctx)
	wg := sync.WaitGroup{}

	// gRPC Server start
	eg.Go(func() error {
		<-ctx.Done() // wait for stop signal
		stopCtx, cancel := context.WithTimeout(a.ctx, 10*time.Second)
		defer cancel()
		return a.grpcServer.GracefulStop()
	})
	wg.Add(1)
	eg.Go(func() error {
		wg.Done() // here is to ensure server start has begun running before register, so defer is not needed
		listen, err := net.Listen("tcp", a.conf.Server.Grpc.Addr)
		if err != nil {
			log.Panic(err)
		}
		return a.grpcServer.Serve()
	})
}

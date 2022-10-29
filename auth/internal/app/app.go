package app

import (
	"context"
	"errors"
	"github.com/Skijetler/alphinium/auth/internal/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type App struct {
	ctx        context.Context
	cancel     func()
	conf       *config.Config
	log        *logrus.Logger
	grpcServer *grpc.Server
	httpServer *http.Server
}

func New(conf *config.Config, logger *logrus.Logger, grpcSrv *grpc.Server, httpSrv *http.Server) *App {
	ctx, cancel := context.WithCancel(context.Background())

	return &App{
		ctx:        ctx,
		cancel:     cancel,
		conf:       conf,
		log:        logger,
		grpcServer: grpcSrv,
		httpServer: httpSrv,
	}
}

func (a *App) Run() error {
	eg, ctx := errgroup.WithContext(a.ctx)
	wg := sync.WaitGroup{}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// gRPC Server start
	eg.Go(func() error {
		<-ctx.Done() // wait for stop signal
		_, cancel := context.WithTimeout(a.ctx, 10*time.Second)
		defer cancel()
		a.log.Info("[Auth] gRPC server stopping")
		a.grpcServer.GracefulStop()
		return nil
	})
	wg.Add(1)
	eg.Go(func() error {
		defer wg.Done()
		listen, err := net.Listen("tcp", a.conf.Server.Grpc.Addr)
		if err != nil {
			a.log.Fatal(err)
			return err
		}
		a.log.Infof("[Auth] gRPC server listening on: %s", a.conf.Server.Grpc.Addr)
		return a.grpcServer.Serve(listen)
	})

	// HTTP Server start
	eg.Go(func() error {
		<-ctx.Done() // wait for stop signal
		stopCtx, cancel := context.WithTimeout(a.ctx, 10*time.Second)
		defer cancel()
		a.log.Info("[Auth] HTTP server stopping")
		return a.httpServer.Shutdown(stopCtx)
	})
	wg.Add(1)
	eg.Go(func() error {
		defer wg.Done()
		a.log.Infof("[Auth] HTTP server listening on: %s", a.conf.Server.Http.Addr)
		return a.httpServer.ListenAndServe()
	})

	wg.Wait()

	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		case <-done:
			return a.Stop()
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (a *App) Stop() error {
	a.log.Info("[Auth] stopping service")
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

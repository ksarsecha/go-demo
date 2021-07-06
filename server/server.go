package server

import (
	"context"
	httpCfg "github.com/ksarsecha/movie_rental/config"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	Start()
}

type appServer struct {
	cfg         httpCfg.HTTPServerConfig
	lgr         *zap.Logger
	router      http.Handler
	tracerFlush func()
}

func (as *appServer) Start() {
	server := newHTTPServer(as.cfg, as.router)

	as.lgr.Sugar().Infof("listening on %s", as.cfg.GetAddress())
	go func() { _ = server.ListenAndServe() }()

	waitForShutdown(server, as.lgr, as.tracerFlush)
}

func waitForShutdown(server *http.Server, lgr *zap.Logger, tracerFlush func()) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigCh

	defer func() { _ = lgr.Sync() }()

	err := server.Shutdown(context.Background())
	if err != nil {
		lgr.Error(err.Error())
		return
	}

	tracerFlush()
	lgr.Info("server shutdown successful")
}

func newHTTPServer(cfg httpCfg.HTTPServerConfig, handler http.Handler) *http.Server {
	return &http.Server{
		Handler:      handler,
		Addr:         cfg.GetAddress(),
		WriteTimeout: time.Second * time.Duration(cfg.GetReadTimeout()),
		ReadTimeout:  time.Second * time.Duration(cfg.GetWriteTimeout()),
	}
}

func NewServer(cfg httpCfg.HTTPServerConfig, lgr *zap.Logger, router http.Handler, tracerFlush func()) Server {
	return &appServer{
		cfg:         cfg,
		lgr:         lgr,
		router:      router,
		tracerFlush: tracerFlush,
	}
}

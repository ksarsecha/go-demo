package main

import (
	"github.com/ksarsecha/movie_rental/config"
	http2 "github.com/ksarsecha/movie_rental/http"
	server2 "github.com/ksarsecha/movie_rental/server"
	"go.uber.org/zap"
)

func main() {
	var logger = zap.NewExample()

	httpServerConfig := config.InitConfig()
	server := server2.NewServer(httpServerConfig, logger, http2.Router(), nil)
	server.Start()
}

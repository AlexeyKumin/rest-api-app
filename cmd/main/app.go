package main

import (
	"first-rest/internal/config"
	"first-rest/internal/user"
	"first-rest/pkg/logging"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, listenErr := filepath.Abs(filepath.Dir(os.Args[0]))
		if listenErr != nil {
			logger.Fatal(listenErr)
		}
		logger.Info("create socket")
		socketPath := appDir + "\\app.sock"

		logger.Info("listen unix socket")
		net.Listen("unix", socketPath)

		listener, listenErr = net.Listen("unix", socketPath)
		logger.Info("server is listening unix socket: %s", socketPath)

	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Infof("server is listening %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	}
	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}

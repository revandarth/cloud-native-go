package main

import (
	"fmt"
	"net/http"
	"xks-go/config"
	"xks-go/http/router"
	app "xks-go/http/server"
	"xks-go/utils/logger"
)

func main() {

	appConf := config.AppConfig()

	logger := logger.NewConsole(appConf.Server.Debug)

	// router := http.NewServeMux()
	// router.HandleFunc("/", Hello)
	appLog := app.New(logger)
	apiRouter := router.New(appLog)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	// log.Println("Starting server: 8080")

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      apiRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// log.Fatal("Server failed to start")
		logger.Fatal().Err(err).Msg("Server failed start")
	}
}

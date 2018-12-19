package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"stroom/config"
	router "stroom/router"
	"time"

	log "github.com/Sirupsen/logrus"
)

var server *http.Server

func main() {
	//This handles graceful shutdown
	//https://gist.github.com/peterhellberg/38117e546c217960747aacf689af3dc2
	//setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	startServer()

	//sample usage of config
	log.Infof("db used is %s", config.GlobalConfig.DB.DatabaseName)
	<-stop

	shutdownServer()
}

func startServer() {
	routerCreated := router.Init()
	server = &http.Server{Addr: ":8080", Handler: routerCreated}
	http.Handle("/", routerCreated)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	log.Info("\nServer running. Waiting for shutdown signal...")
}

func shutdownServer() {
	log.Info("\nGot shutdown signal. Server shutting down")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(ctx)
	log.Info("Server gracefully stopped")
}

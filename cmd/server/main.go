package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	grpcProvider "watermap/cmd/server/providers/grpc"
	httpProvider "watermap/cmd/server/providers/http"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file not found, using OS env vars")
	}

	viper.AutomaticEnv()

	fmt.Println("Starting server... grpc_port=", viper.Get("grpc_port"), " http_port=", viper.Get("http_port"))
	go httpProvider.Provider().Serve()
	go grpcProvider.Provider().Serve()
	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)
	<-shutdown
}

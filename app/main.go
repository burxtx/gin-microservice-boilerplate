package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/burxtx/gin-microservice-boilerplate/app/config"
	"github.com/burxtx/gin-microservice-boilerplate/app/db"
	"github.com/burxtx/gin-microservice-boilerplate/app/endpoints"
	"github.com/burxtx/gin-microservice-boilerplate/app/server"
	"github.com/burxtx/gin-microservice-boilerplate/app/service"
)

func main() {
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*enviroment)
	c := config.GetConfig()
	datasource := c.GetString("database.host")
	dbSession, err := db.Init(datasource)
	if err != nil {
		panic(fmt.Errorf("Fatal error database connection: %s \n", err))
	}
	svc := service.New(dbSession)
	eps := endpoints.New(svc)
	server.Init(eps)
}

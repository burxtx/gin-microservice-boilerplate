package server

import (
	"fmt"

	"github.com/burxtx/gin-microservice-boilerplate/app/config"
	"github.com/burxtx/gin-microservice-boilerplate/app/endpoints"
)

func Init(eps endpoints.AppEndpoint) {
	config := config.GetConfig()
	r := NewRouter(eps)
	fmt.Println("=== router ready ===")
	r.Run(config.GetString("server.port"))
}

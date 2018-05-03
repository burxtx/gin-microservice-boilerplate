package transport

import (
	"fmt"

	"github.com/burxtx/gin-microservice-boilerplate/app/config"
	"github.com/burxtx/gin-microservice-boilerplate/app/endpoints"
)

func NewHttpHandler(eps endpoints.AppEndpoint) {
	config := config.GetConfig()
	r := NewHttpRouter(eps)
	fmt.Println("=== router ready ===")
	r.Run(config.GetString("server.port"))
}

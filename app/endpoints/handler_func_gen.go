package endpoints

import (
	"fmt"

	"github.com/burxtx/gin-microservice-boilerplate/app/service"
	"github.com/gin-gonic/gin"
)

type AppEndpoint struct {
	GetEndpoint gin.HandlerFunc
}

func New(s service.AppService) AppEndpoint {
	eps := AppEndpoint{
		GetEndpoint: MakeGetEndpoint(s),
	}
	fmt.Println("=== endpoints ready ===")
	return eps
}

package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/burxtx/gin-microservice-boilerplate/app/models"
	"github.com/burxtx/gin-microservice-boilerplate/app/service"
	"github.com/gin-gonic/gin"
)

type GetRequest struct {
	Id string `form:"id" json:"id" binding:"required"`
}

type GetResponse struct {
	Data models.App `json:"data"`
	Err  error      `json:"err"`
}

func MakeGetEndpoint(s service.AppService) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("--- gen get handler ---")
		var getReq GetRequest
		if err := c.ShouldBindWith(&getReq, binding.Query); err == nil {

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		app, getErr := s.Get(c, getReq.Id)
		if getErr != nil {
			http.Error(c.Writer, getErr.Error(), http.StatusInternalServerError)
			return
		}
		response := GetResponse{
			Data: app,
			Err:  nil,
		}
		c.JSON(200, response)
	}
}

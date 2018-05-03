package netrpc

import (
	"net/rpc"

	"github.com/burxtx/gin-microservice-boilerplate/app/models"
	"github.com/burxtx/gin-microservice-boilerplate/app/service"
)

func New(cli *rpc.Client) service.AppService {
	return client{cli}
}

type client struct {
	*rpc.Client
}

func (c client) Get(id string) (models.App, error) {
	var reply service.GetResponse
	if err := c.Client.Call("appservice.Get", service.GetRequest{Id: "1"}, &reply); err != nil {
		return models.App{}, err
	}
	return reply.Data, nil
}

/* func (c *client) Find(cond builder.Cond, orderCols ...string) ([]models.App, error) {
	var reply service.FindResponse
	err := a.db.Where(cond).Desc(orderCols...).Find(&all)
	if err != nil {
		return all, err
	}
	return all, nil
}

func (c *client) Add(audit models.App) (int64, error) {
	var reply service.AddResponse
	if err == nil {
		return affected, err
	}
	return affected, nil
}
*/

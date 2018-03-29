package service

import (
	"fmt"

	"github.com/burxtx/gin-microservice-boilerplate/app/models"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
)

type AppService interface {
	Find(ctx *gin.Context, cond builder.Cond, orderCols ...string) ([]models.App, error)
	Get(ctx *gin.Context, id string) (models.App, error)
	Add(ctx *gin.Context, audit models.App) (int64, error)
}

type appService struct{ db *xorm.Engine }

func (a *appService) Find(c *gin.Context, cond builder.Cond, orderCols ...string) ([]models.App, error) {
	all := make([]models.App, 0)
	err := a.db.Where(cond).Desc(orderCols...).Find(&all)
	if err != nil {
		return all, err
	}
	return all, nil
}

func (a *appService) Get(c *gin.Context, id string) (models.App, error) {
	audit := new(models.App)
	has, err := a.db.Where("id=?", id).Get(audit)
	if err != nil {
		return models.App{}, err
	}
	if has != true {
		return models.App{}, nil
	}
	return *audit, nil
}

func (a *appService) Add(c *gin.Context, audit models.App) (int64, error) {
	affected, err := a.db.Insert(audit)
	if err == nil {
		return affected, err
	}
	return affected, nil
}

func NewAppService(db *xorm.Engine) AppService {
	return &appService{db}
}

func New(db *xorm.Engine) AppService {
	var svc AppService = NewAppService(db)
	fmt.Println("=== service ready ===")
	return svc
}

package route

import (
	"eztrust/api/controller"
	"eztrust/bootstrap"
	"eztrust/infra/queue"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTunnelRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	tunnelController := controller.TunnelController{
		Rabbitmq: &queue.Rabbitmq{},
	}
	group.POST("/tunnel", tunnelController.GetTunnelInfo)
}

package route

import (
	"eztrust/api/controller"
	"eztrust/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	userController := controller.UserController{
		Database: db,
	}
	group.POST("/user/create", userController.CreateUser)
}

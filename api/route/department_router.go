package route

import (
	"eztrust/api/controller"
	"eztrust/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewDepartmentRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	departmentController := controller.DepartmentController{
		Database: db,
	}
	group.POST("/department/create", departmentController.CreateDepartment)
}

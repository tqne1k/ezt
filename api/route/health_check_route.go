package route

import (
	"eztrust/bootstrap"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewHealthCheckRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	group.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}

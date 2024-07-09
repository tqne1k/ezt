package route

import (
	"time"

	"eztrust/bootstrap"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "type_login"}
	gin.Use(cors.New(config))
	gin.LoadHTMLGlob("template/templates/*")
	gin.Static("/assets", "./template/assets")
	publicRouter := gin.Group("")
	// Public health check
	NewHealthCheckRouter(env, timeout, db, publicRouter)

	// Public user
	NewUserRouter(env, timeout, db, publicRouter)
}

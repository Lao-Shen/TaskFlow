package router

import (
	"backend/internal/auth"
	"backend/internal/middleware"
	"backend/internal/task"
	"backend/internal/user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func ServerInit() *gin.Engine {
	r = gin.Default()
	r.Use(gin.Recovery())

	// CORS 跨域中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routerInit()

	return r
}

func routerInit() {
	// 健康检查
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "你好"})
	})

	api := r.Group("/api")
	{
		// ---- 公开接口 ----
		api.POST("/register", auth.Register)
		api.POST("/login", auth.Login)

		// ---- 需鉴权接口 ----
		authGroup := api.Group("")
		authGroup.Use(middleware.AuthRequired())
		{
			authGroup.POST("/logout", auth.Logout)
			authGroup.GET("/user/me", user.GetMe)
			authGroup.POST("/tasks", task.CreateTask)
			authGroup.GET("/tasks", task.ListTasks)
		}
	}
}

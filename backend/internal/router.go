package internal

import "github.com/gin-gonic/gin"

var r *gin.Engine

func ServerInit() *gin.Engine {
	r = gin.Default()
	r.Use(gin.Recovery())

	routerInit()

	return r
}

func routerInit() {
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "你好",
		})
	})
}

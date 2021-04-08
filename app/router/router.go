package router

import (
	"goAPISample/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 初期化
func Init() {
	r := router()
	r.Run(":8080")
}

// ルーティング
func router() *gin.Engine {
	r := gin.Default()

	// CORS対応
	r.Use(CORS())

	// ルーティング
	u := r.Group("")
	{
		pc := controller.PostController{}
		u.GET("/posts", pc.Index)
		u.GET("/posts/:id", pc.Show)
		u.POST("/posts", pc.Create)
		u.PUT("/posts/:id", pc.Update)
		u.DELETE("/posts/:id", pc.Delete)
	}

	return r
}

// CORS
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

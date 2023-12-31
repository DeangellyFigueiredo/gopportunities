package router

import (
	"github.com/DeangellyFigueiredo/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context) {
			handler.ShowOpeningHandler(ctx)
		})

		v1.POST("/opening", func(ctx *gin.Context) {
			handler.CreateOpeningHandler(ctx)
		})

		v1.DELETE("/opening", func(ctx *gin.Context) {
			handler.DeleteOpeningHandler(ctx)

		})

		v1.PUT("/opening", func(ctx *gin.Context) {
			handler.UpdateOpeningHandler(ctx)

		})

		v1.GET("/openings", func(ctx *gin.Context) {
			handler.ListAllOpeningHandler(ctx)

		})
	}
}

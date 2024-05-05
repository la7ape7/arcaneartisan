package apiserver

import (
	"github.com/gin-gonic/gin"
	"watIwant/controllers"
	"watIwant/middlewares"
)

func V1Routes(engine *gin.Engine) {
	authMiddleware := middlewares.JwtMiddlewareHandler()

	engine.POST("/auth", authMiddleware.LoginHandler)
	engine.POST("/register", controllers.NewAccountController().Register)

	apiV1 := engine.Group("/v1")
	apiV1.Use(authMiddleware.MiddlewareFunc())
	{
		// ITEMS ROUTES
		items := apiV1.Group("/item")
		{
			items.GET("", controllers.NewItemController().Get)
			items.POST("", controllers.NewItemController().Post)

			item := items.Group(":item_id")
			{
				item.GET("", controllers.NewItemController().GetById)
			}

		}
	}

}

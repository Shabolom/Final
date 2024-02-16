package routes

import (
	"Graduation_Project/iternal/api"
	"Graduation_Project/iternal/middleware"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())

	user := api.NewUserApi()

	authRequired := r.Group("/")
	authRequired.Use(middleware.Logger())
	authRequired.Use(middleware.Authorization())

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/api/user/register", user.Regis)
	r.POST("/api/user/login", user.Login)

	{
		authRequired.POST("/api/user/orders", user.PostOrder)
		authRequired.GET("/api/user/orders", user.GetOrder)
		authRequired.GET("/api/user/balance", user.GetBal)
		authRequired.POST("/api/user/balance/withdraw", user.PostWithD)
		authRequired.GET("/api/user/withdrawals", user.GetWithD)
	}
	return r
}

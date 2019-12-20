package Routes

import "C"
import (
	"github.com/gin-gonic/gin"
	"jwtapp/Controllers"
	"jwtapp/Middlewares"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(Middlewares.CorsHandler())

	UserRouter := r.Group("/user").Use(Middlewares.JWTAuth())
	{
		UserRouter.POST("changePassword", Controllers.ChangePassword)
		UserRouter.POST("getUserList", Controllers.GetUserList)
		UserRouter.POST("setUserAuthority", Controllers.SetUserAuthority)
		UserRouter.POST("logout", Controllers.Logout)

	}

	BaseRouter := r.Group("/")
	{
		BaseRouter.POST("register", Controllers.Regist)
		BaseRouter.POST("login", Controllers.Login)

	}
	AuthorityRouter := r.Group("authority")
	{
		AuthorityRouter.POST("createAuthority", Controllers.CreateAuthority)
		AuthorityRouter.POST("deleteAuthority", Controllers.DeleteAuthority).Use(Middlewares.JWTAuth())
		AuthorityRouter.POST("getAuthorityList",Controllers.GetAuthorityList).Use(Middlewares.JWTAuth())
	}

	ProductRouter := r.Group("/")
	{
		ProductRouter.GET("products", Controllers.GetProducts)
		ProductRouter.POST("products", Controllers.CreateProduct)
		ProductRouter.GET("products/:id", Controllers.GetProduct)
		ProductRouter.PUT("products/:id", Controllers.UpdateProduct).Use(Middlewares.JWTAuth())
	}
	MessageRouter := r.Group("/send")
	{
		MessageRouter.GET("message", Controllers.SendMessages)
		MessageRouter.POST("messages", Controllers.MpesaExpress)

	}

	return r
}

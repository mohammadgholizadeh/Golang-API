package main

import (
	"api-app/controllers"
	"api-app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	docs.SwaggerInfo.Title = "APIs Documentation"
	//docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.1"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	public := r.Group("/api")
	public.POST("login", controllers.Login)
	public.POST("login2", controllers.Login2)
	public.POST("register", controllers.Register)
	public.GET("getcaptcha", controllers.GetCaptcha)
	public.POST("tokenValidation", controllers.Token_validation)
	public.POST("resetpassword", controllers.ResetPassword)
	public.POST("ticketing", controllers.Ticketing)
	public.GET("getTicketsListU", controllers.GetTicketsList_User)
	public.POST("messageForTicketU", controllers.MessageForTicket_User)
	public.GET("GetTicketInfoU", controllers.GetTicketInfo_User)
	public.GET("getTicketsListA", controllers.GetTicketsList_Admin)
	public.POST("messageForTicketA", controllers.MessageForTicket_Admin)
	public.GET("GetTicketInfoA", controllers.GetTicketInfo_Admin)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

}

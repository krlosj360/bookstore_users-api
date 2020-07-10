package app

import (
	"bookstore_users-api/controllers/ping"
	"bookstore_users-api/controllers/users"
	"bookstore_users-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Register and Login MedicAgile
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @host localhost:3001
// @BasePath /api/v1
// @query.collection.format multi

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
func mapUrls() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "MediAgile Register And Login"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3001"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	v1 := router.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET("/ping", ping.Ping)
			accounts.GET("/users/:user_id", users.Get)
			accounts.POST("/users", users.Create)
			accounts.PUT("users/:user_id", users.Update)
			accounts.PATCH("users/:user_id", users.Update)
			accounts.DELETE("/users/:user_id", users.Delete)
			accounts.GET("internal/users/search", users.Search)
			accounts.POST("/users/login", users.Login)
		}
		//...
	}
	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	/*
		router.GET("/ping", ping.Ping)
		router.GET("/users/:user_id", users.Get)
		router.POST("/users", users.Create)
		router.PUT("users/:user_id", users.Update)
		router.PATCH("users/:user_id", users.Update)
		router.DELETE("/users/:user_id", users.Delete)
		router.GET("internal/users/search", users.Search)
		router.POST("/users/login", users.Login)

	*/
}

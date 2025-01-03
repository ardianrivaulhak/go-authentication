	package main

	import (
		"go-authentication/controllers"
		"go-authentication/database"
		"go-authentication/middlewares"
		"github.com/gin-gonic/gin"
	)

	func main() {
		// Initialize Database
		database.Connect("root@tcp(localhost:3306)/jwt_demo?parseTime=true")
		database.Migrate()
		
		  // Initialize Router
		router := initRouter()
		router.Run("localhost:8080")
	}

	func initRouter() *gin.Engine {
		router := gin.Default()
		api := router.Group("/api")
		{
			api.POST("/token", controllers.GenerateToken)
			api.POST("/user/register", controllers.RegisterUser)
	
			categoryRoutes := api.Group("/categories")
			{
				categoryRoutes.GET("/", controllers.GetCategories)             
				categoryRoutes.POST("/", middlewares.Auth(), controllers.CreateCategory) 
				categoryRoutes.GET("/:id", controllers.GetCategory)  
				categoryRoutes.PUT("/:id", middlewares.Auth(), controllers.UpdateCategory)
				categoryRoutes.DELETE("/:id", middlewares.Auth(), controllers.DestroyCategort)
			}
	
			secured := api.Group("/secured").Use(middlewares.Auth())
			{
				secured.GET("/ping", controllers.Ping)
			}
		}
		return router
	}
package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rathink4/backend/controllers"
	"github.com/rathink4/backend/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.GoogleConfig()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/oauth_login", controllers.OAuthLogin)
	r.GET("/oauth/callback", controllers.OAuth)
	r.Run()
}

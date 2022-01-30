package main

import (
	handleProfile "biodata/handlers"
	utils "biodata/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialise a new blank Engine instance without middleware
	router := gin.New()
	// attach global logger and recovery middleware
	router.Use(gin.Logger(), gin.Recovery())
	// define handlers
	// home page handler
	router.GET("/home", handleProfile.Home)
	// authorised admin handlers
	authorised := router.Group("/admin")
	authorised.Use(gin.BasicAuth(utils.InitAccounts()))
	{
		authorised.GET("/bio/:id", func(c *gin.Context) {
			id := c.Param("id")
			message := handleProfile.GetProfileById(id)
			c.JSON(200, gin.H{
				"message": message,
			})
		})
	}
	// attach rounter to http server and start serving requests
	router.Run()
}

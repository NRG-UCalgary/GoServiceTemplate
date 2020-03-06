package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port              string
}


func main() {

	//Read Config File
	configuration := Configuration{}
	err := gonfig.GetConf("config/config.development.json", &configuration)
	if err != nil {
		println(err.Error())
		return
	}

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/echo/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hello " + name,
		})
	})

	println("Starting Server on Port: " + configuration.Port)
	router.Run(configuration.Port)
}

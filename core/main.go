package main

import (
	controllers "github.com/YannaAlghamdi/xenelectronic/core/controllers" // new

	"github.com/YannaAlghamdi/xenelectronic/core/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := models.SetupModels() // new
	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/products", controllers.GetAll)
	r.POST("/products", controllers.Create)       // create
	r.GET("/products/:id", controllers.Get)       // find by id
	r.PATCH("/products/:id", controllers.Update)  // update by id
	r.DELETE("/products/:id", controllers.Delete) // delete by id
	r.Run()
}

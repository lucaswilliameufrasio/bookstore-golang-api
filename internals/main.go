package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lucaswilliameufrasio/bookstore-golang-api/controllers"
	"github.com/lucaswilliameufrasio/bookstore-golang-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}

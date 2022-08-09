package routes

import (
	"TODO-DIGITALENT/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	r.POST("/tasks", controller.CreateTask())
	r.GET("/tasks", controller.GetTasks())
	r.GET("/tasks/:id", controller.GetTask())
	r.PUT("/tasks", controller.EditTask())

	r.LoadHTMLGlob("templates/*.html")
	r.GET("/web", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

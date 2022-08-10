package routes

import (
	"TODO-DIGITALENT/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {

	r.POST("/tasks", controller.CreateTask())
	r.GET("/tasks", controller.GetTasks())
	r.GET("/tasks/:id", controller.GetTask())
	r.PUT("/tasks", controller.EditTask())
	r.PATCH("/tasks/:id", controller.FinishTask())

	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

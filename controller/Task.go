package controller

import (
	"TODO-DIGITALENT/config"
	"TODO-DIGITALENT/models"
	"TODO-DIGITALENT/res"
	"TODO-DIGITALENT/response"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection = config.GetCollection(config.DB, "tasks")
var validate = validator.New()

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input res.CreateTaskBody
		var task models.Task
		defer cancel()

		if err := c.BindJSON(&input); err != nil {
			response.BadRequest(c, err.Error())
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			response.BadRequest(c, validationErr.Error())
			return
		}

		filter := bson.D{{Key: "title", Value: input.Title}}
		duplicateErr := taskCollection.FindOne(ctx, filter).Decode(&task)
		if duplicateErr != mongo.ErrNoDocuments {
			response.BadRequest(c, "Data sudah ada")
			return
		}

		newTask := models.Task{
			ID:        primitive.NewObjectID(),
			Title:     input.Title,
			IsDone:    false,
			Assignee:  input.Assignee,
			DueDate:   input.DueDate,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		result, err := taskCollection.InsertOne(ctx, newTask)
		if err != nil {
			response.BadRequest(c, err.Error())
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "berhasil menambahkan data", "status": 201, "data": result})
	}
}

func GetTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var tasks []models.Task

		defer cancel()

		result, err := taskCollection.Find(ctx, bson.M{})
		if err != nil {
			response.BadRequest(c, err.Error())
		}

		defer result.Close(ctx)
		for result.Next(ctx) {
			var singleTask models.Task
			if err = result.Decode(&singleTask); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			fmt.Println(singleTask)
			tasks = append(tasks, singleTask)
		}
		c.JSON(200, gin.H{"data": tasks})
	}
}

func GetTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
		var task models.Task
		defer cancle()

		id := c.Param("id")

		objId, _ := primitive.ObjectIDFromHex(id)
		err := taskCollection.FindOne(ctx, bson.D{{Key: "id", Value: objId}}).Decode(&task)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "status": 404})
			return
		}
		response.Success(c, map[string]interface{}{"task": task})

	}
}

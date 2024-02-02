package main

import (
	"fmt"
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", postEvents)
	server.Run(":8080")

	fmt.Println("server is listening at port localhost:8080")

}

func getEvents(context *gin.Context) {
	// body := context.Request.Body
	// fmt.Println(body)
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)

}

func postEvents(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	// may be more accurate
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please send a valid body", "error": err})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

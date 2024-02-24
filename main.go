package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("events/:id", getEventById)
	server.POST("/events", postEvents)
	server.Run(":8080")

	fmt.Println("server is listening at port localhost:8080")

}

func getEvents(context *gin.Context) {
	// body := context.Request.Body
	// fmt.Println(body)
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again Later"})
		return
	}
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
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again Later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event. Try again Later"})
		return
	}
	context.JSON(http.StatusOK, event)
}

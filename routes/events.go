package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

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

func updateEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event. Try again Later"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	updatedEvent.ID = eventId
	// may be more accurate
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please send a valid body", "error": err})
		return
	}

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event. Try again Later"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Updated Successfully", "event": updatedEvent})
}

func deletedEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Following Event doesn't exist. Try again Later"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event. Try again Later"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted Successfully"})
}

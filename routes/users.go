package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	// may be more accurate
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please send a valid body", "error": err})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again Later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User Created"})
}

// func getEventById(context *gin.Context) {
// 	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err})
// 		return
// 	}

// 	event, err := models.GetEventById(eventId)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event. Try again Later"})
// 		return
// 	}
// 	context.JSON(http.StatusOK, event)
// }

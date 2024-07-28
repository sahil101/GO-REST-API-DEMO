package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/util"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
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

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	// may be more accurate
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data", "error": err})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}
	fmt.Println(user.ID, user.Email, user.Password)
	token, err := util.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successfully", "token": token})
}

package controller

import (
	"eztrust/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Database *gorm.DB
}

func (userController *UserController) CreateUser(ctx *gin.Context) {
	userRequest := domain.User{}
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid request",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// Check if user already exists
	var user domain.User
	userController.Database.Where("username = ?", userRequest.Username).First(&user)
	if user.Id != "" {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "User already exists",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// Generate client secret key
	// privateKey := wireguard.GeneratePrivateKey()
	// userRequest.PrivateKey = privateKey
	// publicKey := wireguard.GeneratePublicKey(privateKey)
	// userRequest.PublicKey = publicKey

	userController.Database.Create(&userRequest)
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "User created successfully",
		Data:    userRequest,
		Status:  http.StatusOK,
	})
}

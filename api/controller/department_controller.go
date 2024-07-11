package controller

import (
	"eztrust/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DepartmentController struct {
	Database *gorm.DB
}

func (departmentController *DepartmentController) CreateDepartment(ctx *gin.Context) {
	departmentRequest := domain.Department{}
	if err := ctx.ShouldBindJSON(&departmentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid request",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// Check if department already exists
	var department domain.Department
	departmentController.Database.Where("name = ?", departmentRequest.Name).First(&department)
	if department.Id != "" {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Department already exists",
			Status:  http.StatusBadRequest,
		})
		return
	}

	departmentController.Database.Create(&departmentRequest)
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Department created successfully",
		Data:    departmentRequest,
		Status:  http.StatusOK,
	})
}

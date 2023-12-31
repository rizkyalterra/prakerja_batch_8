package controllers

import (
	"net/http"
	"prakerja8/configs"
	"prakerja8/models"

	"github.com/labstack/echo/v4"
)


func GetUsersController(c echo.Context) error {

	var users []models.User
	
	result := configs.DB.Preload("CreditCards").Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false,
			Message: "Failed insert data books",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: users,
	})
}
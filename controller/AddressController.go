package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karthiksbh/go-lang-task/config"
	"github.com/karthiksbh/go-lang-task/models"
)

func GetAddressController(c *gin.Context) {
	addresses := []models.ProceedingAdd{}
	err := config.DB.Find(&addresses).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Error has occurred",
		})
		return
	}
	c.JSON(200, &addresses)
}

func GetIDAddressController(c *gin.Context) {
	id := c.Param("id")
	address := models.ProceedingAdd{}
	err := config.DB.Where("identifier=?", id).First(&address).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Address not found",
		})
		return
	}
	if address.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Address not found",
		})
		return
	}
	c.JSON(http.StatusOK, &address)
}

func PostAddressController(c *gin.Context) {
	var add models.ProceedingAdd
	err := c.BindJSON(&add)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error has Occured",
		})
		return
	}

	if add.Identifier == 0 || add.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Make sure you have entered all the details",
		})
		return
	}

	err = config.DB.Create(&add).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error",
		})
		return
	}
	c.JSON(http.StatusCreated, &add)
}

func EditAddressController(c *gin.Context) {
	var add models.ProceedingAdd
	config.DB.Where("id = ?", c.Param("id")).First(&add)
	c.BindJSON(&add)
	config.DB.Save(&add)
	c.JSON(200, &add)
}

func DeleteAddressController(c *gin.Context) {
	var add models.ProceedingAdd
	err := config.DB.Where("id = ?", c.Param("id")).Delete(&add)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error has Occured",
		})
		return
	}
	if err.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No Record Found",
		})
		return
	}
	c.JSON(200, &add)
}

package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karthiksbh/go-lang-task/config"
	"github.com/karthiksbh/go-lang-task/models"
)

func GetPropertyValsController(c *gin.Context) {
	pageSize := 10
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	propvals := []models.Property_vals{}
	err = config.DB.Offset(offset).Limit(pageSize).Find(&propvals).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Error has occurred",
		})
		return
	}
	c.JSON(http.StatusOK, &propvals)
}

func GetIDPropertyvalsController(c *gin.Context) {
	id := c.Param("id")
	propertyvals := models.Property_vals{}
	err := config.DB.Where("identifier=?", id).First(&propertyvals).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Propertyvals not found",
		})
		return
	}
	if propertyvals.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Propertyvals not found",
		})
		return
	}
	c.JSON(http.StatusOK, &propertyvals)
}

func PostPropertyvalsController(c *gin.Context) {
	var add models.Property_vals
	err := c.BindJSON(&add)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error has Occured",
		})
		return
	}

	if add.Identifier == 0 || add.Serial_Number == 0 {
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

func EditPropertyvalsController(c *gin.Context) {
	var add models.Property_vals
	config.DB.Where("id = ?", c.Param("id")).First(&add)
	c.BindJSON(&add)
	config.DB.Save(&add)
	c.JSON(200, &add)
}

func DeletePropertyvalsController(c *gin.Context) {
	var add models.Property_vals
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

package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karthiksbh/go-lang-task/config"
	"github.com/karthiksbh/go-lang-task/models"
)

func GetAddressController(c *gin.Context) {
	pageSize := 10
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	addresses := []models.ProceedingAdd{}
	err = config.DB.Offset(offset).Limit(pageSize).Find(&addresses).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Error has occurred",
		})
		return
	}
	c.JSON(http.StatusOK, &addresses)
}

func GetAddressFinalController(c *gin.Context) {
	addresses := []models.AddressInformation{}
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

	config.DB.Create(&add)
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
	config.DB.Where("identifier = ?", c.Param("id")).First(&add)
	c.BindJSON(&add)
	config.DB.Save(&add)
	c.JSON(200, &add)
}

func DeleteAddressController(c *gin.Context) {
	var add models.ProceedingAdd
	config.DB.Where("identifier = ?", c.Param("id")).First(&add)
	config.DB.Delete(&add)
	c.JSON(http.StatusOK, gin.H{
		"message": "Address deleted successfully",
		"address": &add,
	})
}

func FilterBased(c *gin.Context) {
	var searchFilter struct {
		TypeCode string `json:"typeCode"`
		Name     string `json:"name"`
		City     string `json:"city"`
	}

	err := c.ShouldBindJSON(&searchFilter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	filter_inps := map[string]interface{}{}
	if searchFilter.TypeCode == "" && searchFilter.Name == "" && searchFilter.City == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Enter atleast one input",
		})
		return
	}
	if searchFilter.TypeCode != "" {
		filter_inps["type_code"] = searchFilter.TypeCode
	}
	if searchFilter.Name != "" {
		filter_inps["name"] = searchFilter.Name
	}
	if searchFilter.City != "" {
		filter_inps["city"] = searchFilter.City
	}

	ops := []models.ProceedingAdd{}
	err = config.DB.Where(filter_inps).Find(&ops).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Error has occurred",
		})
		return
	}
	c.JSON(http.StatusOK, &ops)
}

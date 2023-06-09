package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Asylniet/golang_lab3/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GetItems(c *gin.Context) {
	items, result := services.GetItems()
	if errors.Is(result.Error, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	item, result := services.GetItem(id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func AddItem(c *gin.Context) {
	name := c.Query("name")
	priceStr := c.Query("price")

	if name == "" || priceStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Enter correct data",
		})
		return
	}

	price, priceErr := strconv.Atoi(priceStr)
	if priceErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Enter correct price",
		})
		return
	}

	result := services.AddItem(name, price)
	fmt.Println(result.Error)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't add item, please try again",
		})
		return
	}

	c.Status(http.StatusOK)
}

func RemoveItem(c *gin.Context) {
	id := c.Param("id")
	result := services.RemoveItem(id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't remove item, please try again",
		})
		return
	}
	c.Status(http.StatusOK)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	priceStr := c.Query("price")

	price, priceErr := strconv.Atoi(priceStr)
	if priceErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Enter correct price",
		})
		return
	}

	result, error := services.UpdateItem(name, id, price)

	if error == "No item found" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No item found, please try again",
		})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't update item, please try again",
		})
		return
	}
	c.Status(http.StatusOK)
}

func SearchItems(c *gin.Context) {
	name := strings.ToLower(c.Query("query"))
	filter := c.Query("filter")
	sort := c.Query("sort")

	items, result := services.SearchItems(name, filter, sort)

	if len(items) == 0 || result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No items found, please try again",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

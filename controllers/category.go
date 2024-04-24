package controllers

import (
	"net/http"
	"quiz3/database"
	"quiz3/repository"
	"quiz3/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	categories, err := repository.GetAllCategories(database.DbConnection)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": categories,
		})
	}
}

func GetBooksWithCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	books, err := repository.GetBooksWithCategory(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": books,
		})
	}
}

func InsertCategory(c *gin.Context) {
	var category structs.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request body",
			"result":  err,
		})
		return
	}

	err = repository.InsertCategory(database.DbConnection, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in database",
			"result":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert category",
	})
}

func UpdateCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request body",
			"result":  err,
		})
		return
	}

	category.ID = id
	err = repository.UpdateCategory(database.DbConnection, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in database",
			"result":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success update category",
	})
}

func DeleteCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = id
	err := repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in database",
			"result":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success delete category",
	})
}

package controllers

import (
	"net/http"
	"quiz3/database"
	"quiz3/repository"
	"quiz3/structs"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	books, err := repository.GetAllBooks(database.DbConnection)
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

func InsertBook(c *gin.Context) {
	var book structs.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request body",
			"result":  err,
		})
		return
	}

	regex, err := regexp.Compile(`(http(s?):)([/|.|\w|\s|-])*\.(?:jpg|gif|png)`)
	if !regex.Match([]byte(book.ImageUrl)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request body",
			"result":  "Please input image_url properly",
		})
		return
	}

	if !(book.ReleaseYear >= 1980 && book.ReleaseYear <= 2021) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request body",
			"result":  "release_year must be 1980 to 2021",
		})
		return
	}

	if book.TotalPage <= 100 {
		book.Thickness = "tipis"
	} else if book.TotalPage >= 101 && book.TotalPage <= 200 {
		book.Thickness = "sedang"
	} else if book.TotalPage <= 201 {
		book.Thickness = "tebal"
	}

	err = repository.InsertBook(database.DbConnection, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in database",
			"result":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert book",
	})
}

func UpdateBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request body",
			"result":  err,
		})
		return
	}

	regex, err := regexp.Compile(`(http(s?):)([/|.|\w|\s|-])*\.(?:jpg|gif|png)`)
	if !regex.Match([]byte(book.ImageUrl)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request body",
			"result":  "Please input image_url properly",
		})
		return
	}

	if !(book.ReleaseYear >= 1980 && book.ReleaseYear <= 2021) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request body",
			"result":  "release_year must be 1980 to 2021",
		})
		return
	}

	switch {
	case book.TotalPage <= 100:
		book.Thickness = "tipis"
		break
	case book.TotalPage >= 101 && book.TotalPage <= 200:
		book.Thickness = "sedang"
		break
	case book.TotalPage <= 201:
		book.Thickness = "tebal"
		break
	}

	book.ID = id
	err = repository.UpdateBook(database.DbConnection, book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error in database",
			"result":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success update book",
	})
}

func DeleteBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id
	err := repository.DeleteBook(database.DbConnection, book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error in database",
			"result":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success delete book",
	})
}

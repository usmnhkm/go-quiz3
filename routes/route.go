package routes

import (
	"quiz3/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/segitiga-sama-sisi/", controllers.SegitigaSamaSisi)
	router.GET("/persegi", controllers.Persegi)
	router.GET("/persegi-panjang", controllers.PersegiPanjang)
	router.GET("/lingkaran", controllers.Lingkaran)

	router.GET("/categories", controllers.GetAllCategories)
	router.GET("/categories/:id/books", controllers.GetBooksWithCategory)
	router.POST("/categories", controllers.InsertCategory)
	router.PUT("/categories/:id", controllers.UpdateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)

	router.GET("/books", controllers.GetAllBooks)
	router.POST("/books", controllers.InsertBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
	return router
}

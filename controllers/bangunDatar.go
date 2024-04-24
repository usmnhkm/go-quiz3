package controllers

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func SegitigaSamaSisi(c *gin.Context) {
	alas, _ := strconv.Atoi(c.Query("alas"))
	tinggi, _ := strconv.Atoi(c.Query("tinggi"))
	hitung := c.Query("hitung")
	var hasil float64

	if hitung == "keliling" {
		hasil = float64(alas * 3)
	} else if hitung == "luas" {
		hasil = float64(alas * tinggi / 2)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Metode hitung tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": hasil,
	})
}

func Persegi(c *gin.Context) {
	sisi, _ := strconv.Atoi(c.Query("sisi"))
	hitung := c.Query("hitung")
	var hasil float64

	if hitung == "keliling" {
		hasil = float64(sisi * 4)
	} else if hitung == "luas" {
		hasil = float64(sisi * sisi)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Metode hitung tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": hasil,
	})
}

func PersegiPanjang(c *gin.Context) {
	panjang, _ := strconv.Atoi(c.Query("panjang"))
	lebar, _ := strconv.Atoi(c.Query("lebar"))
	hitung := c.Query("hitung")
	var hasil float64

	if hitung == "keliling" {
		hasil = float64(2 * (panjang + lebar))
	} else if hitung == "luas" {
		hasil = float64(panjang * lebar)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Metode hitung tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": hasil,
	})
}

func Lingkaran(c *gin.Context) {
	jariJari, _ := strconv.Atoi(c.Query("jariJari"))
	hitung := c.Query("hitung")
	var hasil float64

	if hitung == "keliling" {
		hasil = math.Pi * float64(jariJari) * float64(jariJari)
	} else if hitung == "luas" {
		hasil = math.Pi * float64(jariJari) * 2
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Metode hitung tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": hasil,
	})
}

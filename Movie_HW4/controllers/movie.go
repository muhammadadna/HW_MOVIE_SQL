package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/structs"
)

// Get One Movie
func (idb *InDB) GetMovie(c *gin.Context) {
	var (
		movie  structs.Movie
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&movie).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": movie,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Get All Data
func (idb *InDB) GetMovies(c *gin.Context) {
	var (
		movies []structs.Movie
		result gin.H
	)

	idb.DB.Find(&movies)
	if len(movies) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": movies,
			"count":  len(movies),
		}
	}

	c.JSON(http.StatusOK, result)
}

// Insert Data
func (idb *InDB) CreateMovie(c *gin.Context) {
	var (
		movie  structs.Movie
		result gin.H
	)
	title := c.PostForm("title")
	slug := c.PostForm("slug")
	description := c.PostForm("description")
	duration := c.PostForm("duration")
	durationMovie, _ := strconv.Atoi(duration) // convert int to string
	image := c.PostForm("image")
	movie.Title = title
	movie.Slug = slug
	movie.Description = description
	movie.Duration = durationMovie
	movie.Image = image
	idb.DB.Create(&movie)
	result = gin.H{
		"result": movie,
	}
	c.JSON(http.StatusOK, result)
}

// Update Data
func (idb *InDB) UpdateMovie(c *gin.Context) {
	id := c.Query("id")
	title := c.PostForm("title")
	slug := c.PostForm("slug")
	description := c.PostForm("description")
	duration := c.PostForm("duration")
	durationMovie, _ := strconv.Atoi(duration)
	image := c.PostForm("image")
	var (
		movie    structs.Movie
		newMovie structs.Movie
		result   gin.H
	)

	err := idb.DB.First(&movie, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data tidak ditemukan !",
		}
	}
	newMovie.Title = title
	newMovie.Slug = slug
	newMovie.Description = description
	newMovie.Duration = durationMovie
	newMovie.Image = image
	err = idb.DB.Model(&movie).Updates(newMovie).Error
	if err != nil {
		result = gin.H{
			"result": "update gagal!",
		}
	} else {
		result = gin.H{
			"result": "update sukses!",
		}
	}
	c.JSON(http.StatusOK, result)

}

//Delete Data
func (idb *InDB) DeleteMovie(c *gin.Context) {
	var (
		movie  structs.Movie
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&movie, id).Error
	if err != nil {
		result = gin.H{
			"result": "gagal menghapus!",
		}
	} else {
		result = gin.H{
			"result": "sukses menghapus!",
		}
	}
	c.JSON(http.StatusOK, result)
}

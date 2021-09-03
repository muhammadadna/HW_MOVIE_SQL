package main

import (
	"Movie_HW4/config"
	"Movie_HW4/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/movie/:slug", inDB.GetMovie)
	router.GET("/movie", inDB.GetMovie)
	router.POST("/movie", inDB.CreateMovie)
	router.PUT("/movie/:slug", inDB.UpdateMovie)
	router.DELETE("/movie/:slug", inDB.DeleteMovie)
	router.Run(":8080")
}

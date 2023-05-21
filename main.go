package main

import (
	"github.com/ahay12/go-api/controller/page_controller"
	"github.com/ahay12/go-api/helper"
	"github.com/ahay12/go-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	model.ConnectionDB()

	r.GET("/api/page", page_controller.Index)
	r.GET("/api/page/:id", page_controller.Show)
	r.POST("/api/page", page_controller.Create)
	r.PUT("/api/page/:id", page_controller.Update)
	r.DELETE("/api/page", page_controller.Delete)

	s := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	helper.PanicIfError(s.ListenAndServe())
}

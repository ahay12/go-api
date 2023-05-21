package page_controller

import (
	"encoding/json"
	"github.com/ahay12/go-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Index(c *gin.Context) {
	var pages []model.Page

	model.DB.Find(&pages)
	c.JSON(http.StatusOK, gin.H{"pages": pages})
}

func Show(c *gin.Context) {
	var page model.Page

	if err := model.DB.First(&page, c.Param("id")).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Page not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"page": page})
}

func Create(c *gin.Context) {
	var page model.Page

	if err := c.ShouldBindJSON(&page); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	model.DB.Create(&page)
	c.JSON(http.StatusOK, gin.H{"page": page})
}

func Update(c *gin.Context) {
	var page model.Page

	id := c.Param("id")

	if err := c.ShouldBindJSON(&page); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if model.DB.Model(&page).Where("id = ?", id).Updates(&page).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak ada data yang diupdate"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})
}

func Delete(c *gin.Context) {
	var page model.Page

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if model.DB.Delete(&page, input.Id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak ada data yang dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

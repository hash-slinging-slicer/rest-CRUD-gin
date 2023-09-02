package productcontroller

import (
	"api-crud-gin/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(r *gin.Context) {
	var produk []models.Product

	models.DB.Find(&produk)
	r.JSON(http.StatusOK, gin.H{"produk": produk})
}

func Detail(r *gin.Context) {
	var produk models.Product

	// GET ID PARAM
	id := r.Param("id")

	// GET
	if err := models.DB.First(&produk, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			r.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
			return
		default:
			r.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	r.JSON(http.StatusOK, gin.H{"produk": produk})
}

func Tambah(r *gin.Context) {
	var produk models.Product

	if err := r.ShouldBindJSON(&produk); err != nil {
		r.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&produk)
	r.JSON(http.StatusOK, gin.H{"produk": produk})
}

func Update(r *gin.Context) {
	var produk models.Product

	// Variabel Form
	id := r.Param("id")

	if err := r.ShouldBindJSON(&produk); err != nil {
		r.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&produk).Where("id = ?", id).Updates(&produk).RowsAffected == 0 {
		r.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak Bisa Update"})
		return
	}

	r.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})

}

func Hapus(r *gin.Context) {
	var produk models.Product

	// Variabel Form
	var input struct {
		id json.Number
	}

	// Convert to int
	id, _ := input.id.Int64()
	if err := r.ShouldBindJSON(&id); err != nil {
		r.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Delete(&produk, id).RowsAffected == 0 {
		r.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak Bisa Hapus Produk"})
		return
	}

	r.JSON(http.StatusOK, gin.H{"message": "Produk Terhapus"})
}

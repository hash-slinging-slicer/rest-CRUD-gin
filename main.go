package main

import (
	productcontroller "api-crud-gin/controllers/productController"
	"api-crud-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.KonekDB()

	endPointProduk := r.Group("/api/produk")
	{

		endPointProduk.GET("/", productcontroller.Index)
		endPointProduk.GET("/:id", productcontroller.Detail)
		endPointProduk.POST("/tambah", productcontroller.Tambah)
		endPointProduk.PUT("/update/:id", productcontroller.Update)
		endPointProduk.DELETE("/hapus", productcontroller.Hapus)
	}

	r.Run()
}

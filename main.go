package main

import (
	"rest_api_gorm/handler"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Rute untuk mendapatkan daftar Todo
    router.GET("/todos", handler.GetTodos)

    // Rute untuk menambahkan Todo baru, memperbarui, dan menghapus Todo

    router.Run(":8080")
}

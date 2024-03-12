package handler

import (
	"net/http"

	"rest_api_gorm/database"
	"rest_api_gorm/model"

	"github.com/gin-gonic/gin"
)


func GetTodos(c *gin.Context) {
    db, err := database.Connect()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal terhubung ke database"})
        return
    }
    defer database.Close(db) // Menggunakan fungsi Close dari package database

    var todos []model.Todo
    if err := db.Find(&todos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan daftar todos"})
        return
    }

    c.JSON(http.StatusOK, todos)
}

// Implementasi handler untuk menambahkan Todo baru, memperbarui, dan menghapus Todo

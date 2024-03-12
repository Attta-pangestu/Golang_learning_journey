package model

type Todo struct {
    ID     uint   `gorm:"primary_key" json:"id"`
    Title  string `json:"title"`
    Status string `json:"status"`
}

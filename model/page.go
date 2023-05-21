package model

type Page struct {
	Id    int    `gorm:"primaryKey" json:"id"`
	Title string `gorm:"type:varchar(100)" json:"title"`
	Body  string `gorm:"type:varchar(500)" json:"body"`
}

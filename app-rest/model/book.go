package model

type Book struct {
	ID      int64  `gorm:"int(64) primaryKey autoIncrement " form:"id" json:"id"`
	Title   string `gorm:"varchar(40)" json:"title" form:"title"`
	Content string `gorm:"varchar(40)" json:"content" form:"content"`
}

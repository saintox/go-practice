package models

import (
	"github.com/jinzhu/gorm"
	"github.com/saintox/go-practice/freecodecamp/bookstore-management/pkg/configs"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"Name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Year        int8   `json:"year"`
}

func init() {
	configs.ConnectDB()

	db = configs.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)

	return b
}

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)

	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)

	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)

	return book
}
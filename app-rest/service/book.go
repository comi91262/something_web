package service

import "github.com/ykonomi/something_web/model"

type BookService struct{}

func (BookService) SetBook(book *model.Book) error {
	result := dbConn.Create(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (BookService) GetBookList() []model.Book {
	var books []model.Book

	result := dbConn.Select("id", "title", "content").Limit(10).Find(&books)
	if result.Error != nil {
		panic(result.Error)
	}
	return books
}

func (BookService) UpdateBook(newBook *model.Book) error {
	result := dbConn.Save(&newBook)
	if result.Error != nil {
		panic(result.Error)
	}
	return nil
}

func (BookService) DeleteBook(id int) error {
	result := dbConn.Delete(&model.Book{}, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return nil
}

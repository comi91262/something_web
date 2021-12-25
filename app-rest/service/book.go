package service

import "github.com/ykonomi/something_web/model"

type BookService struct{}

func (BookService) SetBook(book *model.Book) error {
	if result := dbConn.Create(&book); result.Error != nil {
		return result.Error
	}
	return nil
}

func (BookService) GetBookList() ([]model.Book, error) {
	var books []model.Book

	result := dbConn.Select("id", "title", "content").Limit(10).Find(&books)
	if result.Error != nil {
		return books, result.Error
	}
	return books, nil
}

func (BookService) UpdateBook(newBook *model.Book) error {
	if result := dbConn.Save(&newBook); result.Error != nil {
		return result.Error
	}
	return nil
}

func (BookService) DeleteBook(id int) error {
	if result := dbConn.Delete(&model.Book{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

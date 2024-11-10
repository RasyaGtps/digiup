package repository

import (
	"base-gin/domain/dao"
	"errors"	
	"time"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(book *dao.Book) error {
	return r.db.Create(&book).Error
}

func (r *BookRepository) GetList() ([]dao.Book, error) {
	var books []dao.Book
	err := r.db.
		Joins("BookPublisher").
		Joins("BookAuthor").
		Find(&books).Error
	return books, err
}

func (r *BookRepository) GetByID(id uint) (dao.Book, error) {
    var book dao.Book
    err := r.db.
        Joins("BookPublisher").
        Joins("BookAuthor").
        First(&book, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return book, errors.New("book not found")
    }
    return book, err
}

// Tambahkan method baru khusus untuk verifikasi delete
func (r *BookRepository) GetByIDUnscoped(id uint) (dao.Book, error) {
    var book dao.Book
    err := r.db.
        Unscoped().
        Joins("BookPublisher").
        Joins("BookAuthor").
        First(&book, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return book, errors.New("book not found")
    }
    return book, err
}

// Di repository/book.go
func (r *BookRepository) Update(book *dao.Book) error {
    result := r.db.Model(&dao.Book{}).Where("id = ?", book.ID).Updates(map[string]interface{}{
        "title":        book.Title,
        "subtitle":     book.Subtitle,
        "publisher_id": book.PublisherID,
        "author_id":    book.AuthorID,
        "updated_at":   time.Now(),
    })
    
    if result.Error != nil {
        return result.Error
    }
    
    if result.RowsAffected == 0 {
        return errors.New("book not found")
    }
    
    return nil
}

func (r *BookRepository) Delete(id uint) error {
    result := r.db.Delete(&dao.Book{}, id)
    if result.Error != nil {
        return result.Error
    }
    
    if result.RowsAffected == 0 {
        return errors.New("book not found")
    }
    
    return nil
}

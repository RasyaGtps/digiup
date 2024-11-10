package repository

import (
	"base-gin/domain/dao"
	"base-gin/exception"
	"errors"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) Create(author *dao.Author) error {
	return r.db.Create(author).Error
}

func (r *AuthorRepository) GetList() ([]dao.Author, error) {
	var authors []dao.Author
	err := r.db.Find(&authors).Error
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *AuthorRepository) GetByID(id uint) (*dao.Author, error) {
	var author dao.Author
	err := r.db.First(&author, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.ErrDataNotFound
	} else if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *AuthorRepository) Update(author *dao.Author) error {
	result := r.db.Model(&dao.Author{}).Where("id = ?", author.ID).Updates(author)
	if result.RowsAffected == 0 {
		return exception.ErrDataNotFound
	}
	return result.Error
}

func (r *AuthorRepository) Delete(id uint) error {
	result := r.db.Delete(&dao.Author{}, id)
	if result.RowsAffected == 0 {
		return exception.ErrDataNotFound
	}
	return result.Error
}

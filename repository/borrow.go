package repository

import (
	"base-gin/domain/dao"
	"errors"
	"time"
	"gorm.io/gorm"
)

type BorrowingRepository struct {
	db *gorm.DB
}

func NewBorrowingRepository(db *gorm.DB) *BorrowingRepository {
	return &BorrowingRepository{db: db}
}

func (r *BorrowingRepository) Create(borrowing *dao.Borrowing) error {
	return r.db.Create(&borrowing).Error
}

func (r *BorrowingRepository) GetList() ([]dao.Borrowing, error) {
	var borrowings []dao.Borrowing
	err := r.db.
		Joins("Book").
		Joins("Person").
		Find(&borrowings).Error
	return borrowings, err
}

func (r *BorrowingRepository) GetByID(id uint) (dao.Borrowing, error) {
    var borrowing dao.Borrowing
    err := r.db.First(&borrowing, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return borrowing, errors.New("borrowing not found")
    }
    return borrowing, err
}
func (r *BorrowingRepository) Update(borrowing *dao.Borrowing) error {
    result := r.db.Model(&dao.Borrowing{}).Where("id = ?", borrowing.ID).Updates(map[string]interface{}{
        "return_date": borrowing.ReturnDate,
        "updated_at":  time.Now(),
    })
    
    if result.Error != nil {
        return result.Error
    }
    
    if result.RowsAffected == 0 {
        return errors.New("borrowing not found")
    }
    
    return nil
}

func (r *BorrowingRepository) Delete(id uint) error {
    result := r.db.Delete(&dao.Borrowing{}, id)
    if result.Error != nil {
        return result.Error
    }
    
    if result.RowsAffected == 0 {
        return errors.New("borrowing not found")
    }
    
    return nil
}
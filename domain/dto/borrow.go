package dto

import (
	"base-gin/domain/dao"
	"time"
)

type BorrowingDTO struct {
	BookID     uint       `json:"book_id" binding:"required"`
	PersonID   uint       `json:"person_id" binding:"required"`
	BorrowDate time.Time  `json:"borrow_date" binding:"required"`
	ReturnDate *time.Time `json:"return_date"`
}

func (b *BorrowingDTO) ToEntity() dao.Borrowing {
	return dao.Borrowing{
		BookID:     b.BookID,
		PersonID:   b.PersonID,
		BorrowDate: b.BorrowDate,
		ReturnDate: b.ReturnDate,
	}
}

type BorrowingResp struct {
	ID         uint       `json:"id"`
	BookID     uint       `json:"book_id"`
	PersonID   uint       `json:"person_id"`
	BorrowDate time.Time  `json:"borrow_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"`
}

func (b *BorrowingResp) FromEntity(borrowing *dao.Borrowing) {
	b.ID = borrowing.ID
	b.BookID = borrowing.BookID
	b.PersonID = borrowing.PersonID
	b.BorrowDate = borrowing.BorrowDate
	b.ReturnDate = borrowing.ReturnDate
}

type BorrowingUpdate struct {
    ID         uint       `json:"-"`
    ReturnDate *time.Time `json:"return_date"`
}

func (b *BorrowingUpdate) ToEntity() *dao.Borrowing {
    return &dao.Borrowing{
        ID:         b.ID,
        ReturnDate: b.ReturnDate,
    }
}

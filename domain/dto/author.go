package dto

import (
	"base-gin/domain/dao"
	"time"
)

type AuthorDTO struct {
	ID        uint       `json:"-"`
	FullName  string     `json:"full_name" binding:"required,min=2,max=56"`
	Gender    *string    `json:"gender" binding:"omitempty,oneof=m f"`
	BirthDate *time.Time `json:"birth_date" binding:"omitempty"`
}

func (a *AuthorDTO) ToEntity() dao.Author {
	return dao.Author{
		ID:        a.ID,
		FullName:  a.FullName,
		Gender:    *a.Gender,
		BirthDate: *a.BirthDate,
	}
}

type AuthorResp struct {
	ID        uint       `json:"id"`
	FullName  string     `json:"full_name"`
	Gender    *string    `json:"gender"`
	BirthDate *time.Time `json:"birth_date"`
}

func (a *AuthorResp) FromEntity(entity *dao.Author) {
	a.ID = entity.ID
	a.FullName = entity.FullName
	a.Gender = &entity.Gender
	a.BirthDate = &entity.BirthDate
}

type AuthorUpdate struct {
	ID        uint       `json:"-"`
	FullName  string     `json:"full_name" binding:"required,min=2,max=56"`
	Gender    *string    `json:"gender" binding:"omitempty,oneof=m f"`
	BirthDate *time.Time `json:"birth_date" binding:"omitempty"`
}

func (a *AuthorUpdate) ToEntity() dao.Author {
	return dao.Author{
		ID:        a.ID,
		FullName:  a.FullName,
		Gender:    *a.Gender,
		BirthDate: *a.BirthDate,
	}
}

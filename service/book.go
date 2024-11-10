package service

import (
	"base-gin/domain/dto"
	"base-gin/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(bookRepo *repository.BookRepository) *BookService {
	return &BookService{repo: bookRepo}
}

func (s *BookService) Create(params *dto.BookDTO) error {
	newItem := params.ToEntity()
	return s.repo.Create(&newItem)
}

func (s *BookService) GetByID(id uint) (dto.BookResp, error) {
	var resp dto.BookResp

	item, err := s.repo.GetByID(id)
	if err != nil {
		return resp, err
	}

	resp.FromEntity(&item)

	return resp, nil
}

func (s *BookService) GetList(params *dto.Filter) ([]dto.BookResp, error) {
	var resp []dto.BookResp

	items, err := s.repo.GetList()
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		var t dto.BookResp
		t.FromEntity(&item)

		resp = append(resp, t)
	}

	return resp, nil
}

// Di service/book_service.go
func (s *BookService) Update(input *dto.BookUpdate) error {
    // Convert DTO to entity
    book := input.ToEntity()
    
    // Update di repository
    err := s.repo.Update(book)
    if err != nil {
        return err
    }
    
    return nil
}

func (s *BookService) Delete(id uint) error {
    // Cek apakah buku ada
    _, err := s.repo.GetByID(id)
    if err != nil {
        return err
    }

    // Hapus buku
    return s.repo.Delete(id)
}

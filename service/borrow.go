package service

import (
	"base-gin/domain/dto"
	"base-gin/repository"
)

type BorrowingService struct {
	repo *repository.BorrowingRepository
}

func NewBorrowingService(borrowingRepo *repository.BorrowingRepository) *BorrowingService {
	return &BorrowingService{repo: borrowingRepo}
}

func (s *BorrowingService) Create(params *dto.BorrowingDTO) error {
	newBorrowing := params.ToEntity()
	return s.repo.Create(&newBorrowing)
}

func (s *BorrowingService) GetByID(id uint) (dto.BorrowingResp, error) {
	var resp dto.BorrowingResp
	borrowing, err := s.repo.GetByID(id)
	if err != nil {
		return resp, err
	}
	resp.FromEntity(&borrowing)
	return resp, nil
}

func (s *BorrowingService) GetList() ([]dto.BorrowingResp, error) {
	var resp []dto.BorrowingResp
	borrowings, err := s.repo.GetList()
	if err != nil {
		return nil, err
	}

	for _, borrowing := range borrowings {
		var borrowingResp dto.BorrowingResp
		borrowingResp.FromEntity(&borrowing)
		resp = append(resp, borrowingResp)
	}

	return resp, nil
}

func (s *BorrowingService) Update(input *dto.BorrowingUpdate) error {
    // Convert DTO to entity
    borrowing := input.ToEntity()
    
    // Update di repository
    return s.repo.Update(borrowing)
}
func (s *BorrowingService) Delete(id uint) error {
    // Cek apakah borrowing ada
    _, err := s.repo.GetByID(id)
    if err != nil {
        return err
    }

    // Hapus borrowing
    return s.repo.Delete(id)
}
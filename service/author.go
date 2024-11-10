package service

import (
	"base-gin/domain/dto"
	"base-gin/repository"
)

type AuthorService struct {
	repo *repository.AuthorRepository
}

func NewAuthorService(repo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) Create(params *dto.AuthorDTO) error {
	author := params.ToEntity()
	return s.repo.Create(&author)
}

func (s *AuthorService) GetList() ([]dto.AuthorResp, error) {
	authors, err := s.repo.GetList()
	if err != nil {
		return nil, err
	}

	var response []dto.AuthorResp
	for _, author := range authors {
		var resp dto.AuthorResp
		resp.FromEntity(&author)
		response = append(response, resp)
	}

	return response, nil
}

func (s *AuthorService) GetByID(id uint) (dto.AuthorResp, error) {
	var response dto.AuthorResp

	author, err := s.repo.GetByID(id)
	if err != nil {
		return response, err
	}

	response.FromEntity(author)
	return response, nil
}

func (s *AuthorService) Update(params *dto.AuthorUpdate) error {
	author := params.ToEntity()
	if err := s.repo.Update(&author); err != nil {
		return err
	}
	return nil
}

func (s *AuthorService) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

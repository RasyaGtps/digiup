package dto

import "base-gin/domain/dao"

type BookDTO struct {
	ID          uint    `json:"-"`
	Title       string  `json:"title" binding:"required,min=2,max=56"`
	Subtitle    *string `json:"subtitle" binding:"min=2,max=56"`
	PublisherID uint    `gorm:"not null;"`
	AuthorID    uint    `gorm:"not null"`
}

func (o *BookDTO) ToEntity() dao.Book {
	return dao.Book{
		ID:          o.ID,
		Title:       o.Title,
		Subtitle:    o.Subtitle,
		PublisherID: o.PublisherID,
		AuthorID:    o.AuthorID,
	}
}

type BookResp struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Subtitle    *string `json:"subtitle"`
	PublisherID uint    `json:"publisher_id"`
	AuthorID    uint    `json:"author_id"`
}

func (o *BookResp) FromEntity(item *dao.Book) {
	o.ID = item.ID
	o.Title = item.Title
	o.Subtitle = item.Subtitle
	o.PublisherID = item.PublisherID
	o.AuthorID = item.AuthorID
}

type BookUpdate struct {
    ID          uint    `json:"-"`
    Title       string  `json:"title" binding:"required,min=2,max=56"`
    Subtitle    *string `json:"subtitle" binding:"min=2,max=56"`
    PublisherID uint    `json:"publisher_id" binding:"required"`
    AuthorID    uint    `json:"author_id" binding:"required"`
}

func (b *BookUpdate) ToEntity() *dao.Book {
    return &dao.Book{
        ID:          b.ID,
        Title:       b.Title,
        Subtitle:    b.Subtitle,
        PublisherID: b.PublisherID,
        AuthorID:    b.AuthorID,
    }
}
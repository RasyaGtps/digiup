package rest

import (
	"base-gin/domain/dto"
	"base-gin/exception"
	"base-gin/server"
	"base-gin/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	hr      *server.Handler
	service *service.BookService
}

func NewBookHandler(handler *server.Handler, bookService *service.BookService) *BookHandler {
	return &BookHandler{hr: handler, service: bookService}
}

func (h *BookHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootBook)
	grp.POST("", h.hr.AuthAccess(), h.create)
	grp.GET("", h.getList)
	grp.GET("/:id", h.getByID)
	grp.PUT("/:id", h.hr.AuthAccess(), h.update)
	grp.DELETE("/:id", h.hr.AuthAccess(), h.delete)
}

// create godoc
//
//	@Summary Create a book
//	@Description Create a new book.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param detail body dto.BookDTO true "Book's detail"
//	@Success 201 {object} dto.SuccessResponse[any]
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /books [post]
func (h *BookHandler) create(c *gin.Context) {
	var req dto.BookDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		h.hr.BindingError(err)
		return
	}

	err := h.service.Create(&req)
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse[any]{
		Success: true,
		Message: "Book created successfully",
	})
}

// getList godoc
//
//	@Summary Get a list of books
//	@Description Get a list of all books.
//	@Produce json
//	@Success 200 {object} dto.SuccessResponse[[]dto.BookDTO]
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /books [get]
func (h *BookHandler) getList(c *gin.Context) {
	filter := &dto.Filter{}
	data, err := h.service.GetList(filter)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}
		return
	}

	var bookDTOs []dto.BookDTO
	for _, item := range data {
		var bookDTO dto.BookDTO
		bookDTO.Title = item.Title
		bookDTO.Subtitle = item.Subtitle
		bookDTO.PublisherID = item.PublisherID
		bookDTO.AuthorID = item.AuthorID
		bookDTOs = append(bookDTOs, bookDTO)
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.BookDTO]{
		Success: true,
		Message: "List of books",
		Data:    bookDTOs,
	})
}

// getByID godoc
//
//	@Summary Get a book's detail
//	@Description Get details of a specific book by ID.
//	@Produce json
//	@Param id path int true "Book ID"
//	@Success 200 {object} dto.SuccessResponse[dto.BookDTO]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /books/{id} [get]
func (h *BookHandler) getByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
		return
	}

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}
		return
	}

	var bookDTO dto.BookDTO
	bookDTO.Title = data.Title
	bookDTO.Subtitle = data.Subtitle
	bookDTO.PublisherID = data.PublisherID
	bookDTO.AuthorID = data.AuthorID

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.BookDTO]{
		Success: true,
		Message: "Book details",
		Data:    bookDTO,
	})
}

// update godoc
//
//	@Summary Update a book's detail
//	@Description Update details of a specific book by ID.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param id path int true "Book ID"
//	@Param detail body dto.BookUpdate true "Book's updated detail"
//	@Success 200 {object} dto.SuccessResponse[any]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /books/{id} [put]
// Di rest/book_handler.go
// Di rest/book_handler.go
func (h *BookHandler) update(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
        return
    }

    var input dto.BookUpdate
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, h.hr.ErrorResponse(err.Error()))
        return
    }

    // Set ID dari parameter ke input
    input.ID = uint(id)

    // Gunakan service untuk update
    err = h.service.Update(&input)
    if err != nil {
        h.hr.ErrorInternalServer(c, err)
        return
    }

    c.JSON(http.StatusOK, dto.SuccessResponse[any]{
        Success: true,
        Message: "Book updated successfully",
    })
}

// delete godoc
//
//	@Summary Delete a book
//	@Description Delete a specific book by ID.
//	@Produce json
//	@Security BearerAuth
//	@Param id path int true "Book ID"
//	@Success 200 {object} dto.SuccessResponse[any]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /books/{id} [delete]
func (h *BookHandler) delete(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
        return
    }

    err = h.service.Delete(uint(id))
    if err != nil {
        if err.Error() == "book not found" {
            c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
            return
        }
        h.hr.ErrorInternalServer(c, err)
        return
    }

    c.JSON(http.StatusOK, dto.SuccessResponse[any]{
        Success: true,
        Message: "Book deleted successfully",
    })
}

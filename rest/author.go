// rest/author.go
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

type AuthorHandler struct {
	hr      *server.Handler
	service *service.AuthorService
}

func NewAuthorHandler(handler *server.Handler, authorService *service.AuthorService) *AuthorHandler {
	return &AuthorHandler{hr: handler, service: authorService}
}

func (h *AuthorHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootAuthor)
	grp.POST("", h.hr.AuthAccess(), h.create)
	grp.GET("", h.getList)
	grp.GET("/:id", h.getByID)
	grp.PUT("/:id", h.hr.AuthAccess(), h.update)
	grp.DELETE("/:id", h.hr.AuthAccess(), h.delete)
}

// create godoc
//
// @Summary Create an author
// @Description Create a new author.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param detail body dto.AuthorDTO true "Author's detail"
// @Success 201 {object} dto.SuccessResponse[any]
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors [post]
func (h *AuthorHandler) create(c *gin.Context) {
	var req dto.AuthorDTO
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
		Message: "Author created successfully",
	})
}

// getList godoc
//
// @Summary Get a list of authors
// @Description Get a list of all authors.
// @Produce json
// @Success 200 {object} dto.SuccessResponse[[]dto.AuthorResp]
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors [get]
func (h *AuthorHandler) getList(c *gin.Context) {
	data, err := h.service.GetList()
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.AuthorResp]{
		Success: true,
		Message: "List of authors",
		Data:    data,
	})
}

// getByID godoc
//
// @Summary Get an author's detail
// @Description Get details of a specific author by ID.
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} dto.SuccessResponse[dto.AuthorResp]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors/{id} [get]
func (h *AuthorHandler) getByID(c *gin.Context) {
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

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.AuthorResp]{
		Success: true,
		Message: "Author details",
		Data:    data,
	})
}

// update godoc
//
// @Summary Update an author's detail
// @Description Update details of a specific author by ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Author ID"
// @Param detail body dto.AuthorUpdate true "Author's updated detail"
// @Success 200 {object} dto.SuccessResponse[any]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors/{id} [put]
func (h *AuthorHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
		return
	}

	var req dto.AuthorUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.hr.BindingError(err)
		return
	}
	req.ID = uint(id)

	err = h.service.Update(&req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Author updated successfully",
	})
}

// delete godoc
//
// @Summary Delete an author
// @Description Delete a specific author by ID.
// @Produce json
// @Security BearerAuth
// @Param id path int true "Author ID"
// @Success 200 {object} dto.SuccessResponse[any]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors/{id} [delete]
func (h *AuthorHandler) delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
		return
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Author deleted successfully",
	})
}

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

type BorrowingHandler struct {
	hr      *server.Handler
	service *service.BorrowingService
}

func NewBorrowingHandler(handler *server.Handler, borrowingService *service.BorrowingService) *BorrowingHandler {
	return &BorrowingHandler{hr: handler, service: borrowingService}
}

func (h *BorrowingHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootBorrowing)
	grp.POST("", h.hr.AuthAccess(), h.create)
	grp.GET("", h.getList)
	grp.GET("/:id", h.getByID)
	grp.PUT("/:id", h.hr.AuthAccess(), h.update)
	grp.DELETE("/:id", h.hr.AuthAccess(), h.delete)
}

// create godoc
//
// @Summary Create a borrowing
// @Description Create a new borrowing record.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param detail body dto.BorrowingDTO true "Borrowing's detail"
// @Success 201 {object} dto.SuccessResponse[any]
// @Failure 401 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrowings [post]
func (h *BorrowingHandler) create(c *gin.Context) {
	var req dto.BorrowingDTO
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
		Message: "Borrowing created successfully",
	})
}

// getList godoc
//
// @Summary Get a list of borrowings
// @Description Get a list of all borrowing records.
// @Produce json
// @Success 200 {object} dto.SuccessResponse[[]dto.BorrowingResp]
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrowings [get]
func (h *BorrowingHandler) getList(c *gin.Context) {
	data, err := h.service.GetList()
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.BorrowingResp]{
		Success: true,
		Message: "List of borrowings",
		Data:    data,
	})
}

// getByID godoc
//
// @Summary Get a borrowing's detail
// @Description Get details of a specific borrowing by ID.
// @Produce json
// @Param id path int true "Borrowing ID"
// @Success 200 {object} dto.SuccessResponse[dto.BorrowingResp]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrowings/{id} [get]
func (h *BorrowingHandler) getByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
		return
	}

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, exception.ErrDataNotFound) {
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		} else {
			h.hr.ErrorInternalServer(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.BorrowingResp]{
		Success: true,
		Message: "Borrowing details",
		Data:    data,
	})
}

// update godoc
//
// @Summary Update a borrowing's detail
// @Description Update details of a specific borrowing by ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Borrowing ID"
// @Param detail body dto.BorrowingUpdate true "Borrowing's updated detail"
// @Success 200 {object} dto.SuccessResponse[any]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrowings/{id} [put]
func (h *BorrowingHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
		return
	}

	var req dto.BorrowingUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.hr.BindingError(err)
		return
	}
	req.ID = uint(id)

	err = h.service.Update(&req)
	if err != nil {
		if errors.Is(err, exception.ErrDataNotFound) {
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		} else {
			h.hr.ErrorInternalServer(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Borrowing updated successfully",
	})
}

// delete godoc
//
// @Summary Delete a borrowing
// @Description Delete a specific borrowing by ID.
// @Produce json
// @Security BearerAuth
// @Param id path int true "Borrowing ID"
// @Success 200 {object} dto.SuccessResponse[any]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrowings/{id} [delete]
func (h *BorrowingHandler) delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
		return
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		if errors.Is(err, exception.ErrDataNotFound) {
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		} else {
			h.hr.ErrorInternalServer(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Borrowing deleted successfully",
	})
}

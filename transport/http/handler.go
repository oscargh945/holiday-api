package http

import (
	"encoding/xml"
	nethttp "net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oscargh945/holiday-api/domain/entities"
	"github.com/oscargh945/holiday-api/domain/usecase"
)

type HolidayHandler struct {
	useCase *usecase.HolidayUseCase
}

func NewHolidayHandler(useCase *usecase.HolidayUseCase) *HolidayHandler {
	return &HolidayHandler{
		useCase: useCase,
	}
}

type HolidayListResponse struct {
	XMLName  xml.Name           `json:"-" xml:"holidays"`
	Holidays []entities.Holiday `json:"holidays" xml:"holiday"`
}

type ErrorResponse struct {
	Message string `json:"message" xml:"message"`
}

func (h *HolidayHandler) List(c *gin.Context) {
	filter := entities.HolidayFilter{
		Type: c.Query("type"),
		From: c.Query("from"),
		To:   c.Query("to"),
	}

	holidays, err := h.useCase.List(filter)
	if err != nil {
		h.respond(c, nethttp.StatusBadRequest, ErrorResponse{
			Message: "invalid date format, expected YYYY-MM-DD",
		})
		return
	}

	h.respond(c, nethttp.StatusOK, HolidayListResponse{
		Holidays: holidays,
	})
}

func (h *HolidayHandler) respond(c *gin.Context, status int, payload any) {
	accept := c.GetHeader("Accept")

	if strings.Contains(accept, "application/xml") {
		c.XML(status, payload)
		return
	}

	c.JSON(status, payload)
}

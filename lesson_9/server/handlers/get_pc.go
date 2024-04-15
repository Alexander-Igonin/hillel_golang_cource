package handlers

import (
	computer "hillel_go_cource/lesson_9/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetResponcePcBody struct {
	Pc *computer.Computer
}

func (h *Handlers) GetPcHandler(c echo.Context) error {
	var pc GetResponcePcBody
	pc.Pc = h.Pc
	
	return c.JSON(http.StatusOK, pc)
}

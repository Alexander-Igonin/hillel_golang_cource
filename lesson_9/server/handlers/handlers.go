package handlers

import (
	computer "hillel_go_cource/lesson_9/models"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Pc *computer.Computer
}


func NewHandlers(pc *computer.Computer) *Handlers {
	return &Handlers{
		Pc: pc,
	}
}

func (h Handlers) RegisterRouts(e *echo.Echo) {
	e.GET("/", h.GetPcHandler)
	e.POST("/new", h.PostNewPcHandler)
	e.PUT("/new/hardware", h.PutNewHardwareHandler)
	e.DELETE("/delete/hardware/", h.DeleteHardwareHandler)
}

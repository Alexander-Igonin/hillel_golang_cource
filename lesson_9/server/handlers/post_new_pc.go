package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostNewPcRequestBody struct {
	Cpu string
	Gpu string
	Motherboard string
	Ram int
	Ssd int
	PowerBlock int
	Water bool
}


func (h *Handlers) PostNewPcHandler(c echo.Context) error {
	
	if err := c.Bind(&h.Pc); err != nil {
		return fmt.Errorf("post request parse error: %w", err)
	}

	return c.JSON(http.StatusOK, &h.Pc)
}

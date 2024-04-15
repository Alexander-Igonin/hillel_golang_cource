package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) PostNewPcHandler(c echo.Context) error {

	if err := c.Bind(&h.Pc); err != nil {
		return fmt.Errorf("post request parse error: %w", err)
	}

	return c.JSON(http.StatusOK, h.Pc)
}

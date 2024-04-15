package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PutNewHardwareRequestBody struct {
	Name string
	Value string
}

func (h *Handlers) PutNewHardwareHandler(c echo.Context) error {
	var hardware PutNewHardwareRequestBody
	if err := c.Bind(&hardware); err != nil {
		return fmt.Errorf("put request parse error: %w", err)
	}

	switch hardware.Name {
	case "gpu":
		h.Pc.Gpu = hardware.Value
	case "cpu":
		h.Pc.Cpu = hardware.Value
	case "motherboard":
		h.Pc.Motherboard = hardware.Value
	case "ram":
		h.Pc.Ram = hardware.Value
	case "ssd":
		h.Pc.Ssd = hardware.Value
	case "powerBlock":
		h.Pc.PowerBlock = hardware.Value
	case "water":
		h.Pc.Water = hardware.Value
	}

	return c.JSON(http.StatusOK, &h.Pc)
}

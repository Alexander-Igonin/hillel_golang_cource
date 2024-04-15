package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) DeleteHardwareHandler(c echo.Context) error {
	param := c.QueryParam("name")

	switch param {
	case "gpu":
		h.Pc.Gpu = ""
	case "cpu":
		h.Pc.Cpu = ""
	case "motherboard":
		h.Pc.Motherboard = ""
	case "ram":
		h.Pc.Ram = ""
	case "ssd":
		h.Pc.Ssd = ""
	case "powerBlock":
		h.Pc.PowerBlock = ""
	case "water":
		h.Pc.Water = ""
	}

	return c.JSON(http.StatusOK, &h.Pc)
}

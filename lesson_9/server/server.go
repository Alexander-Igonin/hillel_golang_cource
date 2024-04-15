package server

import (
	computer "hillel_go_cource/lesson_9/models"
	"hillel_go_cource/lesson_9/server/handlers"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e echo.Echo
	handlers handlers.Handlers
}

func NewServer(pc *computer.Computer) *Server {
	return &Server{
		e: *echo.New(),
		handlers: *handlers.NewHandlers(pc),
	}
}

package main

import (
	"hillel_go_cource/lesson_9/config"
	computer "hillel_go_cource/lesson_9/models"
	"hillel_go_cource/lesson_9/server/handlers"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Computer struct {
	Cpu         string
	Gpu         string
	Motherboard string
	Ram         int
	Ssd         int
	PowerBlock  int
	Water       bool
}

func main() {

	pc := computer.NewComputer(
		"amd ryzen 5600",
		"amd rx 6650",
		"MSI B450M",
		"16",
		"512",
		"700",
		"no",
	)
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	server := echo.New()
	handlers := handlers.NewHandlers(pc)
	handlers.RegisterRouts(server)

	err = server.Start(cfg.Port)
	if err != nil {
		logrus.Fatal(err)
	}
}

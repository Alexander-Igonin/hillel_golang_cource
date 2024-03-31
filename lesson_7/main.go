package main

import (
	"flag"
	"fmt"
	"hillel_go_cource/lesson_7/tcpserver"
	"hillel_go_cource/lesson_7/webserver"
)

func main() {
	startServer := flag.String("server", "web", "Choose web or tcp server")
	flag.Parse()

	switch *startServer {
	case "web":
		err := webserver.Start()
		if err != nil {
			fmt.Println("Failed to run web server")
		}
	case "tcp":
		err := tcpserver.Start()
		if err != nil {
			fmt.Println("Failed to run tcp server")
		}
	}




}

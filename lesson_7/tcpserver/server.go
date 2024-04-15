package tcpserver

import (
	"errors"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

func Start() error {
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		return errors.New("failed to up tcp server")
	}

	defer listener.Close()

	fmt.Println("TCP Server start on port 5050")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		logrus.Info("somebody connected")

		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) error {
	tmp := make([]byte, 256)
	defer conn.Close()

	_, err := conn.Read(tmp)
	if err != nil {
		logrus.Info("failed to read ", err)
		return err
	}

	fmt.Println(string(tmp))

	return nil
}

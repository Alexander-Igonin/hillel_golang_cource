package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type HandlePostRequest struct {
	Name string `json:"name"`
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error("Wrong post request")
		return
	}

	req := HandleGetResponse{}

	err = json.Unmarshal(reqBytes, &req)
	if err != nil {
		logrus.Error("Failed to unmarshal post request")
		return
	}
	fmt.Println(req)

	io.WriteString(w, "Done")
}

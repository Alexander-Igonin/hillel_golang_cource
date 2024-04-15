package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type HandleGetResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	id := query.Get("id")
	name := query.Get("name")

	resp := HandleGetResponse{
		Id: id,
		Name: name,
	}

	marshResp, err := json.Marshal(resp)
	if err != nil {
		logrus.Error("Failed marshal get response")
		return
	} 

	w.Write(marshResp)
}

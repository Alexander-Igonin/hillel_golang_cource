package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Dog struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func HandlePutRequest(w http.ResponseWriter, r *http.Request, d *Dog) (Dog, error) {
	result := d
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		return *result, errors.New("failed to read put body")
	}

	err = json.Unmarshal(reqBody, &result)
	if err != nil {
		return *result, errors.New("failed to read put body")
	}

	return *result, nil
}

package handlers

import (
	"errors"
	"net/http"
	"strconv"
)

var Films = make(map[string]string)

func HandleDeleteRequest(w http.ResponseWriter, r *http.Request, m *map[string]string) (map[string]string, error) {
	quaery := r.URL.Query()

	index := quaery.Get("index")
	indexInt, err := strconv.Atoi(index)
	if err != nil || indexInt < 0 || indexInt > len(*m) {
		return *m, errors.New("wrong index")
	}

	delete(*m, index)

	return *m, nil
}

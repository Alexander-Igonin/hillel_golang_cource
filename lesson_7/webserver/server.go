package webserver

import (
	"errors"
	"fmt"
	"hillel_go_cource/lesson_7/webserver/handlers"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Start() error {
	http.HandleFunc("/", Handle)

	fmt.Println("web Server start on port 5050")
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		return errors.New("there is an error")
	}

	return nil
}

func Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.HandleGet(w, r)
	case http.MethodPost:
		handlers.HandlePost(w, r)
	case http.MethodPut:
		dog := handlers.Dog{
			Name: "Garry",
			Age: 3,
		}

		fmt.Println("Before request: ", dog)

		dogAfterRequest, err := handlers.HandlePutRequest(w, r, &dog)
		if err != nil {
			logrus.Error("Wrong put request")
		}

		fmt.Println("After request: ", dogAfterRequest)
	case http.MethodDelete:
		films := handlers.Films
		films["1"] = "IronMan"
		films["2"] = "SpiderMan"

		fmt.Println("Map before delete", films)

		_, err := handlers.HandleDeleteRequest(w, r, &films)
		if err != nil {
			logrus.Error("Wrong index")
		}

		fmt.Println("Map after delete", films)
	}
}

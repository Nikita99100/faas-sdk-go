package framework

import (
	"net/http"
)

func Start(port string, function http.HandlerFunc) error {
	http.HandleFunc("/", function)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return err
	}
	return nil
}

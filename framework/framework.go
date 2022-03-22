package framework

import (
	"net/http"
)

func Start(port string, function http.HandlerFunc) {
	http.HandleFunc("/", function)
	http.ListenAndServe(":"+port, nil)
}

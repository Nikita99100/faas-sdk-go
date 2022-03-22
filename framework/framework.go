package framework

import (
	"net/http"
)

func Start(port string, function http.HandlerFunc) error {
	http.HandleFunc("/", function)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) })
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return err
	}
	return nil
}

// Package handlers
package handlers

import "net/http"

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	data := "Hello world"
	if r.Method == http.MethodPost {
		_, err := w.Write([]byte(data))
		if err != nil {
		}
	} else {
		http.Error(w, "Method no valid", http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-type", "Application/json")
}

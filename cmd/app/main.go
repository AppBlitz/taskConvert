package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/AppBlitz/taskConvert/internal/handlers"
)

func main() {
	servers := http.NewServeMux()
	servers.HandleFunc("/length", handlers.LengthHandler)
	erro := http.ListenAndServe(":3333", servers)
	if errors.Is(erro, http.ErrServerClosed) {
		fmt.Printf("%s\n", "Serve closed")
	} else if erro != nil {
		fmt.Printf("error starting server: %s\n", erro)
		os.Exit(1)
	}
}

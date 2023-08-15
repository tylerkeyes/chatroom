package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/tylerkeyes/chatroom/pkg/routes"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", routes.GetRoot)
    mux.HandleFunc("/hello", routes.GetHello)
    
    err := http.ListenAndServe(":3333", mux)
    if errors.Is(err, http.ErrServerClosed) {
        fmt.Println("server closed")
    } else if err != nil {
        fmt.Printf("error starting server: %s\n", err)
        os.Exit(1)
    }
}

package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// globals
var router *chi.Mux

// TODO: create global variable for db access

func main() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	server := Server{}
	server.router = chi.NewRouter()
	server.router.Use(middleware.Recoverer)
	server.routes()

	var err error

	/*
	router.Use(ChangeMethod)
	router.Get("/", server.GetRoot())
	router.Post("/", CreateRoom)
	router.Route("/room/{roomId}", func(r chi.Router) {
		r.Get("/", GetRoomMessages)
		r.Post("/", SendRoomMessage)
	})
	*/

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", routes.GetRoot)
	//mux.HandleFunc("/hello", routes.GetHello)
	//mux.HandleFunc("/room/", routes.GetRoomMessages)

	err = http.ListenAndServe(":3333", router)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

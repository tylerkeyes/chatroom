package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/gocql/gocql"
)

type Server struct {
	router *chi.Mux
	db     *gocql.Session
}

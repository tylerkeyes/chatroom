package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type message struct {
	id   uint64
	chat string
	room string
	user string
}

func (s *Server) routes() {
	fmt.Println("Initialize routes")
	s.router.Use(ChangeMethod)
	s.router.Get("/", s.GetRoot())
	s.router.Post("/", s.CreateRoom())
	s.router.Route("/room/{roomId}", func(r chi.Router) {
		r.Get("/", s.GetRoomMessages())
		r.Post("/", s.SendRoomMessage())
	})
}

func (s *Server) GetRoot() http.HandlerFunc {
	fmt.Println("Get root handler")
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("In return func")
		hasFirst := r.URL.Query().Has("first")
		first := r.URL.Query().Get("first")
		hasSecond := r.URL.Query().Has("second")
		second := r.URL.Query().Get("second")

		fmt.Println("Read body")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("counld not read request body: %s\n", err)
		}

		fmt.Printf("got / request, first(%t)=%s, second(%t)=%s, body:\n%s\n",
			hasFirst, first, hasSecond, second, body)
		http.ServeFile(w, r, "./templates/index.html")
		io.WriteString(w, "Welcome to chatbot\n")
	}
}

func (s *Server) GetHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got /hello request\n")

		myName := r.PostFormValue("myName")
		if myName == "" {
			w.Header().Set("x-missing-field", "myName")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		io.WriteString(w, fmt.Sprintf("Hello, %s!\n", myName))
	}
}

func (s *Server) CreateRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got create room request\n")

		roomName := r.FormValue("roomName")
		if roomName == "" {
			w.Header().Set("x-missing-field", "roomName")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Printf("Room name: %s\n", roomName)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		// TOTO: add room to database, if already exists return warning
	}
}

// Route handler for chat rooms
func (s *Server) GetRoomMessages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomName := strings.TrimPrefix(r.URL.Path, "/room/")

		// TOTO: get chat history data from database for given room, read into messages slice
		messages := make([]message, 0)
		messages = append(messages, message{id: 1, chat: "hello", room: "myroom", user: "tkeyes"})

		tmpl, _ := template.ParseFiles("templates/test.html") //New("itemlist").

		input := map[string]interface{}{}
		input["room_name"] = roomName
		input["messages"] = messages

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "text/html")
		err := tmpl.Execute(w, input)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) SendRoomMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//roomName := strings.TrimPrefix(r.URL.Path, "/room/")
		// TODO: get all messages from room
	}
}

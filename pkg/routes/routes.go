package routes

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
    hasFirst := r.URL.Query().Has("first")
    first := r.URL.Query().Get("first")
    hasSecond := r.URL.Query().Has("second")
    second := r.URL.Query().Get("second")

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Printf("counld not read request body: %s\n", err)
    }

    fmt.Printf("got / request, first(%t)=%s, second(%t)=%s, body:\n%s\n",
        hasFirst, first, hasSecond, second, body)
    io.WriteString(w, "Welcome to chatbot\n")
}

func GetHello(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got /hello request\n")
    io.WriteString(w, "Hello!\n")
}

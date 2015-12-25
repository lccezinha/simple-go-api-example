package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
)

type Todo struct {
    Name        string
    Completed   bool
    Due         time.Time
}

type Todos []Todo // Cria um novo tipo "Todos", que é um slice de Todo.

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "TodoIndex!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // Pega os paramêtros enviados na requisição
    todoId := params["todoId"]
    fmt.Fprintln(w, "TodoShow: ", todoId)
}
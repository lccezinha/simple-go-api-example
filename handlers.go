package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Estuda Golang"},
        Todo{Name: "Zerar Dead Space"},
    }

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // Pega os paramêtros enviados na requisição
    todoId := params["todoId"]
    fmt.Fprintln(w, "TodoShow: ", todoId)
}
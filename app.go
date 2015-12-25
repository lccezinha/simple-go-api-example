package main

import (
    "fmt"
    "log"
    "net/http"
    "html"
    "time"
    "encoding/json"

    "github.com/gorilla/mux"
)

type Todo struct {
    Name        string      `json:"name"`
    Completed   bool        `json:"completed"`
    Due         time.Time   `json:"due"`
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
    todos := Todos{
        Todo{Name: "Estuda Golang"},
        Todo{Name: "Zerar Dead Space"},
    }

    json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // Pega os paramêtros enviados na requisição
    todoId := params["todoId"]
    fmt.Fprintln(w, "TodoShow: ", todoId)
}
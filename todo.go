package main

import "time"

type Todo struct {
    Name        string      `json:"name"`
    Completed   bool        `json:"completed"`
    Due         time.Time   `json:"due"`
}

type Todos []Todo // Cria um novo tipo "Todos", que é um slice de Todo.
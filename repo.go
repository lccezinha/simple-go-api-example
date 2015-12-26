package main

import "fmt"

var currentId int
var todos Todos

func init() {
    RepoCreateTodo(Todo{Name: "Estudar Golang"})
    RepoCreateTodo(Todo{Name: "Zerar Dead Space"})
}

func generateId() int {
    currentId += 1
    return currentId
}

func RepoFindTodo(id int) Todo {
    for _, t := range todos {
        if t.Id == id {
            return t
        }
    }

    return Todo{}
}

func RepoCreateTodo(todo Todo) Todo {
    todo.Id = generateId()
    todos = append(todos, todo)

    return todo
}

func RepoDestroyTodo(id int) error {
    for i, t := range todos {
        if t.Id == id {
            todos = append(todos[:i], todos[i+1:]...)
            return nil
        }
    }

    return fmt.Errorf("Não encontrou o ID %d do Todo para remover", id)
}
package main

import (
	"lesson7/api"

	"lesson7/todo"
)

func main() {
	todo.Justdoit()
	api.InitRouter_gin()
}

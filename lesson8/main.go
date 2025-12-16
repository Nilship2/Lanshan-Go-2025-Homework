package main

import (
	"lesson7/api"

	"lesson7/todo"
)

func main() {
	//永生redis，颗秒缓存击穿和雪崩
	todo.Justdoit()
	api.InitRouter_gin()
}

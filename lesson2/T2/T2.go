package main

import (
	"fmt"
)

func Fac(c string) func(int, int) int {
	if c == "+" {
		return func(x int, y int) int {
			return x + y
		}
	}
	if c == "-" {
		return func(x int, y int) int {
			return x - y
		}
	}
	if c == "*" {
		return func(x int, y int) int {
			return x * y
		}
	}
	if c == "/" {
		return func(x int, y int) int {
			return x / y
		}
	}
	return func(x int, y int) int {
		return 114514
	}
}

func main() {
	Add := Fac("+")
	Mi := Fac("-")
	Mu := Fac("*")
	Di := Fac("/")

	a := 1
	b := 2
	fmt.Scan(a)
	fmt.Scan(b)
	fmt.Println("a+b=", Add(a, b))
	fmt.Println("a-b=", Mi(a, b))
	fmt.Println("a*b=", Mu(a, b))
	fmt.Println("a/b=", Di(a, b))

}

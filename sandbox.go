package main

import (
	"fmt"
	"time"
)

func fun1(abc string) {
	fmt.Println(abc)
}

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
	fun1("hello")
}

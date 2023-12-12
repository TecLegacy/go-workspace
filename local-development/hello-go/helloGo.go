package main

import (
	"fmt"

	"github.com/teclegacy/test-go"
)

func main() {
	fmt.Println("hello-go")
	val := testgo.Add(10, 20)
	fmt.Println(val)
}

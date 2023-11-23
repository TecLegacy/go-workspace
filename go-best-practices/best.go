package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println(err)
	}
}

func test() error {
	err := errors.New("Something went wrong")
	return err
}

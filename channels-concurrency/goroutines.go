package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 10
	}()
	val := <-ch
	fmt.Println(val)

	chz := normalCh("keshav")
	receiverChannel(chz)
}

// * first channel
func normalCh(name string) chan string {
	ch := make(chan string, 1)

	ch <- name

	return ch

}

func receiverChannel(ch chan string) {
	val := <-ch
	fmt.Println(val)
}

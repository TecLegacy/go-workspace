package main

import "fmt"

func main() {
	ch := normalCh("keshav")

	receiverChannel(ch)
}

// * first channel
func normalCh(name string) chan string {
	ch := make(chan string, 2)

	ch <- name

	return ch

}

func receiverChannel(ch chan string) {
	val := <-ch
	fmt.Println(val)
}

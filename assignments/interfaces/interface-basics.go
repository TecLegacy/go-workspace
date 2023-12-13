package main

import "fmt"

func main() {
	val := plusOne{
		val: 50,
	}
	check := val.addOne(1)
	fmt.Println(check)

	val.decree(4)
}

type plusOne struct {
	val int
}

func (p plusOne) decree(val int) {
	p.val -= val
	fmt.Println(p.val)
}

type increase interface {
	addOne(value int) int
}

func (p plusOne) addOne(val int) int {
	p.val += val
	return p.val
}

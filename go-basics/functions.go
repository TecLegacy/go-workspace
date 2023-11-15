package main

import "fmt"

func main() {

	//struct invoke
	addToDB(Person{
		name: "keshav",
		age:  20,
	})

	/* Pass by value for variables
	var value int
	value = add(10,20)
	fmt.Println(value)
	fmt.Printf(value)
	fmt.Println()
	*/

}

// Pass by Value for variables
func add(x int, y int) int {
	return x + y
}

// Structs
type Person struct {
	name string
	age  int
}

type Home struct {
}

func addToDB(p Person) {
	fmt.Printf("Your Name is %s And Your age %d", p.name, p.age)
}

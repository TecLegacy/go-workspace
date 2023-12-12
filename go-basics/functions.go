package main

import "fmt"

func main() {

  currentColor:=car{
    color: "blue",
  }

  currentColor.changeColor("red")
  
  fmt.Println(currentColor.color)

  
}

//pointer receiver 
type car struct{
  color string
}

func (c *car) changeColor(color string){
  c.color = color
}

// func main() {

// 	//struct invoke
// 	addToDB(Person{
// 		name: "keshav",
// 		age:  20,
// 	})

// 	OwnerDetails(Home{
// 		Landmark: "Delhi",
// 		AreaCode: 30,
// 		Owner: Person{
// 			name: "Kanu",
// 			age:  26,
// 		},
// 	})

	/* Pass by value for variables
	var value int
	value = add(10,20)
	fmt.Println(value)
	fmt.Printf(value)
	fmt.Println()
	*/

}

// Pass by Value for variables
// func add(x int, y int) int {
// 	return x + y
// }

// Structs
type Person struct {
	name string
	age  int
}

type Home struct {
	Landmark string
	AreaCode int
	Owner    Person
}

func addToDB(p Person) {
	fmt.Printf("Your Name is %s And Your age %d \n", p.name, p.age)
}

func OwnerDetails(records Home) {
	fmt.Printf("Your area code %d \n", records.AreaCode)
	fmt.Printf("Your Owner Age is %d & name is %s", records.Owner.age, records.Owner.name)
}

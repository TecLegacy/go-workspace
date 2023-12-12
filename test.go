package main

import (
	"fmt"
)

type Shape interface {
	calculate() float64
}
type Circle struct {
	radius int
}

func (c Circle) calculate() float64 {
	return float64(c.radius)
}

func test(s Shape) {
	c, ok := s.(Circle)
	if !ok {
		fmt.Println("fatal panic")
	}
	// fmt.Printf("vals %v", fuk.radius)
	fmt.Printf("vals %T", c)
}

func main() {

	some := make(map[string]int)
	some["d"] = 1

	do := map[string]int{
		"asd": 1,
	}

	test(Circle{
		radius: 20,
	})

	test1(dog{
		breed: "german",
		age:   10,
		trait: Trait{
			sound:  "woof",
			energy: 10,
		},
		wearables: struct {
			pawCloth  string
			tailCloth string
		}{
			pawCloth:  "wool",
			tailCloth: "wool",
		},
	})
}

func test1(d dog) {
	fmt.Printf("dog bread is %s \n"+
		"is %d old \n"+
		"he likes to %s & is always full of %d \n "+
		"his fav cloths are % s & for the tail he prefers %s \n", d.breed, d.age, d.trait.sound, d.trait.energy, d.wearables.pawCloth, d.wearables.tailCloth)

}

type dog struct {
	breed     string
	age       int
	trait     Trait
	wearables struct {
		pawCloth  string
		tailCloth string
	}
}

type Trait struct {
	sound  string
	energy int
}

//#HEADING: Interfaces & Struct
// func main() {
// 	test(morning{
// 		message: "First Report",
// 		hours:   time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
// 	})
// 	test(health{
// 		doctorName: "John Doe",
// 		visits:     10,
// 	})
// }

// func test(m sendMessage) {
// 	fmt.Printf(m.showMessages())
// }

// // Show morning messages or doctor visits
// type sendMessage interface {
// 	showMessages() string
// }
// type morning struct {
// 	message string
// 	hours   time.Time
// }
// type health struct {
// 	doctorName string
// 	visits     int
// }

// func (h health) showMessages() string {
// 	return fmt.Sprintf("Your doctors name is %s and this is your %d visit \n", h.doctorName, h.visits)
// }

// func (m morning) showMessages() string {
// 	return fmt.Sprintf("%s . its a new day %s \n", m.message, m.hours.Format(time.RFC3339))
// }

/*
#TODO: Refactor to reuse interfaces on struct
func message(m morning) {
	fmt.Printf("your morning message is %s which came at time %s", m.message, m.hours.Format(time.RFC3339))
}
*/

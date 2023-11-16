package main

import (
	"fmt"
	"time"
)

func main() {
	test(morning{
		message: "First Report",
		hours:   time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test(health{
		doctorName: "John Doe",
		visits:     10,
	})
}

func test(m sendMessage) {
	fmt.Printf(m.showMessages())
}

// Show morning messages or doctor visits
type sendMessage interface {
	showMessages() string
}
type morning struct {
	message string
	hours   time.Time
}
type health struct {
	doctorName string
	visits     int
}

func (h health) showMessages() string {
	return fmt.Sprintf("Your doctors name is %s and this is your %d visit \n", h.doctorName, h.visits)
}
func (m morning) showMessages() string {
	return fmt.Sprintf("%s . its a new day %s \n", m.message, m.hours.Format(time.RFC3339))
}

/*
#TODO: Refactor to reuse interfaces on struct
func message(m morning) {
	fmt.Printf("your morning message is %s which came at time %s", m.message, m.hours.Format(time.RFC3339))
}
*/

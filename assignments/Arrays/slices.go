package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	cost := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		messages := messages[i]
		cost[i] := len(messages) * 0.01
	}
	return cost
}

// don't edit below this line

func test(messages []string) {
	costs := getMessageCosts(messages)
	fmt.Println("Messages:")
	for i := 0; i < len(messages); i++ {
		fmt.Printf(" - %v\n", messages[i])
	}
	fmt.Println("Costs:")
	for i := 0; i < len(costs); i++ {
		fmt.Printf(" - %.2f\n", costs[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})
	test([]string{
		"I don't want to be here anymore",
		"Can we go home?",
		"I'm hungry",
		"I'm bored",
	})
	test([]string{
		"Hello",
		"Hi",
		"Hey",
		"Hi there",
		"Hey there",
		"Hi there",
		"Hello there",
		"Hey there",
		"Hello there",
		"General Kenobi",
	})
}

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// const (
// 	planFree = "free"
// 	planPro  = "pro"
// )

// /*
// *Complete the getMessageWithRetriesForPlan function. It takes a plan variable as input, and you've been provided with constants for the plan types at the top of the program.

// *If the plan is a pro plan, return all the strings from getMessageWithRetries().
// *If the plan is a free plan, return the first 2 strings from getMessageWithRetries().
// *If the plan isn't either of those, return an error that says unsupported plan.
//  */
// func getMessageWithRetriesForPlan(plan string) ([]string, error) {
// 	allMessages := getMessageWithRetries()
// 	if plan == planPro {
// 		return allMessages[1:], nil
// 	}
// 	if plan == planFree {
// 		return allMessages[1:3], nil
// 	}
// 	return allMessages[:], errors.New("unsupported plan")
// }

// // don't touch below this line

// func getMessageWithRetries() [3]string {
// 	return [3]string{
// 		"click here to sign up",
// 		"pretty please click here",
// 		"we beg you to sign up",
// 	}
// }

// func test(name string, doneAt int, plan string) {
// 	defer fmt.Println("=====================================")
// 	fmt.Printf("sending to %v...", name)
// 	fmt.Println()

// 	messages, err := getMessageWithRetriesForPlan(plan)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	for i := 0; i < len(messages); i++ {
// 		msg := messages[i]
// 		fmt.Printf(`sending: "%v"`, msg)
// 		fmt.Println()
// 		if i == doneAt {
// 			fmt.Println("they responded!")
// 			break
// 		}
// 		if i == len(messages)-1 {
// 			fmt.Println("no response")
// 		}
// 	}
// }

// func main() {
// 	test("Ozgur", 3, planFree)
// 	test("Jeff", 3, planPro)
// 	test("Sally", 2, planPro)
// 	test("Sally", 3, "no plan")
// }

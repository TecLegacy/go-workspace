// * Custom Errors
package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (d divideError) Error() string {
	return fmt.Sprintf("Dividend cannot be %f by zero", d.dividend)
}

// don't edit below this line

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		// We convert the `divideError` struct to an `error` type by returning it
		// as an error. As an error type, when it's printed its default value
		// will be the result of the Error() method
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func test(dividend, divisor float64) {
	defer fmt.Println("====================================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", dividend, divisor)
	quotient, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

func main() {
	test(10, 0)
	test(10, 2)
	test(15, 30)
	test(6, 3)
	test(0, 10)
}

//* HEADING: Basic error handling
// package main

// import (
// 	"fmt"
// )

// func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
// 	customerSMS, err := sendSMS(msgToCustomer)
// 	if err != nil {
// 		return 0.0, err
// 	}
// 	spouseSMS, err := sendSMS(msgToSpouse)
// 	if err != nil {
// 		return 0.0, err
// 	}

// 	return customerSMS + spouseSMS, nil
// }

// // don't edit below this line

// func sendSMS(message string) (float64, error) {
// 	const maxTextLen = 25
// 	const costPerChar = .0002
// 	if len(message) > maxTextLen {
// 		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
// 	}
// 	return costPerChar * float64(len(message)), nil
// }

// func test(msgToCustomer, msgToSpouse string) {
// 	defer fmt.Println("========")
// 	fmt.Println("Message for customer:", msgToCustomer)
// 	fmt.Println("Message for spouse:", msgToSpouse)
// 	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Printf("Total cost: $%.4f\n", totalCost)
// }

// func main() {
// 	test(
// 		"Thanks for coming in to our flower shop today!",
// 		"We hope you enjoyed your gift.",
// 	)
// 	test(
// 		"Thanks for joining us!",
// 		"Have a good day.",
// 	)
// 	test(
// 		"Thank you.",
// 		"Enjoy!",
// 	)
// 	test(
// 		"We loved having you in!",
// 		"We hope the rest of your evening is absolutely fantastic.",
// 	)
// }

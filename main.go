package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b Bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose an option: (a) - add item, (s) - save bill, (t) - add tip, (q) - quit: ", reader)

	switch option {
	case "a":
		user_item, _ := getInput("What is the name of the item? ", reader)
		user_cost, _ := getInput("What is the item price? ", reader)

		// convert string to float
		price, err := strconv.ParseFloat(user_cost, 64)
		if err != nil {
			fmt.Println("This price must be a number")
			promptOptions(b)
		}
		b.addItem(user_item, price)
		promptOptions(b)

	case "s":
		b.saveBill()
	case "t":
		user_tip, _ := getInput("Enter a tip amount: ", reader)
		// convert string to float
		tip, err := strconv.ParseFloat(user_tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(tip)
		promptOptions(b)
	case "q":
		break
	default:
		fmt.Println("Invalid Option")
		promptOptions(b)

	}
}

func createBill() Bill {
	// read from standard input
	reader := bufio.NewReader(os.Stdin)

	billName, _ := getInput("Enter a Bill Name: ", reader)

	newBill := newBill(billName)

	return newBill
}

func main() {

	myBill := createBill()
	fmt.Printf("Bill created: %v", myBill)
	promptOptions(myBill)

}

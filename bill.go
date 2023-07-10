package main

import (
	"fmt"
	"os"
)

type Bill struct {
	Name  string
	Items map[string]float64
	Tip   float64
}

// make a new bill
func newBill(name string) Bill {
	bill := Bill{Name: name, Items: map[string]float64{"pie": 6.50, "coffee": 4.50}, Tip: 0}
	return bill

}

// format the bill (receiver function)
func (b *Bill) format() string {
	fs := "Bill breakdown: \n"
	totalCost := 0

	for name, cost := range b.Items {
		fs += fmt.Sprintf("Item: %-25v .... Price: $%v\n", name, cost)
		totalCost += int(cost)
	}

	totalCost += int(b.Tip)

	fs += fmt.Sprintf("Tip: %-25v \n %-41v $%0.2f", b.Tip, "Total:", float64(totalCost))
	return fs
}

// update the tip
func (b *Bill) updateTip(tip float64) {
	b.Tip = tip
}

// add item to bill
func (b *Bill) addItem(item string, cost float64) {
	b.Items[item] = cost
}

// save the bill
func (b *Bill) saveBill(){
	data := []byte(b.format())

	err := os.WriteFile("bills" + b.Name + ".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file")
}
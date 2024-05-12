package main

import (
	"fmt"
	"strings"
	"time"
)

type Item struct {
	Name     string
	Price    float64
	Quantity int
}

type ShoppingCart struct {
	Items []Item
}

func main() {
	var cart ShoppingCart
	var customerName string

	fmt.Println("WELCOME TO SEMICOLON CHECKOUT APP :)")
	fmt.Print("What is the customer's name: ")
	fmt.Scanln(&customerName)

	addItemsToCart(&cart)

	var cashierName string
	var discount float64

	fmt.Print("What is your name? ")
	fmt.Scanln(&cashierName)

	fmt.Print("How much discount will you apply?: ")
	fmt.Scanln(&discount)

	printInvoice(customerName, cashierName, cart, discount)
}

func addItemsToCart(cart *ShoppingCart) {
	for {
		var itemName string
		var price float64
		var quantity int

		fmt.Println("What did the customer buy? ")
		fmt.Scanln(&itemName)

		fmt.Println("How much per unit?")
		fmt.Scanln(&price)

		fmt.Println("How many pieces?")
		fmt.Scanln(&quantity)

		item := Item{Name: itemName, Price: price, Quantity: quantity}
		cart.Items = append(cart.Items, item)

		var addMoreItems string
		fmt.Print("Add more items? Type 'yes' or 'no': ")
		fmt.Scanln(&addMoreItems)

		if strings.ToLower(addMoreItems) != "yes" {
			break
		}
	}
}

func calculateSubtotal(cart ShoppingCart) float64 {
	var subtotal float64
	for _, item := range cart.Items {
		subtotal += item.Price * float64(item.Quantity)
	}
	return subtotal
}

func calculateDiscount(subtotal float64, discount float64) float64 {
	return (discount / 100) * subtotal
}

func calculateTotal(subtotal float64, discount float64) float64 {
	subtotalAfterDiscount := subtotal - calculateDiscount(subtotal, discount)
	vat := 0.175 * subtotal
	return subtotalAfterDiscount + vat
}

func printInvoice(customerName string, cashierName string, cart ShoppingCart, discount float64) {
	fmt.Println("================================================================")
	fmt.Println("                 SEMICOLON STORES                               ")
	fmt.Println("                 MAIN BRANCH                                   ")
	fmt.Println("         LOCATION: 321, HERBERT MACAULAY WAY, SABO, YABA, LAGOS.")
	fmt.Println("                 TELL: 03293828343                           ")
	fmt.Println("================================================================")
	fmt.Printf("Date: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("Cashier: %s\n", cashierName)
	fmt.Printf("Customer: %s\n", customerName)
	fmt.Println("================================================================")
	fmt.Println(" ITEM \t\tPRICE \t\tQTY \t\tTOTAL(NGN)")
	fmt.Println("----------------------------------------------------------------")

	subtotal := calculateSubtotal(cart)
	discountAmount := calculateDiscount(subtotal, discount)
	total := calculateTotal(subtotal, discount)

	for _, item := range cart.Items {
		fmt.Printf("%s\t\t%.2f\t\t%d\t\t%.2f\n", item.Name, item.Price, item.Quantity, item.Price*float64(item.Quantity))
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("\t\tSub Total:         %.2f\n", subtotal)
	fmt.Printf("\t\tDiscount:          %.2f\n", discountAmount)
	fmt.Printf("\t\tVAT @ 17.5%%:       %.2f\n", 0.175*subtotal)
	fmt.Println("==================================================")
	fmt.Printf("\t\tTotal Bill:        %.2f\n", total)
	fmt.Println("==================================================")
	fmt.Println("       THIS IS NOT A RECEIPT - KINDLY PAY        ")
	fmt.Println("==================================================")
}

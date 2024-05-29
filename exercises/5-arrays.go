package main

import "fmt"

type Product struct {
	price int
	name  string
}

func printStat(list [4]Product) {
	var cost, totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItems += 1
		}
	}

	fmt.Println("Last item on the list:", list[totalItems-1])
	fmt.Println("Total items", totalItems)
	fmt.Println("Total cost", cost)
}

func main() {
	shoppingList := [4]Product{
		{1, "Banana"},
		{6, "Meat"},
		{3, "Salad"},
	}

	printStat(shoppingList)
	shoppingList[3] = Product{4, "Bread"}
	printStat(shoppingList)
}

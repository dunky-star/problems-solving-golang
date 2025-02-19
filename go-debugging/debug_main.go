package main

import "fmt"

func mainDebug() {
	itemOpeniongStock := 500
	itemSold := 100
	itemReturned := 50
	itemMissing := 5
	var itemClosingStockCalc int

	itemActualClosingStock := 445

	fmt.Println("Available Inventory (Check): ", itemActualClosingStock)
	fmt.Println("--------------------------------")

	itemClosingStockCalc = finalCalc(itemOpeniongStock, itemSold, itemReturned, itemMissing)
	fmt.Println("Available Inventory (calc): ", itemClosingStockCalc)

	fmt.Println("--------------------------------")
	if itemClosingStockCalc != itemActualClosingStock {
		fmt.Println("Warning: Calculation ERROR!")

	} else {
		fmt.Println("Great Work: All Good!")
	}
}

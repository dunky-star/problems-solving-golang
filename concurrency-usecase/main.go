package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
)

func main() {
	var receivedOrderCh = make(chan order)
	var validOrderCh = make(chan order)
	var invalidOrderCh = make(chan invalidOrder)
	var wg sync.WaitGroup

	go receivedOrders(receivedOrderCh, &wg)
	go validateOrders(receivedOrderCh, validOrderCh, invalidOrderCh)

	wg.Add(1)
	go func() {
		order := <-validOrderCh
		fmt.Printf("Valid order received: %s", order)
	}()
	wg.Wait()

}

func validateOrders(in, out chan order, errChan chan invalidOrder) {
	order := <-in
	if order.Quantity <= 0 {
		// Error condition
		errChan <- invalidOrder{order: order, err: fmt.Errorf("Quantity %.2f must be grater that zero for product %s",
			order.Quantity, order.ProductCode)}
	} else {
		// Success handling
		out <- order
	}

}

func receivedOrders(out chan order, wg *sync.WaitGroup) {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		out <- newOrder
	}
	wg.Done()
}

var rawOrders = []string{
	fmt.Sprintf(`{"productCode": "%s", "quantity": 5, "status": 1}`, uuid.NewString()),
	fmt.Sprintf(`{"productCode": "%s", "quantity": 42.3, "status": 1}`, uuid.NewString()),
	fmt.Sprintf(`{"productCode": "%s", "quantity": 19, "status": 1}`, uuid.NewString()),
	fmt.Sprintf(`{"productCode": "%s", "quantity": 8, "status": 1}`, uuid.NewString()),
}

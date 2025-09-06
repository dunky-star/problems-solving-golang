package main

import (
	"encoding/json"
	"errors"
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
		wg.Done()
	}()
	go func() {
		order := <-invalidOrderCh
		fmt.Printf("Invalid order received: %s. Issue: %s", order.order, order.err)
		wg.Done()
	}()
	wg.Wait()

}

func validateOrders(in <-chan order, out chan<- order, errChan chan<- invalidOrder) {
	for order := range in {
		if order.Quantity <= 0 {
			// Error condition
			errChan <- invalidOrder{order: order, err: errors.New("Quantity must be grater that zero for product")}
		} else {
			// Success handling
			out <- order
		}
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

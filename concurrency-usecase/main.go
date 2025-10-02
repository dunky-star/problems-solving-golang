package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/google/uuid"
)

func main() {
	var wg sync.WaitGroup

	receivedOrderCh := receivedOrders()
	validOrderCh, invalidOrderCh := validateOrders(receivedOrderCh)

	wg.Add(1)
	go func(validOrderCh <-chan order, invalidOrderCh <-chan invalidOrder) {
	loop:
		for {
			select {
			case order, ok := <-validOrderCh:
				if ok {
					fmt.Printf("Valid order received: %s", order)
				} else {
					break loop
				}

			case order, ok := <-invalidOrderCh:
				if ok {
					fmt.Printf("Invalid order received: %s. Issue: %s", order.order, order.err)
				} else {
					break loop
				}

			}
		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)
	wg.Wait()
}

func validateOrders(in <-chan order) (<-chan order, <-chan invalidOrder) {
	out := make(chan order)
	errCh := make(chan invalidOrder)
	go func() {
		for order := range in {
			if order.Quantity <= 0 {
				// Error condition
				errCh <- invalidOrder{order: order, err: errors.New("quantity must be grater that zero for product")}
			} else {
				// Success handling
				out <- order
			}
		}
		close(out)
		close(errCh)
	}()
	return out, errCh
}

func receivedOrders() <-chan order {
	out := make(chan order)
	go func() {
		for _, rawOrder := range rawOrders {
			var newOrder order
			err := json.Unmarshal([]byte(rawOrder), &newOrder)
			if err != nil {
				log.Print(err)
				continue
			}
			out <- newOrder
		}
		close(out)
	}()
	return out
}

var rawOrders = []string{
	fmt.Sprintf(`{"productCode": "%s", "quantity": 5, "status": 1}`, uuid.NewString()),
	fmt.Sprintf(`{"productCode": "%s", "quantity": 42.3, "status": 1}`, uuid.NewString()),
	fmt.Sprintf(`{"productCode": "%s", "quantity": 19, "status": 1}`, uuid.NewString()),
	fmt.Sprintf(`{"productCode": "%s", "quantity": 8, "status": 1}`, uuid.NewString()),
}

// Non blocking error channel
var (
	in = make(chan string)
)

func worker(in <-chan string) (chan int, chan error) {
	out := make(chan int)
	errCh := make(chan error, 1)
	go func() {
		for msg := range in {
			i, err := strconv.Atoi(msg)
			if err != nil {
				errCh <- err
				continue
			}
			out <- i
		}
	}()
	return out, errCh
}

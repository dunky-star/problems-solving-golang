package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
)

func main() {
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	s := []int{}
	var m sync.Mutex

	receivedOrderCh := receivedOrders()
	validOrderCh, invalidOrderCh := validateOrders(receivedOrderCh)
	reservedInventoryCh := reserveInventory(validOrderCh)

	wg.Add(2)
	go func(invalidOrderCh <-chan invalidOrder) {
		for order := range invalidOrderCh {
			fmt.Printf("Invalid order received: %v. Issue: %v\n", order.order, order.err)
		}
		wg.Done()
	}(invalidOrderCh)

	go func(reservedInventoryCh <-chan order) {
		for order := range reservedInventoryCh {
			fmt.Printf("Inventory reserved for: %v\n", order)
		}
		wg.Done()
	}(reservedInventoryCh)

	// 	go func(validOrderCh <-chan order, invalidOrderCh <-chan invalidOrder) {
	// 	loop:
	// 		for {
	// 			select {
	// 			case order, ok := <-validOrderCh:
	// 				if ok {
	// 					fmt.Printf("Valid order received: %s", order)
	// 				} else {
	// 					break loop
	// 				}

	// 			case order, ok := <-invalidOrderCh:
	// 				if ok {
	// 					fmt.Printf("Invalid order received: %s. Issue: %s", order.order, order.err)
	// 				} else {
	// 					break loop
	// 				}

	//			}
	//		}
	//		wg.Done()
	//	}(validOrderCh, invalidOrderCh)
	//
	wg.Wait()

	// Demonstrating how to use Mutexes
	ctx, cancel := context.WithCancel(context.Background())
	const interations = 1000
	wg2.Add(interations)
	for i := range interations {
		go func(ctx context.Context) {
			m.Lock()
			s = append(s, i)
			defer m.Unlock()
			wg2.Done()
			for range time.Tick(500 * time.Millisecond) {
				if ctx.Err() != nil {
					log.Println(ctx.Err())
					return
				}
				fmt.Println("Tick...")
			}
		}(ctx)
		time.Sleep(2 * time.Millisecond)
		cancel()
	}
	wg2.Wait()
	fmt.Printf("Length of s: %d\n", len(s))
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

func reserveInventory(in <-chan order) <-chan order {
	out := make(chan order)
	go func() {
		for o := range in {
			// Simulate inventory reservation
			o.Status = reserved
			out <- o
		}
		close(out)
	}()
	return out
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

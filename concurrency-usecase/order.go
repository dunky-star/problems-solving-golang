package main

import "fmt"

type order struct {
	ProductCode int
	Quantity    float64
	Status      orderStatus
}

func (o order) String() string {
	return fmt.Sprintf("Order{ProductCode: %d, Quantity: %.2f, Status: %s}", o.ProductCode, o.Quantity, orderStatusToText(o.Status))
}

func orderStatusToText(o orderStatus) string {
	switch o {
	case none:
		return "none"
	case new:
		return "new"
	case received:
		return "received"
	case reserved:
		return "reserved"
	case filled:
		return "filled"
	default:
		return "unknown status"
	}
}

type orderStatus int

const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)

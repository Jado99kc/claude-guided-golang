package main

import (
	"fmt"
	"strconv"
)

// OrderStatus is the lifecycle state of an order.
type OrderStatus int

// Order statuses. StatusUnknown is the zero value, so an order whose
// status was never set is detectable instead of silently "pending".
const (
	StatusUnknown   OrderStatus = iota // 0
	StatusPending                      // 1
	StatusPaid                         // 2
	StatusShipped                      // 3
	StatusDelivered                    // 4
	StatusCancelled                    // 5
)

var statusNames = map[OrderStatus]string{
	StatusUnknown:   "unknown",
	StatusPending:   "pending",
	StatusPaid:      "paid",
	StatusShipped:   "shipped",
	StatusDelivered: "delivered",
	StatusCancelled: "cancelled",
}

// Compile-time check that OrderStatus satisfies fmt.Stringer.
var _ fmt.Stringer = StatusUnknown

// String returns the lowercase name of the status, or "OrderStatus(N)"
// for a value that is not one of the declared constants.
func (s OrderStatus) String() string {
	if name, ok := statusNames[s]; ok {
		return name
	}

	return "OrderStatus(" + strconv.Itoa(int(s)) + ")"
}

//chose a value receiver since the function String will not modify the original Order status

var _ fmt.Stringer = OrderStatus(0)

func main() {
	fmt.Println(StatusUnknown)
	fmt.Println(StatusShipped)
	fmt.Println(OrderStatus(99))
	fmt.Println(OrderStatus(-1))
}

// if String() had a pointer receiver it would not satisfy the Stringer interface since Stringer expects the method on the value and not the pointer

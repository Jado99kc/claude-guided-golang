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

func (o OrderStatus) String() string {
	m := map[int]string{0: "unknown", 1: "pending", 2: "paid", 3: "shipped", 4: "delivered", 5: "cancelled"}

	if _, ok := m[o]; ok {
		return m[o]
	}

	if o == StatusShipped {
		return "shipped"
	}
	return strconv.FormatInt(0, 10)
}

//chose a value receiver since the function String will not modify the original Order status

var _ fmt.Stringer = (OrderStatus)(0)

func main() {
	fmt.Println(StatusShipped)
	fmt.Println(99)
}

// if String() had a pointer receiver it would not satisfy the Stringer interface since Stringer expects the method on the value and not the pointer

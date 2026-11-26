package main

import (
	"errors"
	"fmt"
)

// import (
// 	"fmt"
// 	"math"
// )

// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// type Rect struct{ W, H float64 }
// type Circle struct{ R float64 }

// func (r Rect) Area() float64        { return r.W * r.H }
// func (r Rect) Perimeter() float64   { return 2 * (r.W + r.H) }
// func (c Circle) Area() float64      { return math.Pi * c.R * c.R }
// func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.R }

// func main() {
// 	var s Shape = Circle{2}

// 	c, ok := s.(Circle) // "comma ok" prevents a panic if s isn't a Circle

// 	fmt.Println(c)
// 	fmt.Println(ok)
// 	switch v := s.(type) {
// 	case Rect:
// 		fmt.Println("rect", v.W)
// 	case Circle:
// 		fmt.Println("circle", v.R)
// 	}

// }

type Notifier interface {
	Notify(to, message string) error
}

type MultiNotifier []Notifier

type EmailNotifier struct{}
type SMSNotifier struct{}

func (en EmailNotifier) Notify(to, message string) error {
	fmt.Printf("to: %s, message: %s", to, message)
	return nil
}
func (sn SMSNotifier) Notify(to, message string) error {
	fmt.Printf("to: %s, message: %s", to, message)
	return nil
}
func (mn MultiNotifier) Notify(to, message string) error {
	var errs []error
	for _, notification := range mn {

		err := notification.Notify(to, message)

		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

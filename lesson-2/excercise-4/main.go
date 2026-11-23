package main

import (
	"errors"
	"fmt"
)

// Product Schema
type Product struct {
	SKU        string
	Name       string
	PriceCents int64
	Qty        int
}

// Inventory holds products keyed by SKU.
type Inventory map[string]Product

// Product constructor
func NewProduct(sku, name string, priceCents int64, qty int) (Product, error) {
	if sku == "" {
		return Product{}, errors.New("sku is required")
	}
	if priceCents < 0 {
		return Product{}, fmt.Errorf("product %q: negative price %d", sku, priceCents)
	}
	if qty < 0 {
		return Product{}, fmt.Errorf("product %q: negative quantity %d", sku, qty)
	}
	return Product{SKU: sku, Name: name, PriceCents: priceCents, Qty: qty}, nil
}

// AddProduct adds p to inv. It returns an error if a product with the
// same SKU already exists; the existing product is left unchanged.
func AddProduct(inv Inventory, p Product) error {
	if _, exists := inv[p.SKU]; exists {
		return fmt.Errorf("product with SKU %q already exists", p.SKU)
	}
	inv[p.SKU] = p
	return nil
}

// Restock adds n units to the product with the given SKU.
// It returns an error if n <= 0 or the SKU is not in inv.
func Restock(inv Inventory, sku string, n int) error {
	if n <= 0 {
		return fmt.Errorf("restock %q with invalid quantity %d", sku, n)
	}
	item, exists := inv[sku]
	if !exists {
		return fmt.Errorf("cannot restock item SKU %q as it does not exist in inventory", sku)
	}
	item.Qty += n
	inv[sku] = item
	return nil
}

// TotalValueCents returns the total stock value of inv in cents
// (the sum of PriceCents * Qty over all products).
func TotalValueCents(inv Inventory) int64 {
	var total int64
	for _, item := range inv {
		total += int64(item.Qty) * item.PriceCents
	}
	return total
}

func formatCents(c int64) string {
	sign := ""
	if c < 0 {
		sign = "-"
		c = -c
	}
	return fmt.Sprintf("%s$%d.%02d", sign, c/100, c%100)
}
func main() {
	inv := make(Inventory)
	seed := []struct {
		sku   string
		name  string
		price int64
		qty   int
	}{
		{"1", "soap", 300, 2},
		{"2", "cornflakes", 852, 3},
		{"2", "cornflakes", 852, 2},
		{"3", "soy", 1369, 1},
	}
	for _, v := range seed {
		item, err := NewProduct(v.sku, v.name, v.price, v.qty)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if err := AddProduct(inv, item); err != nil {
			fmt.Println(err)
		}

	}
	if err := Restock(inv, "2", 2); err != nil {
		fmt.Println(err)
	}
	if err := Restock(inv, "5", 2); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total inventory value:", formatCents(TotalValueCents(inv)))
}

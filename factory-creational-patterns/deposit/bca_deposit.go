package deposit

import "fmt"

type BCA struct {
	// Category string
}

func (bca *BCA) SetCategory() string {
	return "Virtual Account"
}

func (bca *BCA) Payment(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using BCA \n", amount)
}

package deposit

import "fmt"

type BNI struct{}

func (bca *BNI) SetCategory() string {
	return "Bank Transfer"
}

func (bni *BNI) Payment(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using BNI \n", amount)
}

package deposit

type IPayment interface {
	Payment(amount float32) string
	SetCategory() string
}

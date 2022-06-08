package services

import (
	"factory-creational-patterns/deposit"
	"strings"
)

type Deposit struct {
	categories []string
}

func (d *Deposit) DepositRegister(depoStruct deposit.IPayment) deposit.IPayment {
	getCategory := depoStruct.SetCategory()

	d.categories = append(d.categories, getCategory)

	return depoStruct
}

func (d *Deposit) ListCategory() string {
	return strings.Join(d.categories[:], ",")
}

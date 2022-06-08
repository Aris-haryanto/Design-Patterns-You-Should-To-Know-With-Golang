package main

import (
	"factory-creational-patterns/deposit"
	"factory-creational-patterns/services"
	"fmt"
	"net/http"
)

func main() {

	depoSvc := services.Deposit{}

	//register your deposit category
	bca := depoSvc.DepositRegister(new(deposit.BCA))
	bni := depoSvc.DepositRegister(new(deposit.BNI))

	http.HandleFunc("/depo-bca", func(w http.ResponseWriter, r *http.Request) {
		result := bca.Payment(5000)
		w.Write([]byte(result))
		return
	})

	http.HandleFunc("/depo-bni", func(w http.ResponseWriter, r *http.Request) {
		result := bni.Payment(10000)
		w.Write([]byte(result))
		return
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		list := depoSvc.ListCategory()
		w.Write([]byte(list))
		return
	})

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

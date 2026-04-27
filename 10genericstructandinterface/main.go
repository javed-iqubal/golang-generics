package main

import (
	"errors"
	"fmt"
)

type CreditCard struct {
	CardNumber string
	CardHolder string
}

type PayPal struct {
	EmailId string
}

type Payer interface {
	pay(amount float64) string
}

func (c CreditCard) pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card number ending with %s", amount, c.CardNumber[len(c.CardNumber)-4:])
}

func (p PayPal) pay(amount float64) string {

	return fmt.Sprintf("Paid %.2f Using PayPal account %s", amount, p.EmailId)

}

// func ProcessPayment(payer Payer, amount float64) {
// 	fmt.Println(payer.pay(amount))
// }

type Result[T any] struct {
	value T
	err   error
}

func ProcessPayment[T Payer](p T, amount float64) Result[string] {
	if amount <= 0 {
		return Result[string]{
			err: errors.New("amount must be positive"),
		}
	}
	return Result[string]{
		value: p.pay(amount),
		err:   nil,
	}
}

func main() {

	payer := CreditCard{
		CardNumber: "012-345-6789",
		CardHolder: "John Deo",
	}

	result := ProcessPayment(payer, 260)

	if result.err != nil {
		fmt.Println(result.err)
	} else {
		fmt.Println(result.value)
	}
	payer1 := PayPal{
		EmailId: "john.deo@mymail.com",
	}
	result = ProcessPayment(payer1, 300)
	if result.err != nil {
		fmt.Println(result.err)
	} else {
		fmt.Println(result.value)
	}
}

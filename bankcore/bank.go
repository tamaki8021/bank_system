package bankcore

import "errors"

type Customer struct {
	Name string
	Address string
	Phone string
}

type Account struct {
	Customer
	Number int32
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance -= amount
	return nil
}
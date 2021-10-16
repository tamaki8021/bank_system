package bankcore

import (
	"errors"
	"fmt"
)

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

type Bank interface {
	Statement() string
}

// 預金
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

// 引き出し
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}

	a.Balance -= amount
	return nil
}

// 明細
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

//送金
func (a *Account) Transfer(amount float64, dest *Account) error {
	if amount <= 0 {
		return errors.New("the amount to transfer should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to transfer should be greater than accaunt's balance")
	}

	a.Withdraw(amount)
	dest.Deposit(amount)
	return nil
}

func Statement(b Bank) string {
	return b.Statement()
}
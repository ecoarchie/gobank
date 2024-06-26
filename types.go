package main

import (
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type Account struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Number int64 `json:"number"`
	Balance int64 `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount int `json:"amount"`
}

func NewAccount(fName, lName string) *Account {
	return &Account{
		// ID:        rand.Intn(10000),
		FirstName: fName,
		LastName:  lName,
		// Number:    int64(rand.Intn(1000000)),
		Balance:   0,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
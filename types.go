package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	ToAccount int64 `json:"toAccount"`
	Amount    int64 `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		//ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Int63(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(), //DOES mysql do it for me.
		// balance goes to 0 normally with go
	}
}

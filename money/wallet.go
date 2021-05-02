package money

import (
	"fmt"
	"math/rand"
)

type Wallet struct {
	Id      string
	Value   int64
	Address string
}

func NewWallet() Wallet {
	a := Wallet{}
	b := make([]byte, 16)
	rand.Read(b)
	a.Id = fmt.Sprintf("%X", b)
	return a
}
func NewWalletWithValue(id string, value int64) Wallet {
	a := NewWallet()
	a.Id = id
	a.Value = value
	return a
}

func (a *Wallet) AssociatedFund() int64 {
	return a.Value
}

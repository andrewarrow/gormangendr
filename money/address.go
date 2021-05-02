package money

import (
	"fmt"
	"math/rand"
)

type Address struct {
	Id    string
	Value int64
}

func NewAddress() Address {
	a := Address{}
	b := make([]byte, 16)
	rand.Read(b)
	a.Id = fmt.Sprintf("%X", b)
	return a
}
func NewAddressWithValue(id string, value int64) Address {
	a := NewAddress()
	a.Id = id
	a.Value = value
	return a
}

func (a *Address) AssociatedFund() string {
	return "AssociatedFund"
}

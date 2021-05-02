package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Address struct {
	Val string
}

func NewAddress() Address {
	a := Address{}
	b := make([]byte, 16)
	rand.Read(b)
	a.Val = fmt.Sprintf("%X", b)
	return a
}

type ConfigurationBuilder struct {
	Txs []Tx
}

func NewConfigurationBuilder() *ConfigurationBuilder {
	cb := ConfigurationBuilder{}
	return &cb
}

func (cb *ConfigurationBuilder) GenesisBlockHash() string {
	return "ABC"
}
func ConfigurationBuilderWithFunds(txs []Tx) *ConfigurationBuilder {
	cb := ConfigurationBuilder{}
	cb.Txs = txs
	return &cb
}

type Tx struct {
	Address string
	Val     int64
}

func NewTx(address string, val int64) Tx {
	tx := Tx{}
	tx.Address = address
	tx.Val = val
	return tx
}

/*
  let leader_config = ConfigurationBuilder::new()
        .with_funds(vec![InitialUTxO {
            address: sender.address(),
            value: 100.into(),
        }])

    let transaction_message = jcli
        .transaction_builder(block0_hash)
        .build_transaction_from_utxo(
            &utxo,
            *utxo.associated_fund(),
            &sender,
            *utxo.associated_fund(),
            &reciever,
        );
*/

func TwoNodesCommunication() {
	sender := NewAddress()
	receiver := NewAddress()

	fmt.Println(sender.Val)
	fmt.Println(receiver.Val)

	txs := []Tx{}
	txs = append(txs, NewTx("FROG", 123))
	leaderConfig := ConfigurationBuilderWithFunds(txs)
	fmt.Println(leaderConfig)
}
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("testing...")
	TwoNodesCommunication()
}

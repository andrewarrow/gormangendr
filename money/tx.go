package money

type Tx struct {
	Address string
	Val     int64
}

func NewTx() Tx {
	tx := Tx{}
	return tx
}

func (tx *Tx) BuildFromHash(hash string) {
}
func (tx *Tx) SetFromUtxo(utxo *Address, af1 string,
	sender *Address, af2 string, receiver *Address) {
}

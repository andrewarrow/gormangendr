package network

type Blockchain struct {
	//branches: Branches,
	//ref_cache: RefCache,
	//ledgers: Multiverse<Ledger>,
	//storage: Storage,
	Block0 string
}

func NewBlockchain(genesisBlockHash string) *Blockchain {
	bc := Blockchain{}
	bc.Block0 = genesisBlockHash
	return &bc
}

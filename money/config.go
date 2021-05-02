package money

type Config struct {
	Addresses    []Address
	TrustedPeers []string
	BlockHash    string
}

func NewConfig() *Config {
	c := Config{}
	c.TrustedPeers = []string{}
	return &c
}

func (c *Config) Block0UtxoForAddress(sender *Wallet) Address {
	a := NewAddress()
	return a
}

func (c *Config) GenesisBlockHash() string {
	return "ABC"
}
func (c *Config) AddTrustedPeer(tp string) {
	c.TrustedPeers = append(c.TrustedPeers, tp)
}
func ConfigWithFunds(adx []Address) *Config {
	c := Config{}
	c.Addresses = adx
	return &c
}

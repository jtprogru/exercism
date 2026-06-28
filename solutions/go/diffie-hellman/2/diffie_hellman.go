package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

var two = big.NewInt(2)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	// Pick a private key uniformly from [2, p).
	max := new(big.Int).Sub(p, two)
	n, _ := rand.Int(rand.Reader, max)
	return n.Add(n, two)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	return private, PublicKey(private, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}

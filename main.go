package main

import (
	"math"
)

const (
	targetBits = 24
	maxNonce   = math.MaxInt64
)

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to Roman")
	bc.AddBlock("Send 2 BTC to Roman")
}

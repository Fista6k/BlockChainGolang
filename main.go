package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	targetBits = 24
	maxNonce   = math.MaxInt64
)

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to Roman")
	bc.AddBlock("Send 2 BTC to Roman")

	for _, block := range bc.blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

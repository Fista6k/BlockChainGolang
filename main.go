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
	defer bc.db.Close()

	cli := CLI{
		bc: bc,
	}
	cli.Run()
}

package main

import (
	"fmt"
	"github.com/filipk999/go-chain/blockchain"
	"strconv"
)

func main() {
	chain := blockchain.InitChain()
	chain.AddBlock("Hi")
	chain.AddBlock("Hi")
	chain.AddBlock("Hi")

	for _, block := range chain.Blocks {
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("Block Hash: %x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("Verified: %s\n", strconv.FormatBool(pow.Validate()))
	}
}

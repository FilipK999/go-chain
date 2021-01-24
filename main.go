package main

import (
	"github.com/filipk999/go-chain/blockchain"
)

func main() {
	chain := blockchain.InitChain()
	chain.Iterator().Next()
	chain.AddBlock("Hi")
	chain.AddBlock("Hi 2")
	chain.AddBlock("Hi 3")

	// for _, block := range chain.Blocks {
	// 	fmt.Printf("Block Data: %s\n", block.Data)
	// 	fmt.Printf("Block Hash: %x\n", block.Hash)
	// 	fmt.Printf("Date mined: %s\n", ParseTimeStamp(block.TimeStamp))
	// 	pow := blockchain.NewProof(block)
	// 	fmt.Printf("Verified: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }
}

//
// func ParseTimeStamp(ts int64) time.Time {
// 	return time.Unix(ts, 0)
// }

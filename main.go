package main

import (
	"flag"
	"fmt"
	"github.com/filipk999/go-chain/blockchain"
	"os"
	"runtime"
	"strconv"
	"time"
)

type CommandLine struct {
	blockchain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("You haven't specified any arguments. Available arguments are:")
	fmt.Println(" add BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - Prints the blocks")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
}
func (cli *CommandLine) printChain() {
	iter := cli.blockchain.Iterator()

	for {
		block := iter.Next()

		fmt.Println()
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Date mined: %s\n", time.Unix(block.TimeStamp, 0))
		pow := blockchain.NewProof(block)

		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func main() {
	defer os.Exit(0)
	chain := blockchain.InitChain()
	defer chain.Database.Close()

	cli := CommandLine{chain}
	cli.run()
}

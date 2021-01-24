package blockchain

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

type BlockChain struct {
	LastHash []byte
	Database *redis.Client
}

type BlockChainIterator struct {
	LastHash []byte
	Database *redis.Client
}

var ctx = context.TODO()

func InitChain() *BlockChain {
	var lastHash []byte

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := client.Get(ctx, "lastHash").Result()

	// If there's no lastHash (meaning the BlockChain doesn't exist) create a genesis block
	if err == redis.Nil {
		println("No blockchain found in database, creating a new one...")
		genesis := genesis()

		_, err := client.Set(ctx, StrHash(genesis.Hash), genesis, 0).Result()
		Handle(err)

		lastHash = genesis.Hash
		client.Set(ctx, "lastHash", lastHash, 0)
	}

	item, err := client.Get(ctx, "lastHash").Result()
	Handle(err)
	lastHash = []byte(item)

	return &BlockChain{lastHash, client}
}

// Gets the lastHash from database, creates a new block using the hash, save the block
// and it's hash as the new lastHash
func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	item, err := chain.Database.Get(ctx, "lastHash").Result()
	Handle(err)
	lastHash = []byte(item)
	newBlock := createBlock(data, lastHash)

	_, err = chain.Database.Set(ctx, StrHash(newBlock.Hash), newBlock, 0).Result()
	Handle(err)
	_, err = chain.Database.Set(ctx, "lastHash", newBlock.Hash, 0).Result()
}

// -----------------
// Utility Functions
// -----------------

// This function gets called automatically by go-redis
func (b *Block) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

// Converts the bytes into a string representation of the hash
func StrHash(bytes []byte) string {
	return fmt.Sprintf("%x", bytes)
}

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

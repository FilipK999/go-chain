package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"
)

const difficulty = 20

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) joinData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.Block.PrevHash,
		pow.Block.Data,
		toBytes(time.Now().Unix()),
		toBytes(int64(nonce)),
		toBytes(difficulty),
	}, []byte{})

	return data
}

func (pow *ProofOfWork) Mine() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.joinData(nonce)
		hash = sha256.Sum256(data)
		intHash.SetBytes(hash[:])

		fmt.Printf("\r%x", hash)

		if intHash.Cmp(pow.Target) == -1 {
			pow.Block.TimeStamp = time.Now().Unix()
			break
		} else {
			nonce++
		}

	}
	fmt.Println()

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	hash := sha256.Sum256(pow.joinData(pow.Block.Nonce))
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

func toBytes(n int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, n)
	if err != nil {
		log.Fatal(err)
	}
	return buff.Bytes()
}

package blockchain

type Block struct {
	Data      []byte
	Hash      []byte
	PrevHash  []byte
	Nonce     int
	TimeStamp int64
}

func createBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Data:      []byte(data),
		Hash:      []byte{},
		PrevHash:  prevHash,
		Nonce:     0,
		TimeStamp: 0,
	}
	pow := NewProof(block)
	nonce, hash := pow.Mine()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

func genesis() *Block {
	return createBlock("Genesis", []byte{})
}

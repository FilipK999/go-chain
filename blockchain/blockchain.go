package blockchain

type BlockChain struct {
	Blocks []*Block
}

func InitChain() *BlockChain {
	return &BlockChain{Blocks: []*Block{genesis()}}
}

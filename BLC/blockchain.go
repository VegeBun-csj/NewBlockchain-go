package BLC

//区块链结构体
type BlockChain struct {
	//是一个Block结构体数组，用来存储Block
	Blocks []*Block
}


//调用BlockChain对象新增区块
func (blockChain *BlockChain)AddBlock(data string)  {
	//1.创建一个新的区块Block
	//当前blocks数组的最后一个区块的hash就是当前要添加的区块的上一个hash
	CurBlockHash := blockChain.Blocks[len(blockChain.Blocks) - 1].Hash

	newblock := NewBlock(data, CurBlockHash)

	//2.将区块添加到Blocks数组中
	blockChain.Blocks = append(blockChain.Blocks, newblock)
}


//创建一个带有创世区块的区块链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
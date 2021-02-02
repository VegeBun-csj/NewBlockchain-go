package BLC

import (
	"fmt"
	"testing"
	"time"
)

func TestNewBlockChain(t *testing.T) {
	blockchain := NewBlockChain()
	blockchain.AddBlock("a transfer b 20 BTC")
	blockchain.AddBlock("a transfer c 10 BTC")
	blockchain.AddBlock("a transfer d 30 BTC")
	blockchain.AddBlock("a transfer e 20 BTC")

	fmt.Printf("当前区块链有%v个区块 \n",len(blockchain.Blocks))

	for blockNum,block :=range blockchain.Blocks{
		fmt.Printf("第%v个区块的区块数据为%v,区块产生时间为%v,当前区块的hash为%x,上一个区块的hash为%x\n",
			blockNum,string(block.Data),
			time.Unix(block.Timestamp,0).Format("2006-01-02 15:04:05"),
			block.Hash,
			block.PrevBlockHash)
	}

}

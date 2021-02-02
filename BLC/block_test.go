package BLC

import (
	"fmt"
	"testing"
)

func TestNewBlock(t *testing.T) {

	//区块中的hash为256位，16进制表示，32个字节

	//block := NewBlock("Genesis Block", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	//fmt.Println(block)
	//fmt.Println(block.Timestamp)
	//fmt.Printf("PreBlockHash:%x\n", block.PrevBlockHash)
	//fmt.Println(string(block.Data))
	//fmt.Printf("Current Block Hash: %x",block.Hash)


	Genesisblock :=NewGenesisBlock()
	fmt.Println(Genesisblock)
	fmt.Println(Genesisblock.Timestamp)
	fmt.Printf("PreBlockHash:%x\n", Genesisblock.PrevBlockHash)
	fmt.Println(string(Genesisblock.Data))
	fmt.Printf("Current Block Hash: %x",Genesisblock.Hash)
}
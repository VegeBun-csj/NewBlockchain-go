package main

import (
	"NewBlockchain-go/BLC"
	"fmt"
)

func main() {

	block := BLC.NewBlock("Genesis Block", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(block)
	fmt.Println(block.Timestamp)
	fmt.Printf("PreBlockHash:%x\n", block.PrevBlockHash)
	fmt.Println(string(block.Data))
	fmt.Printf("Current Block Hash: %x",block.Hash)
}

package main

import (
	"fmt"
	"huhuBLC/BLC"
)

func main() {
	bc := BLC.CreateBlockChainWithGenesisBlock()
	fmt.Printf("blockchain:%v\n", bc.Blocks[0])

	//上链
	bc.AddBlock(bc.Blocks[len(bc.Blocks)-1].Height+1,bc.Blocks[len(bc.Blocks)-1].Hash,[]byte("Peter send 10 to Wawa"))

	for _, block := range bc.Blocks {
		fmt.Printf("PrevBlockHash:%x\n", block.PrevBlockHash)
		fmt.Printf("currentHash:%x\n", block.Hash)
	}

	//第二次上链
	bc.AddBlock(bc.Blocks[len(bc.Blocks)-1].Height+1,bc.Blocks[len(bc.Blocks)-1].Hash,[]byte("Wawa send 20 to Huhu"))

	for _, block := range bc.Blocks {
		fmt.Printf("PrevBlockHash:%x\n", block.PrevBlockHash)
		fmt.Printf("currentHash:%x\n", block.Hash)
	}
}

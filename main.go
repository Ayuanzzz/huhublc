package main

import "HUHUBLC/BLC"

func main() {
	// bc := BLC.CreateBlockChainWithGenesisBlock()

	// bc.AddBlock([]byte("alice send 100 eth to bob"))
	// bc.PrintChain()
	cli := BLC.CLI{}
	cli.Run()

}

package main

import "bitcoin/bkc/c16-new-tx/BLC"

//启动
func main(){
	//bc:=BLC.CreateBlockChainWithGenesisBlock()
	//bc.AddBlock([]byte("a send 100 eth to b"))
	//bc.AddBlock([]byte("b send 100 eth to a"))
	//bc.PrintChain()
	cli:=BLC.CLI{}
	cli.Run()
}
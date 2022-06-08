package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//对blockchain的命令行操作进行管理

//client对象
type CLI struct {
}

//用法展示
func PrintUsage() {
	fmt.Printf("Usage:\n")
	//初始化区块链
	fmt.Printf("createblockchain\n")
	//添加区块
	fmt.Printf("addblock\n")
	//打印完整的区块信息
	fmt.Printf("printchain\n")

}

//初始化区块链
func (cli *CLI) createBlockchain() {
	CreateBlockChainWithGenesisBlock()
}

//添加区块
func (cli *CLI) addBlock(data string) {
	//判断数据库是佛存在
	if !dbExist() {
		fmt.Printf("数据库不存在...\n")
		os.Exit(1)
	}
	//获取blockchain的对象实例
	blockChain := BlockchainObject()
	blockChain.AddBlock([]byte(data))
}

//打印完整区块链信息
func (cli *CLI) printchain() {
	//判断数据库是佛存在
	if !dbExist() {
		fmt.Printf("数据库不存在...\n")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	blockchain.PrintChain()
}

//参数数量检查函数
func IsValidArgs() {
	if len(os.Args) < 2 {
		PrintUsage()
		//直接退出
		os.Exit(1)
	}
}

//命令行运行函数
func (cli *CLI) Run() {
	//检测参数数量
	IsValidArgs()
	//新建相关命令
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	//输出区块链完整信息
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	//创建区块链
	createBLCWithGenesisBlockCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	//数据参数处理
	flagAddBlockArg := addBlockCmd.String("data", "send 100 to wawa", "添加区块数据")

	//判断命令
	switch os.Args[1] {
	case "addblock":
		if err := addBlockCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse addBlockCmd failed %v\n", err)
		}
	case "printchain":
		if err := printChainCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse printChainCmd failed %v\n", err)
		}
	case "createblockchain":
		if err := createBLCWithGenesisBlockCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse createblockchain failed %v\n", err)
		}
	default:
		PrintUsage()
		os.Exit(1)
	}

	//添加区块命令
	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.addBlock(*flagAddBlockArg)
	}
	//输出区块链信息
	if printChainCmd.Parsed() {
		cli.printchain()
	}
	//创建区块链命令
	if createBLCWithGenesisBlockCmd.Parsed() {
		cli.createBlockchain()
	}
}

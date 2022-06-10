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
	fmt.Printf("\tcreateblockchain -address address\n")
	//添加区块
	fmt.Printf("\taddblock -data DATA\n")
	//打印完整的区块信息
	fmt.Printf("\tprintchain\n")

}

//初始化区块链
func (cli *CLI) createBlockchain(address string) {
	CreateBlockChainWithGenesisBlock(address)
}

//添加区块
func (cli *CLI) addBlock(txs []*Transaction) {
	//判断数据库是佛存在
	if !dbExist() {
		fmt.Printf("数据库不存在...\n")
		os.Exit(1)
	}
	//获取blockchain的对象实例
	blockChain := BlockchainObject()
	blockChain.AddBlock(txs)
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
	//添加区块
	flagAddBlockArg := addBlockCmd.String("data", "send 100 to wawa", "添加区块数据")
	//创建区块链时指定的矿工地址(接受奖励)
	flagCreateBlockchainArg := createBLCWithGenesisBlockCmd.String("address", "huhu", "指定接受系统奖励的矿工地址")

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
		cli.addBlock([]*Transaction{})
	}
	//输出区块链信息
	if printChainCmd.Parsed() {
		cli.printchain()
	}
	//创建区块链命令
	if createBLCWithGenesisBlockCmd.Parsed() {
		if *flagCreateBlockchainArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.createBlockchain(*flagCreateBlockchainArg)
	}
}

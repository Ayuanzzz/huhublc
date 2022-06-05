package BLC

import (
	"log"

	"github.com/boltdb/bolt"
)

//数据库名称
const dbName = "block.db"

//表名称
const blockTableName = "blocks"

//区块链基本结构
type BlockChain struct {
	DB  *bolt.DB //数据库对象
	Tip []byte   //保存最新区块的哈希值
}

//初始化区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	//保存最新区块的哈希值
	var blockHash []byte
	//1.创建或者打开一个数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Panicf("create db [%s] failed %v\n", dbName, err)
	}
	//2.创建桶，把生成的创世区块存入数据库中
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			//没找到桶
			b, err := tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panicf("create bucket [%s] failed %v\n", dbName, err)
			}
			//生成创世区块
			genesisBlock := CreateGenesisBlock([]byte("init blockchain"))
			//存储
			//key--hash
			//value--序列化
			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panicf("insert the gensis block failed %v/n", err)
			}
			blockHash = genesisBlock.Hash
			//存储最新区块的哈希
			//1:latest
			err = b.Put([]byte("1"), genesisBlock.Hash)
			if err != nil {
				log.Panicf("save the hash of genesis block failed %v/n", err)
			}
		}
		return nil
	})
	return &BlockChain{DB: db, Tip: blockHash}
}

//添加区块到区块链中
func (bc *BlockChain) AddBlock(data []byte) {
	//更新区块数据(insert)
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		//1.获取bucket
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//2.获取最后插入的区块
			blockBytes := b.Get(bc.Tip)
			//3.区块数据反序列化
			latest_block := DeserialiezBlock(blockBytes)
			//4.新建区块
			newBlock := NewBlock(latest_block.Height+1, latest_block.Hash, data)
			//5.存入数据库
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panicf("insert the new block to db failed %v/n", err)
			}
			//6.更新数据库中最新区块的哈希
			err = b.Put([]byte("1"), newBlock.Hash)
			if err != nil {
				log.Panicf("update the latest block hash to db failed %v", err)
			}
			//7.更新区块链对象中的最新区块哈希
			bc.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panicf("insert block to db failed %v", err)
	}
}

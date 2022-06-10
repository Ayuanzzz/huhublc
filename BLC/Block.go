package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

//基本区块结构
type Block struct {
	TimeStamp     int64          //区块时间戳，代表区块时间
	Hash          []byte         //当前区块哈希
	PrevBlockHash []byte         //前区块哈希
	Height        int64          //区块高度
	Txs           []*Transaction //交易数据(交易列表)
	Nonce         int64          //在运行pow时生成的哈希变化值，也代表pow运行时动态修改的数据
}

//新建区块
func NewBlock(height int64, preBlockHash []byte, txs []*Transaction) *Block {
	var block Block
	block = Block{
		TimeStamp:     time.Now().Unix(),
		Hash:          nil,
		PrevBlockHash: preBlockHash,
		Height:        height,
		Txs:           txs,
	}
	//通过POW生成新的哈希
	pow := NewProofOfWork(&block)
	//执行工作量证明算法
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = int64(nonce)
	//执行工作量证明算法
	return &block
}

//生成创世区块函数
func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(1, nil, txs)
}

//实现int64转[]byte
func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, data)
	if nil != err {
		log.Panicf("int transact to []byte failed! %v/n", err)
	}
	return buffer.Bytes()
}

//区块结构序列化
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	//新建编码对象
	encoder := gob.NewEncoder(&buffer)
	//序列化
	if err := encoder.Encode(block); err != nil {
		log.Panicf("serialize the []byte to block failed %v/n", err)
	}
	return buffer.Bytes()
}

//区块数据反序列化
func DeserialiezBlock(blockBytes []byte) *Block {
	var block Block
	//新建decode对象
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&block); err != nil {
		log.Panicf("deserialize the []byte to block failed %v/n", err)
	}
	return &block
}

//把指定区块中所有交易结构都序列化
func (block *Block) HashTransaction() []byte {
	var txHashes [][]byte
	//将指定区块中所有交易哈希进行拼接
	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxHash)
	}
	txHash := sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

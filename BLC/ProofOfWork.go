package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//目标难度值
const targetBit = 16

type ProofOfWork struct {
	//需要共识验证的区块
	Block *Block
	//目标难度的哈希(大数据存储)
	target *big.Int
}

//创建POW对象
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	//左移targetBit位
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{Block: block, target: target}
}

//执行pow，比较哈希
//返回哈希值，以及碰撞次数
func (proofOfWork *ProofOfWork) Run() ([]byte, int) {
	//碰撞次数
	var nonce = 0
	var hashInt big.Int
	//生成的哈希值
	var hash [32]byte
	//无限循环，生成符号条件的哈希值
	for {
		//生成准备数据
		dataBytes := proofOfWork.prepareData(int64(nonce))
		hash = sha256.Sum256(dataBytes)
		hashInt.SetBytes(hash[:])
		//检测生成的哈希值是否符合条件
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			//找到符合条件的哈希值，中断循环
			break
		}
		nonce++
	}
	fmt.Printf("\n碰撞次数:%d\n", nonce)
	return hash[:], nonce
}

//生成准备数据
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	var data []byte
	//拼接区块属性,进行哈希计算
	timeStampBytes := IntToHex(pow.Block.TimeStamp)
	heightBytes := IntToHex(pow.Block.Height)
	data = bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		pow.Block.PrevBlockHash,
		pow.Block.HashTransaction(),
		IntToHex(nonce),
		IntToHex(targetBit),
	}, []byte{})
	return data
}

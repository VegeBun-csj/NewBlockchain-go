package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//这是一个最简单的区块的结构
type Block struct {
	//时间戳
	Timestamp int64

	//上一个区块的hash
	PrevBlockHash []byte

	//当前区块的hash
	Hash []byte

	//data交易数据
	Data []byte
}


//有一个创建区块的方法（工厂方法）
/**
 * @Description:
 * @param data	区块数据
 * @param preBlockHash 上一个区块hash
 * @return *Block
 */
func NewBlock(data string,preBlockHash []byte) *Block {
	//创建区块(一开始创建区块的hash暂时设置为空)
	//time.Now()生成的是一个Time对象，要将其转化为int64,使用Unix()方法
	block := &Block{time.Now().Unix(),preBlockHash,[]byte{},[]byte(data)}

	//设置当前区块的hash
	block.SetHash()

	//返回区块
	return block

}


//设置区块的hash值
/**
 * @Description:
 * @receiver block
 */
func (block *Block) SetHash() {

	//1. 将时间戳转换为字节数组（（1）先将时间戳int64转换为字符串
	//第二个参数范围是(2~36),它是转的进制数(此处就是转为二进制)
	timeString := strconv.FormatInt(block.Timestamp, 2)
	//（2）然后转为字节数组）
	timeStamp := []byte(timeString)
	//2. 将除了Hash之外的字段拼接为字节数组
	header := bytes.Join([][]byte{timeStamp, block.PrevBlockHash, block.Data}, []byte{})
	//3. 将拼接起来的字节数组进行sha256 hash
	CurHash := sha256.Sum256(header)
	//4. 将该hash赋值给当前区块的hash
	block.Hash = CurHash[:]
}



//创建创世区块

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}




























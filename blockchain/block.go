package blockchain

import (
	"DetaCerProject/uilt"
	"bytes"
	"encoding/gob"
	"time"
)

/**
区块结构的定义
*/

type Block struct {
	Height     int64    //区块高度
	Size       int    //区块大小
	TiemeStamp int64  //时间戳
	Hash       []byte //区块的Hash
	Data       []byte //数据
	PrevHash   []byte //上一个区块的Hash
	Version    string //版本号
	Nonce      int64  //随机数

}
/**
生成创世区块，并返回区块信息
 */
func CreateGenesisBlock() Block{
	block:=NewBlock(0,[]byte{},[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
return  block
}

/**
新建一个区块链实例
*/

func NewBlock(height int64, data []byte, prevHash []byte) Block {
	block := Block{
		Height: height ,

		TiemeStamp: time.Now().Unix(),

		Data:     data,
		PrevHash: prevHash,
		Version:  "0x01",
	}
	//2.为新生的block，寻找合适的nonce值
	pow := NewPoW(block)
	nonce := pow.Run()
	//3.将block的Nonce设置为找到合适的nonce
	block.Nonce =nonce
	//调用uilt.SHA256Hash进行hash计算
	heightBytes, _ := uilt.IntToBytes(block.Height)
	timeBytes, _ := uilt.IntToBytes(block.TiemeStamp)
	versionBytes := uilt.StringToByte(block.Version)
	nonceBytes, _ := uilt.IntToBytes(block.Nonce)
	//bytes.Join函数，用于[]byte的拼接
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		timeBytes,
		data,
		prevHash,
		versionBytes,
		nonceBytes,
	}, []byte{})
//设置第7个字段hash
	block.Hash = uilt.SH256HashBlock(blockBytes)
	return block

}
/**
区块序列化
 */
func (bk Block) Serialize() ([]byte, error){
	buff := new(bytes.Buffer)
	//Encoder 编码器
	err:=gob.NewEncoder(buff).Encode(bk)
	if err !=nil {
		return nil,err
	}
	return buff.Bytes(),nil
}
/**
反序列化
 */
func DeSerialize(data []byte)(*Block,error){
	var block  Block
	//解码器
	err:=gob.NewDecoder(bytes.NewReader(data)).Decode(&block)
	if err != nil {
		return  nil,err


	}
	return &block,nil
}
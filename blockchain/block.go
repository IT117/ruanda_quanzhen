package blockchain

import (
	"DetaCerProject/uilt"
	"time"
)

/**
区块结构的定义
 */

type Block struct {

Height  int    //区块高度
Size int       //区块大小
TiemeStamp int64//时间戳
Hash   []byte //区块的Hash
Data  []byte //数据
PrevHash  []byte//上一个区块的Hash
Version   string//版本号


}
/**
新建一个区块链实例
 */

func NewBlock(height int,data []byte,prevHash []byte) (Block) {
	block:=Block{
		Height:     height +1,

		TiemeStamp: time.Now().Unix(),

		Data:     data,
		PrevHash:   prevHash,
		Version:    "0x01",
	}
	block.Hash =uilt.SH256HashBlock(block)
	return block

}






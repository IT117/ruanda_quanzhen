package blockchain

import (
	"DetaCerProject/uilt"
	"bytes"
	"crypto/sha256"
	"math/big"
)

const DIFFICU  =16

/**
工作量证明结构体
 */
type ProofOfWork struct {
	//目标值
	Target  *big.Int
	Block   Block
}
/**
实例化一个pow算法实例
 */
func NewPoW(block Block) ProofOfWork{
	target := big.NewInt(1)
	target.Lsh(target,255-DIFFICU)
	pow :=ProofOfWork{
		Target:target,
		Block:block,

	}
	return pow
}
func (p ProofOfWork) Run() (int64) {
	var nonce int64
	bigBlock :=new(big.Int)

	for {
		block:=p.Block
		heightBytes, _ :=uilt.IntToBytes(block.Height)
		timeBytes, _ :=uilt.IntToBytes(block.TiemeStamp)
		versionBytes:=uilt.StringToByte(block.Version)
		nonceBytes,_:=uilt.IntToBytes(nonce)
		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			versionBytes,
			nonceBytes,

		},[]byte{})
		sha256Hash:=sha256.New()
		sha256Hash.Write(blockBytes)
		block256Hash :=sha256Hash.Sum(nil)
		bigBlock =bigBlock.SetBytes(block256Hash)
		if p.Target.Cmp(bigBlock)==1{//如果满足条件时，退出循环
			break
		}
		nonce++  //如果不满足条件，nonce++，继续循环
	}
	return  nonce

}
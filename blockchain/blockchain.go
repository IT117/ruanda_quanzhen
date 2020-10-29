package blockchain

import (
	"fmt"
	"github.com/bolt"
	"math/big"
)

//桶的名称用于该装区块的信息
var BUCKET_NAME = "blocks"

//表示最新区块的KEY
var LAST_KEY = "lasthash"

//储存区块数据的文件
var CHAINDB = "chain.db"

/**
区块链结构体实例定义： 用于代表一条区块链
该区块链包含以下功能：
1.将新产生的区块与已有的区块连接起来，并保存
2.可以查询某个区块的信息
3.可以将所有区块进行遍历，输出区块信息
*/
type BlockChain struct {
	LastHash []byte //最新的区块hash
	BoltDb   *bolt.DB
}

/**
查询所有的区块信息，并返回。将所有的区块放入到切片中。
*/
func (bc BlockChain) QueryAllBlocks() []*Block {
	blocks := make([]*Block, 0)
	db := bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("查询数据错误")

		}
		eachKey := bc.LastHash
		preHashBig := new(big.Int) //初始化一个大整数
		zerpoBig := big.NewInt(0)  //0的大整数
		for {
			eachBlockBytes := bucket.Get(eachKey)
			//拿到区块后反序列化出来查看
			eachBlock, _ := DeSerialize(eachBlockBytes)
			//将遍历到的每一个区块结构体指针放到[]byte容器中
			blocks = append(blocks, eachBlock)
			preHashBig.SetBytes(eachBlock.PrevHash)
			if preHashBig.Cmp(zerpoBig) == 0 { //通过if条件语句判断区块链遍历是否已经遍历到了创世区块
				break
			} //否则，继续向前遍历
			eachKey = eachBlock.PrevHash
		}
		return nil

	})
	return blocks
}

/**
通过区块高度查询某个具体的区块，返回区块实例
*/
func (bc BlockChain) QueryBlockByHeight(height int64) *Block {
	if height < 0 {
		//如果查询的高度小于0，那么就输入有误
		return nil
	}
	var block *Block
	bd := bc.BoltDb
	bd.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("查询数据失败")
		}
		hashKey := bc.LastHash
		for {
			lastBlockBytes := bucket.Get(hashKey)
			eachBlock, _ := DeSerialize(lastBlockBytes)
			if eachBlock.Height < height {
				//给定的高度超过区块链中的高度直接返回
				break
			}
			if eachBlock.Height == height { //高度和目标高度一致，则已经找到了，结束循环
				block = eachBlock
				break

			}
			//遍历当前的区块的高度 和目标高度不一致，继续向前遍历
			//以eachBlock.PrevHash为key，使用get获取上一个区块的数据
			hashKey = eachBlock.PrevHash
			}
			return nil

	})
	return block
}

/**
用于创建一条区块链，并返回该区块链实例
由于区块链就是有一个个的区块组成，因此，如果要创建一条区块链，那么必须要先创建一个区块，该区块
作为该条区块链的创世区块。
*/

func NewBlockChain() BlockChain {
	//0.打开储存区块数据的chanin.db文件
	db, err := bolt.Open("chain.db", 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	var bl BlockChain
	//先查看区块链中的创世区块是否存在
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}

		lastHash := bucket.Get([]byte(LAST_KEY))
		if len(lastHash) == 0 {
			//没有创世区块
			//1、创建创世区块
			genesis := CreateGenesisBlock() //创世区块

			//2.创建储存数据库的文件
			fmt.Printf("genesis的值：%x\n", genesis.Hash)
			bl = BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
			genesisBytes, _ := genesis.Serialize()
			bucket.Put(genesis.Hash, genesisBytes)
			bucket.Put([]byte(LAST_KEY), genesis.Hash)
		} else { //有创世区块
			lastHash := bucket.Get([]byte(LAST_KEY))
			lastBlockBytes := bucket.Get(lastHash) //创世区块的[]byte
			lastBlock, err := DeSerialize(lastBlockBytes)
			if err != nil {
				panic("读取区块链数据失败")
			}
			bl = BlockChain{
				LastHash: lastBlock.Hash,
				BoltDb:   db,
			}

		}

		//3.把新创建的创世区块存入到chain.db当中的一个桶中

		//将创世区块保存到桶中(用序列化)
		//serializeBlock, err := genesis.Serialize()
		//if err != nil {
		//	panic(err.Error())
		//}
		////把创世区块存入到桶里
		//bucket.Put(genesis.Hash, serializeBlock)
		////更新指向最新区块的hash值
		//bucket.Put([]byte(LAST_KEY), genesis.Hash)
		return nil

	})
	return bl

}

/**
调用BlockChain的该SaveBlock 方法，该方法可以将一个生成的新区块链保存到chain.db
*/
func (bc BlockChain) SaveData(data []byte) (Block, error) {
	var e error
	db := bc.BoltDb
	var lastBlock *Block
	//先查一查chain.db中存储的最新的区块
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("boltdb未创建，请重试")
		}
		lastHash := bucket.Get([]byte(LAST_KEY))
		lastBlockByte := bucket.Get(lastHash)
		lastBlock, _ = DeSerialize(lastBlockByte)
		return nil
	})
	//1.先生成一个区块把data存入到新生成的区块中
	newBlock := NewBlock(lastBlock.Height+1, data, lastBlock.Hash)
	//更新chain.db
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//key=hash  value=block的byte
		//区块序列化
		newBlockBytes, _ := newBlock.Serialize()
		//把区块信息保存到bolt.db中
		bucket.Put(newBlock.Hash, newBlockBytes)
		//更新代表最以后一个区块hash值得记录
		bucket.Put([]byte(LAST_KEY), newBlock.Hash)
		bc.LastHash = newBlock.Hash

		return nil

	})
	return newBlock, e

}

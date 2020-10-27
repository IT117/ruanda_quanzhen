package blockchain

import "github.com/bolt"

//桶的名称用于该装区块的信息
var BUCKET_NAME = "blocks"

//表示最新区块的KEY
var LAST_KEY = "lasthash"

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
用于创建一条区块链，并返回该
*/

func NewBlockChain() BlockChain {

	genesis := CreateGenesisBlock() //创世区块
	//2.创建储存数据库的文件
	db, err := bolt.Open("chain.db", 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	bl := BlockChain{
		LastHash: genesis.Hash,
		BoltDb:   db,
	}
	//3.把新创建的创世区块存入到chain.db当中的一个桶中
	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte(BUCKET_NAME))
		if err != nil {
			panic(err.Error())
		}
		//将创世区块保存到桶中(用序列化)
		serializeBlock, err := genesis.Serialize()
		if err != nil {
			panic(err.Error())
		}
		//把创世区块存入到桶里
		bucket.Put(genesis.Hash, serializeBlock)
		//更新指向最新区块的hash值
		bucket.Put([]byte(LAST_KEY), genesis.Hash)
		return nil

	})
	return bl

}

/**
调用BlockChain的该SaveBlock 方法，该方法可以将一个生成的新区块链保存到chain.db
*/
func (bc BlockChain) SaveData(data []byte) {
	db:=bc.BoltDb
	var lastBlock *Block
	//先查一查chain.db中存储的最新的区块
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil{
			panic("boltdb未创建，请重试")
		}
		lastHash:=bucket.Get([]byte(LAST_KEY))
		lastBlockByte:=bucket.Get(lastHash)
		lastBlock,_=DeSerialize(lastBlockByte)
		return nil
	})
	//1.先生成一个区块把data存入到新生成的区块中
	NewBlock(lastBlock.Height+1,data,lastBlock.Hash)
	//更新chain.db
	db.Update(func(tx *bolt.Tx) error {
		return nil

	})
	return

}

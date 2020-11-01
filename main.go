package main

import (
	"DetaCerProject/blockchain"
	"DetaCerProject/db_mysql"
	_ "DetaCerProject/routers"

	"github.com/astaxie/beego"
)

func main() {
	//1.实例化一个区块实例

	blockchain.NewBlockChain()
	//fmt.Printf("最新区块的Hash值：%x\n",bc.LastHash)
	//fmt.Printf("创世区块的Hash值：%x\n",bc.LastHash)
	//block,err:=bc.SaveData([]byte("这里储存链上的数据信息"))
	//if err!=nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//blocks := bc.QueryAllBlocks()
	//if len(blocks)==0{
	//	fmt.Println("暂时没有查询到区块数据")
	//	return
	//}
	// for _,block :=range blocks{
	// 	fmt.Printf("高度：%d,哈希：%x\n,",block.Height,block.Hash,block.PrevHash)
	// }
	//return
	//fmt.Printf("区块的高度：%d\n",block.Height)
	//fmt.Printf("区块的PrevHash：%x\n",block.PrevHash)
	//return
	////1.生成第一个区块链
	//block :=blockchain.CreateGenesisBlock()
	//fmt.Println("第一个区块：",block)
	//fmt.Printf("区块的Hash值：%x",block.Hash)
	//return
	//连接数据库
	db_mysql.ConnectDB()
	//静态资源路径设置
	beego.SetStaticPath("js", "./static/js")
	beego.SetStaticPath("css", "./static/css")
	beego.SetStaticPath("img", "./static/img")

	beego.Run()
}

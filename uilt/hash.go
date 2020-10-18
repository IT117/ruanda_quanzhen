package uilt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

/**
对一个字符串进行MD哈希计算，并返回hash
 */
func MD5HashString(data string) string{
	md5Hash:=md5.New()
	md5Hash.Write([]byte(data))
	passwordBytes:=md5Hash.Sum(nil)
	return hex.EncodeToString(passwordBytes)

}
/**
对io操作的reaer（通常为文件）进行数据读取，并计算hash，返回md5哈希值
 */
func MD5HashReader(reader io.Reader)(string,error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
		return "", err

	}
	md5Hash := md5.New()
	md5Hash.Write(bytes)
	hashBytes:=md5Hash.Sum(nil)
	return  hex.EncodeToString(hashBytes),nil


}
/*
func SH256HashBlock(bolck blockchain.Block)([]byte){

	//1.对block字段进行拼接
	//2.对拼接后的数据进行sha256
	sha256Hash:=sha256.New()
	sha256Hash.Write([]byte(""))
	return  sha256Hash.Sum(nil)
}

 */
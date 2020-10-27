package uilt

import (
	"bytes"
	"encoding/binary"
)

/**
int转[]byte
 */
func IntToBytes(num int64)([]byte,error){
	//大端位序排列binary.BigEndian
	//小端位序排序binary.LittleEndian
	buff:=new(bytes.Buffer)
	err:=binary.Write(buff,binary.BigEndian,num)
	if err!=nil{
		return nil, err
	}
	return buff.Bytes(),nil



}
/**
string转[]byte
 */
func StringToByte(st string)[]byte{
	return []byte(st)
}
package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

//实现int64转[]byte
func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, data)
	if nil != err {
		log.Panicf("int transact to []byte failed %v/n", err)
	}
	return buffer.Bytes()
}

//标准JSON格式转切片
//执行命令格式
// send -from [\"huhu\"] -to [\"jiji\"] -amount [\"100\"]
func JSONToSlice(jsonString string) []string {
	var strSlice []string
	//json
	if err := json.Unmarshal([]byte(jsonString), &strSlice); nil != err {
		log.Panicf("json to []string failed! %v\n", err)
	}
	return strSlice
}

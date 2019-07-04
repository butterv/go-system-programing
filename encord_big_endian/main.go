package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 32ビットのビッグエンディアンのデータ(10000)
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	// エンディアンの変換
	if err := binary.Read(bytes.NewReader(data), binary.BigEndian, &i); err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	fmt.Printf("data: %d\n", i)
}
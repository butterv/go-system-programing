package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func dummyChunk(chunk io.Reader) {
	var length int32
	if err := binary.Read(chunk, binary.BigEndian, &length); err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	buffer := make([]byte, 4)
	if _, err := chunk.Read(buffer); err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
	// チャンクを格納する配列
	var chunks []io.Reader
	// 最初の8バイトを飛ばす
	if _, err := file.Seek(8, 0) ; err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return chunks
	}

	var offset int64 = 8
	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}

		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
		// 次のチャンクの先頭に移動
		// 現在位置は長さを読み終わった箇所なので
		// チャンク名(4バイト) + データ長 + CRC(4バイト)先に移動
		offset, _ = file.Seek(int64(length + 8), 1)
	}
	return chunks
}

func main() {
	file, err := os.Open(".png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dummyChunk(chunk)
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type bencode struct {
	bufio.Reader
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (bencode *bencode) genLen() int {
	len, err := bencode.ReadSlice(':')
	check(err)

	s := strings.TrimSuffix(string(len), ":")

	i, _ := strconv.Atoi(s)

	return i
}

func (bencode *bencode) readItem() {
	var key string
	//var value string

	for {

		len := bencode.genLen()
		fmt.Println(len)
		b := make([]byte, len)
		_, _ = io.ReadFull(bencode, b)
		key = string(b)
		fmt.Println(key)
	}

}

func (bencode *bencode) readDictionary() {
	bencode.readItem()
	/*
		d := make(map[string]interface{})
		b, err := bencode.ReadByte()
		check(err)

		fmt.Println(string(b))
	*/
}

func readInteger() {

}

func readBytes() {

}

func readList() {

}

func main() {
	f, err := os.Open("debian.torrent")
	check(err)

	bencode := bencode{*bufio.NewReader(f)}
	if b, err := bencode.ReadByte(); err != nil {
		fmt.Println("Not bencode! %s", b)
	}

	bencode.readDictionary()

}

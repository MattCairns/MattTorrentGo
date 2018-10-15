package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
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

func (bencode *bencode) getLen() int {
	len, err := bencode.ReadSlice(':')
	check(err)

	s := strings.TrimSuffix(string(len), ":")

	i, _ := strconv.Atoi(s)

	return i
}

func (bencode *bencode) checkType() interface{} {
	switch b, _ := bencode.ReadByte(); b {
	case 'd':
		color.Red("Dictionary")
		return bencode.readItem()
	case 'i':
		color.Red("Integer")
		return bencode.readInteger()
	case 'l':
		color.Red("List")
		return bencode.readList()
	default:
		err := bencode.UnreadByte()
		check(err)
		return 1
	}
}

func (bencode *bencode) readItem() int {
	var key string
	//var value string

	for {

		len := bencode.getLen()
		b := make([]byte, len)
		_, _ = io.ReadFull(bencode, b)
		key = string(b)
		fmt.Println(key)

		bencode.checkType()
	}

	return 1

}

func (bencode *bencode) readDictionary() interface{} {
	bencode.readItem()
	/*
		d := make(map[string]interface{})
		b, err := bencode.ReadByte()
		check(err)

		fmt.Println(string(b))
	*/

	return nil
}

func (bencode *bencode) readInteger() int {
	byteString, err := bencode.ReadSlice('e')
	check(err)

	s := strings.TrimSuffix(string(byteString), "e")

	i, err := strconv.Atoi(s)

	fmt.Println(i)

	return i

}

func (bencode *bencode) readBytes() int {
	return 1
}

func (bencode *bencode) readList() int {
	for {
		len := bencode.getLen()
		b := make([]byte, len)
		_, _ = io.ReadFull(bencode, b)
		key := string(b)
		fmt.Println(key)

		if b2, err := bencode.ReadByte(); b2 == 'e' {
			check(err)
			return 1

		} else {
			err := bencode.UnreadByte()
			check(err)
		}

	}

	return 1
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

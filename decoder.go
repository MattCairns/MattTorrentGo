package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
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
	time.Sleep(2 * time.Second)
	color.Magenta(s)

	i, _ := strconv.Atoi(s)

	return i
}

func (bencode *bencode) checkType() interface{} {
	switch b, _ := bencode.ReadByte(); b {
	case 'd':
		return bencode.readItem()
	case 'i':
		return bencode.readInteger()
	case 'l':
		return bencode.readList()
	default:
		color.Green(string(b))
		err := bencode.UnreadByte()
		check(err)
		return 1
	}
}

func (bencode *bencode) readItem() int {
	var key string
	for {
		bencode.checkType()

		len := bencode.getLen()
		b := make([]byte, len)
		_, _ = io.ReadFull(bencode, b)
		key = string(b)
		fmt.Println(key)

		if bencode.isEnd() {
			return 1
		}

	}

	return 1

}

func (bencode *bencode) readInteger() int {
	byteString, err := bencode.ReadSlice('e')
	check(err)

	s := strings.TrimSuffix(string(byteString), "e")

	i, err := strconv.Atoi(s)

	fmt.Println(i)

	return i

}

func (bencode *bencode) readList() int {
	for {
		color.Green("readList")
		bencode.checkType()
		len := bencode.getLen()
		b := make([]byte, len)
		_, _ = io.ReadFull(bencode, b)
		key := string(b)
		fmt.Println(key)

		if bencode.isEnd() {
			return 1
		}
	}

	return 1
}

func (bencode *bencode) readBytes() int {
	return 1
}

func (bencode *bencode) readDictionary() interface{} {
	bencode.readItem()
	return nil
}

func (bencode *bencode) isEnd() bool {
	if b, err := bencode.ReadByte(); b == 'e' {
		check(err)
		return true
	}

	bencode.UnreadByte()
	return false
}

func main() {
	f, err := os.Open("ubuntu.torrent")
	check(err)

	bencode := bencode{*bufio.NewReader(f)}
	if b, err := bencode.ReadByte(); err != nil {
		fmt.Println("Not bencode! %s", b)
	} else {
		color.Red(string(b))
	}

	bencode.readDictionary()

}

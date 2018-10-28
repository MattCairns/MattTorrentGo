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
		color.Red("readDictionary")
		return bencode.readDictionary()
	case 'i':
		color.Red("readInteger")
		return bencode.readInteger()
	case 'l':
		color.Red("readList")
		return bencode.readList()
	default:
		err := bencode.UnreadByte()
		check(err)
		return nil
	}
}

func (bencode *bencode) readInteger() interface{} {
	byteString, err := bencode.ReadSlice('e')
	check(err)

	s := strings.TrimSuffix(string(byteString), "e")

	i, err := strconv.Atoi(s)

	return i

}

func (bencode *bencode) readList() []interface{} {
	var l []interface{}
	for {
		color.Green("readList")
		bencode.checkType()
		len := bencode.getLen()
		b := make([]byte, len)
		_, _ = io.ReadFull(bencode, b)
		key := string(b)

		l = append(l, key)

		if bencode.isEnd() {
			return l
		}
	}

	return l
}

func (bencode *bencode) readBytes() int {
	return 1
}

func (bencode *bencode) readDictionary() map[string]interface{} {
	var key string
	var value interface{}

	m := make(map[string]interface{})
	for {

		len := bencode.getLen()
		b := make([]byte, len)
		_, _ = io.ReadFull(bencode, b)
		key = string(b)
		color.Magenta(key)

		value = bencode.checkType()
		fmt.Println(value)

		if value != nil {
			m[key] = value
		}

		if bencode.isEnd() {
			return m
		}

	}

	return m
}

func (bencode *bencode) isEnd() bool {
	if b, err := bencode.ReadByte(); b == 'e' {
		check(err)
		return true
	}

	bencode.UnreadByte()
	return false
}

func decode(torrent string) map[string]interface{} {
	f, err := os.Open("debianC.torrent")
	check(err)

	bencode := bencode{*bufio.NewReader(f)}
	if b, err := bencode.ReadByte(); err != nil {
		fmt.Println("Not bencode! %s", b)
	} else {
		color.Red(string(b))
	}

	data := bencode.readDictionary()
	fmt.Println(data)
	return data
}

func main() {
	decode("ubuntu.torrent")
}

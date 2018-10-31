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
		return bencode.readDictionary()
	case 'i':
		return bencode.readInteger()
	case 'l':
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
		key := bencode.getKey()

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
		key = bencode.getKey()

		value = bencode.checkType()
		if value == nil {
			value = bencode.getKey()
		}

		m[key] = value

		if bencode.isEnd() {
			return m
		}

	}

	return m
}

func (bencode *bencode) getKey() string {
	var key string
	len := bencode.getLen()
	b := make([]byte, len)
	_, _ = io.ReadFull(bencode, b)
	key = string(b)

	return key

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
	f, err := os.Open(torrent)
	check(err)

	bencode := bencode{*bufio.NewReader(f)}
	if b, err := bencode.ReadByte(); err != nil {
		fmt.Println("Not bencode! %s", b)
	}

	data := bencode.readDictionary()
	return data
}

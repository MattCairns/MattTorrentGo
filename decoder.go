package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readDictionary(r *bufio.Reader) {
	b, err := r.ReadByte()
	check(err)

	fmt.Println(string(b))
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

	read := bufio.NewReader(f)

	b, err := read.ReadByte()
	check(err)
	if string(b) == "d" {
		fmt.Println("Dict")
		d_size, err := read.ReadByte()
		check(err)
		fmt.Printf("Dict size %s", string(d_size))
	}
}

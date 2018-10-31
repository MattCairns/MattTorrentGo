package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInfoHash(filename string) string {
	buffer, err := ioutil.ReadFile(filename)
	check(err)

	s := string(buffer)

	r, err := regexp.Compile("info(.+)$")
	check(err)

	r.FindString(s)

	h := sha1.New()

	h.Write([]byte(s))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

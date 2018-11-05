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

	fmt.Println(s)
	r, err := regexp.Compile("info(.+)ee$")
	check(err)

	rs := r.FindStringSubmatch(s)

	h := sha1.New()

	h.Write([]byte(rs[0]))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

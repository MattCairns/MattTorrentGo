package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"os"
)

func main() {
	m := decode("debianC.torrent")

	fmt.Println(m["announce"])

	response, err := http.Get(m["announce"].(string))
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		check(err)
		fmt.Println("%s\n", string(contents))
	}
}

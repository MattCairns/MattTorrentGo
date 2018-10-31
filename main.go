package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	file := "debian.torrent"
	m := decode(file)
	left := url.QueryEscape(strconv.Itoa(m["info"].(map[string]interface{})["length"].(int)))
	info_hash := url.QueryEscape(string(getInfoHash(file)))
	peer_id := url.QueryEscape("MT20-111111111111111")

	fmt.Println(m["announce"])
	fmt.Println(m["info"].(map[string]interface{})["length"])

	t := fmt.Sprintf("%s?info_hash=%s&peer_id=%s&left=%s", m["announce"], info_hash, peer_id, left)
	fmt.Println(t)

	response, err := http.Get(t)
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

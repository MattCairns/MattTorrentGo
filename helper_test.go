package main

import (
	"testing"
)

func TestGetInfoHash(t *testing.T) {
	file := "torrents/alaska.torrent"
	expected := "3d03ce50e2ef38b5519705ffc28d5f260c5d15ee"
	actual := getInfoHash(file)

	if actual != expected {
		t.Errorf("Expected hash %s but got %s", expected, actual)
	}
}

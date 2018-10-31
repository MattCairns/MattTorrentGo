package main

import (
	"testing"
)

func TestGetInfoHash(t *testing.T) {
	file := "debian.torrent"
	expected := "95892469229115cb6f5addc1f6f7953e1eb6cef9"

	actual := getInfoHash(file)

	if actual != expected {
		t.Errorf("Expected hash %s but got %s", expected, actual)
	}
}

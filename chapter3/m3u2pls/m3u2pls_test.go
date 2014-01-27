package main

import (
	"testing"
	"io/ioutil"
)

func TestReadM3u(t *testing.T) {
	rawBytes, _ := ioutil.ReadFile("data/m3u_input.m3u")
	songs := readM3uPlaylist(string(rawBytes))
	if songs[0].Title != "Allah-Las - Catamaran" {
		t.Errorf("Song title doesn't match. Got: %s Expected %s", songs[0].Title, "Allah-Las - Catamaran") 
	}
}

func TestReadPls(t *testing.T) {
	rawBytes, _ := ioutil.ReadFile("data/pls_input.pls")
	songs := readPlsPlaylist(string(rawBytes))
	if songs[0].Title != "Allah-Las - Catamaran" {
		t.Errorf("Song title doesn't match. Got: %s Expected %s", songs[0].Title, "Allah-Las - Catamaran") 
	}
}

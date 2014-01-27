package main

import (
	"testing"
)

func TestM3u2Pls(t *testing.T) {
	if rawBytes, err:= ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		songs := readM3uPlaylist(string(rawBytes))
		writePlsPlaylist(songs)
	}
}

func TestPls2M3u(t *testing.T) {

}

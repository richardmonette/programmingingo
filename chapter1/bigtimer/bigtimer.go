package main

import (
	"log"
	"io/ioutil"
	"strings"
	"fmt"
)

func main() {
	var bigDigits [7][10]string
	
	content, err := ioutil.ReadFile("digits.txt")
	if err != nil {
		log.Fatal(err)
	}
	stringcontent := string(content)
	lines := strings.Split(stringcontent, "\n")
	fmt.Println(lines)
}
package main

import (
	"log"
	"io/ioutil"
	"strings"
	"fmt"
	//"strconv"
)

func main() {

	content, err := ioutil.ReadFile("digits.txt")
	if err != nil {
		log.Fatal(err)
	}

	stringcontent := string(content)
	lines := strings.Split(stringcontent, "\r\n")
	
	var digits [70][7]string
	
	for i, line := range lines {
		for x := 0; x < 7; x++ {
			digits[i][x] = string(line[x])
			//fmt.Print(string(line[x]))
		}
		//fmt.Print("\n")
	}
	
	toDraw := []int{1, 3, 3, 7}

	for x := 0; x < 7; x++ {
		for _, value := range toDraw {
			line := digits[value * 7 + x]
			for _, char := range line {
				fmt.Print(char)
			}
		}
		fmt.Print("\n")
	}
	
	/*for i, _ := range lines {
		for x := 0; x < 7; x++ {
			fmt.Print(digits[i][x])
		}
		fmt.Print("\n")
	}*/

}
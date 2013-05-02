package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "World!" // := operator is a 'short variable declaration' which declares and initalizes a variable at the same time
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ") // note how having already initialized with the := beforehand, we can just assign now
	}
	fmt.Println("Hello", who)
}
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var d string = "hello world"
	//var r string
	for i := 0; i < 5; i++ {
		r := sha256.Sum256([]byte(d + string(i)))
		fmt.Println(i, " - sha256: ", string(r[:]))
	}
}

package main

import (
	"fmt"
	"crypto/sha1"
)

func main() {
	h := sha1.New()
	h.Write([]byte("Future crypto-text"))
	
	crypto := h.Sum([]byte{})
	fmt.Println(crypto)
}

package main

import(
	"fmt"
)

func main() {
	for i := 0;  i <= 10; i++ {
		go fmt.Println("Hello World:", i)
	}

	var input string
	fmt.Scanln(&input)
}
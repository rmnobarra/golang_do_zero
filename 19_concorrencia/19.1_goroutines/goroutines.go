package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Hell'yo")
	go escrever("Yahooo")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

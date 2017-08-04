package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		help(args[0])
		return
	}

}

func help(name string) {
	fmt.Println("This program calculates n-th digit of binary log2 or hexadecimal pi.")
	fmt.Println("Usage:", name, "[log2/pi] n")
}

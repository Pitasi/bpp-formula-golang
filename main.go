package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
	"time"
)
import b "./bppformula"

func main() {
	filename, args := os.Args[0], os.Args[1:]

	if len(args) != 2 || (args[0] != "log2" && args[0] != "pi") {
		help(filename)
		return
	}

	d, err := parseInt(filename, args[1])
	if (err != nil) { return }


	var res string
	fmt.Printf("Computing...")
	timeStart := time.Now()
	if (args[0] == "log2") {
		res, err = b.ToStringBase(b.Log2(d), 2, 8)
		fmt.Printf("\rBinary")
	} else if (args[0] == "pi") {
		res, err = b.ToStringBase(b.Pi(d), 16, 8)
		fmt.Printf("\rHexadecimal")
	}
	elapsed := time.Since(timeStart)

	fmt.Printf(" digits of %s starting at position %s: ", args[0], args[1])
	if (err != nil) {
		fmt.Println("Unexpected error:", err)
	}
	fmt.Println(res)
	fmt.Println("Computation took", elapsed)
}

func help(filename string) {
	fmt.Println("This program calculates n-th digit of binary log2 or hexadecimal pi.")
	fmt.Println("Usage:", filename, "[log2/pi] n")
}

func parseInt(filename string, s string) (int, error) {
	d, err := strconv.Atoi(s)
	if (err != nil || d <= 0) {
		fmt.Println("ERROR Invalid number n provided. A positive integer is required.")
		fmt.Println("Try running:", filename, "pi 10000")
		return -1, errors.New("Invalid number")
	}
	return d, nil
}
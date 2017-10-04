package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	b "./bppformula"
)

func main() {
	filename, args := os.Args[0], os.Args[1:]

	invalidArguments := len(args) != 2 || (args[0] != "log2" && args[0] != "pi")
	if invalidArguments {
		help(filename)
		return
	}

	// Parse second argument as integer
	d, err := strconv.Atoi(args[1])
	if err != nil || d <= 0 {
		fmt.Println("ERROR Invalid number n provided. A positive integer is required.")
		fmt.Println("Try running:", filename, "pi 10000")
		return
	}

	// Start computing
	var res string
	fmt.Printf("Computing...")
	timeStart := time.Now()

	// Case switch for log2 or pi
	if args[0] == "log2" {
		res, err = b.ToStringBase(b.Log2(d), 2, 8)
		fmt.Printf("\rBinary")
	} else if args[0] == "pi" {
		res, err = b.ToStringBase(b.Pi(d), 16, 8)
		fmt.Printf("\rHexadecimal")
	}

	// Print results
	elapsed := time.Since(timeStart)
	fmt.Printf(" digits of %s starting at position %s: ", args[0], args[1])
	if err != nil {
		fmt.Println("Unexpected error:", err)
	}
	fmt.Println(res)
	fmt.Println("Computation took", elapsed)
}

func help(filename string) {
	fmt.Println("This program calculates n-th digit of binary log2 or hexadecimal pi.")
	fmt.Println("Usage:", filename, "[log2/pi] <n>")
}

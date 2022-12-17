package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	// create csv file from puzzle input using
	// find & replace consisting of columns :
	// [ n, from, to ]

	// create representation for ~ starting stacks of crates
	var stacks []string

	stacks = append(stacks, "")         // 0
	stacks = append(stacks, "BPNQHDRT") // 1
	stacks = append(stacks, "WGBJTV00") // 2
	stacks = append(stacks, "NRHDSVMQ") // 3
	stacks = append(stacks, "PZNMC")    // 4
	stacks = append(stacks, "BZD")      // 5
	stacks = append(stacks, "VCWZ")     // 6
	stacks = append(stacks, "GZNCVQLS") // 7
	stacks = append(stacks, "LGJMDNV")  // 8
	stacks = append(stacks, "TPMFZCG")  // 9

	// open puzzle-input, close after use
	pattern, _ := os.Open("puzzle-input.csv")
	defer pattern.Close()

	// prepare to read procedures
	r := csv.NewReader(pattern)

	// create required vars
	n, from, to := 0, 0, 0

	for {
		// read a record
		rec, err := r.Read()
		// until,
		if err == io.EOF {
			break
		}

		n, _ = strconv.Atoi(rec[0])
		from, _ = strconv.Atoi(rec[1])
		to, _ = strconv.Atoi(rec[2])

		// UPDATE THIS !!
		crateMover := "9001" // for part 1 or "9001" for part 2

		// add to desired stack
		switch crateMover {
		case "9000":
			stacks[to] += getReverse(stacks[from][len(stacks[from])-n:])
		case "9001":
			stacks[to] += stacks[from][len(stacks[from])-n:]
		}
		// remove from desired stack
		stacks[from] = stacks[from][:len(stacks[from])-n]

		// visualise
		visualise(stacks, n, from, to)

	}
}

func getReverse(s string) (result string) {
	for _, i := range s {
		result = string(i) + result
	}
	return
}

func visualise(stacks []string, n, from, to int) {
	// clear screen
	fmt.Print("\033[H\033[2J")

	// print procedure
	fmt.Printf("[ Move %d from %d to %d ]\n", n, from, to)

	// print stacks
	for i := 1; i < len(stacks); i++ {
		fmt.Printf("%d : %v\n", i, stacks[i])
	}
	// wait ...
	time.Sleep(5 * time.Millisecond)
}

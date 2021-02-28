package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(bytes []byte) {
	n := len(bytes)
	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-i-1] = bytes[n-i-1], bytes[i]
	}
}

func recat(in *os.File, w *bufio.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		reverse(bytes)
		w.Write(bytes)
		w.WriteByte(0x0A) // \n
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "recat: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	args := os.Args[1:]

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	if len(args) > 0 {
		for _, in := range args {
			file, err := os.Open(in)
			if err != nil {
				fmt.Fprintf(os.Stderr, "recat: %s\n", err)
				os.Exit(1)
			}
			defer file.Close()
			recat(file, writer)
		}
		return
	}

	recat(os.Stdin, writer)
}

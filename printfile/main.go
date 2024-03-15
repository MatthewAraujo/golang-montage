package main

import (
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)

}

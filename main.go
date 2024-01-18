package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		fmt.Printf("%s\n", line)
	}

}

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
		if len(line) == 0 {
			continue
		}
		line, abbrev := extractAbbrev(line)
		fmt.Printf("%s - %s\n", line, abbrev)
	}

}

func extractAbbrev(s string) (string, string) {
	line := ""
	abbrev := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] <= 'z' && s[i] >= 'A' {
			abbrev = abbrev + string(s[i])
		} else {
			s = s[0:i]
			break
		}
	}
	abbrevR := reverse(abbrev)
	line = strings.TrimSpace(s)
	return line, abbrevR
}

func reverse(s string) string {
	answer := ""
	for i := len(s) - 1; i >= 0; i-- {
		answer = answer + string(s[i])
	}
	return answer
}

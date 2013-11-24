package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func HasSquare(lines []string) bool {
	start, end, count := -1, -1, -1
	consecCount := 0
	blank := false
	for _, line := range lines {
		s, e, c := strings.Index(line, "#"), strings.LastIndex(line, "#"), strings.Count(line, "#")

		if start == -1 {
			// not started yet
			if s == -1 {
				continue
			}
			// first line
			start, end, count = s, e, c
			consecCount += 1
		} else {
			// blank line
			if s == -1 {
				blank = true
				continue
			}
			// found a matching line
			if s == start && e == end && c == count && blank != true {
				consecCount += 1
				continue
			}
			return false
		}
	}
	if consecCount == count && count == end-start+1 {
		return true
	}
	return false
}

func Run(r io.Reader) []bool {
	answers := []bool{}
	scanner := bufio.NewScanner(r)

	// read number of cases
	scanner.Scan()
	numCases, _ := strconv.ParseInt(scanner.Text(), 10, 8)

	for c := 0; c < int(numCases); c++ {
		// read number of lines in this case
		scanner.Scan()
		numLines, _ := strconv.ParseInt(scanner.Text(), 10, 8)

		lines := []string{}
		for l := 0; l < int(numLines); l++ {
			scanner.Scan()
			lines = append(lines, scanner.Text())
		}
		hasSquare := HasSquare(lines)
		answers = append(answers, hasSquare)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return answers
}

func main() {
	answers := Run(os.Stdin)
	for a, hasSquare := range answers {
		ans := "NO"
		if hasSquare {
			ans = "YES"
		}
		fmt.Printf("Case #%d: %s\n", a+1, ans)
	}

}

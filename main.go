// go-sort/main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
)

func readLines(scanner *bufio.Scanner) []string {
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func removeDuplicates(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	result := []string{lines[0]}
	for i := 1; i < len(lines); i++ {
		if lines[i] != lines[i-1] {
			result = append(result, lines[i])
		}
	}
	return result
}

func main() {
	// -u 옵션 추가
	unique := flag.Bool("u", false, "중복 제거")
	flag.Parse()

	var lines []string

	if flag.NArg() > 0 {
		// 파일 입력
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "파일 열기 실패: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lines = readLines(scanner)
	} else {
		// 표준 입력
		scanner := bufio.NewScanner(os.Stdin)
		lines = readLines(scanner)
	}

	sort.Strings(lines)

	if *unique {
		lines = removeDuplicates(lines)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

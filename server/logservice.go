package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func scanLogs(query string) ([]string, int32) {
	file, err := os.Open("../sample.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count int32
	var results []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // GET the line string
		if strings.Contains(line, query) {
			fmt.Println(line)
			count += 1
			results = append(results, line)
		}
	}
	return results, count
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileTextLines {
		// fmt.Println(line)
		stringSplit := strings.Split(line, ": ")
		regex, password := stringSplit[0], stringSplit[1]
		fmt.Printf("Regex: '%s'; Password: '%s'\n", regex, password)
	}

}

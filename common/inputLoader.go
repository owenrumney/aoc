package common

import (
	"bufio"
	"log"
	"os"
)

func GetScanner(inputFile string) *bufio.Scanner {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	// create a scanner to open the file
	scanner := bufio.NewScanner(file)
	return scanner
}

func LoadInput(inputFile string) string {
	scanner := GetScanner(inputFile)
	content := ""
	for scanner.Scan() {
		content = scanner.Text()
	}
	return content
}

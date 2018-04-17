package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	length := 4
	wordlist, err := readLines("/usr/share/dict/words")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	rand.Seed(time.Now().Unix())
	for length > 0 {
		fmt.Print(wordlist[rand.Intn(len(wordlist))], " ")
		length--
	}
	fmt.Println()

}

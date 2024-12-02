package shared

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func ReadTextFile(filename string) []string {
	var lines []string

	cwd, err := os.Getwd()
	if err != nil {
		log.Println("Error getting working directory:", err)
		return nil
	}

	fullpath := filepath.Join(cwd, filename)
	file, err := os.Open(fullpath)
	if err != nil {
		log.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	log.Printf("Reading input file from '%s'", fullpath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}

	return lines
}

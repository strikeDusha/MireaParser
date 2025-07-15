package main

import (
	"bufio"
	"log"
	"os"
	"vuzparser/parser"
)

func main() {
	lines, err := readLines("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	pages := make([]*parser.Page, 0, len(lines))
	for _, v := range lines {
		f, err := parser.GetTable(v)
		if err != nil {
			log.Fatal(err)
		}
		pages = append(pages, f)
	}
	parser.MultiSheetExcelFile(pages, "mireafiles")

}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

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
	channel := make(chan *parser.Page, len(lines))
	for _, v := range lines[:] {
		go parser.GetTable(v, channel) // about 5 seconds for
	}
	for i := 0; i < len(lines); i++ {
		v := <-channel
		pages = append(pages, v)
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

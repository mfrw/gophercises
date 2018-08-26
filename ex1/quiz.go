package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	qfile = flag.String("file", "problems.csv", "Input CSV file for questions")
)

func quiz(r io.Reader, w io.Writer) (int, int) {
	var correct, incorrect int
	var ans string
	r = bufio.NewReader(r)
	csvreader := csv.NewReader(r)

	for {
		record, err := csvreader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record) != 2 {
			continue
		}

		fmt.Fprintf(w, "%s: ", record[0])
		fmt.Scanf("%s", &ans)
		if ans == strings.TrimSpace(record[1]) {
			correct++
		} else {
			incorrect++
		}
	}
	return correct, incorrect

}

func main() {
	flag.Parse()
	f, err := os.Open(*qfile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	correct, incorrect := quiz(f, os.Stdout)
	fmt.Println("Total Correct: ", correct, "Out of:", incorrect+correct)
}

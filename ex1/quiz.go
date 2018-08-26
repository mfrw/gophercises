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
	"time"
)

var (
	qfile = flag.String("file", "problems.csv", "Input CSV file for questions")
	qtime = flag.Int("time", 10, "default quiz time")
)

type Quiz struct {
	q string
	a string
}

func getQuestions(r io.Reader) chan Quiz {
	qchan := make(chan Quiz)

	r = bufio.NewReader(r) // Interface Chaining
	csvreader := csv.NewReader(r)

	go func() {
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
			qchan <- Quiz{q: strings.TrimSpace(record[0]), a: strings.TrimSpace(record[1])}
		}
		close(qchan)
	}()

	return qchan
}

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

func takeQuiz(questions <-chan Quiz, done <-chan time.Time, w io.Writer) (int, int) {
	var correct, total int
	var ans string
	run := true
	for run {
		select {
		case q := <-questions:
			fmt.Fprintf(w, "%s: ", q.q)
			fmt.Scanf("%s", &ans)
			if ans == q.a {
				correct++
			}
			total++
		case <-done:
			run = false
		}
	}
	return correct, total
}

func main() {
	flag.Parse()
	f, err := os.Open(*qfile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// correct, incorrect := quiz(f, os.Stdout)
	// fmt.Println("Total Correct: ", correct, "Out of:", incorrect+correct)
	questions := getQuestions(f)
	ticker := time.After(time.Duration(*qtime) * time.Second)
	correct, total := takeQuiz(questions, ticker, os.Stdout)
	fmt.Println("Total Correct: ", correct, "Out of:", total)
}

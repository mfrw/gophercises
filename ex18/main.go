package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func trackTime(s time.Time, msg string) {
	e := time.Since(s)
	fmt.Println(msg, ":", e)
}

func main() {
	b, err := primitive("test.png", "out.png", 10, ModeTriangle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}

type Mode int

const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedRect
	ModeBeziers
	ModeRotatedEllipse
	ModePloygon
)

func primitive(inputFile, outputFile string, numShapes int, mode Mode) (string, error) {
	defer trackTime(time.Now(), "PRIMITIVE")
	inp := fmt.Sprintf("-i %s -o %s -n %d -m %d", inputFile, outputFile, numShapes, mode)
	cmd := exec.Command("primitive", strings.Fields(inp)...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}

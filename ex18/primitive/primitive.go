package primitive

import (
	"fmt"
	"os/exec"
	"strings"
)

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

func Primitive(inputFile, outputFile string, numShapes int, mode Mode) (string, error) {
	inp := fmt.Sprintf("-i %s -o %s -n %d -m %d", inputFile, outputFile, numShapes, mode)
	cmd := exec.Command("primitive", strings.Fields(inp)...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}

package primitive

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

func WithMode(mode Mode) func() []string {
	return func() []string {
		return []string{"-n", fmt.Sprintf("%d", mode)}
	}
}

func Transform(image io.Reader, opts ...func() []string) (io.Reader, error) {
	in, err := ioutil.TempFile("", "in_")
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	out, err := ioutil.TempFile("", "out_")
	if err != nil {
		return nil, err
	}
	defer os.Remove(out.Name())

	_, err = io.Copy(in, image)
	if err != nil {
		return nil, err
	}
	std, err := primitive(in.Name(), out.Name(), numShapes, ModeCombo)
	if err != nil {
		return nil, err
	}
	fmt.Println(std)
	var b bytes.Buffer
	_, err = io.Copy(&b, out)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func primitive(inputFile, outputFile string, numShapes int, mode Mode) (string, error) {
	inp := fmt.Sprintf("-i %s -o %s -n %d -m %d", inputFile, outputFile, numShapes, mode)
	cmd := exec.Command("primitive", strings.Fields(inp)...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}

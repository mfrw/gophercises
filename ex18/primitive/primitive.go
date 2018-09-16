package primitive

import (
	"bytes"
	"errors"
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

func Transform(image io.Reader, ext string, numShapes int, opts ...func() []string) (io.Reader, error) {
	in, err := tempFile("in_", ext)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary input file")
	}
	defer os.Remove(in.Name())
	out, err := tempFile("in_", ext)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary input file")
	}
	defer os.Remove(out.Name())

	_, err = io.Copy(in, image)
	if err != nil {
		return nil, errors.New("primitive: failed to copy image to temporary input file")
	}
	std, err := primitive(in.Name(), out.Name(), numShapes, ModeCombo)
	if err != nil {
		return nil, fmt.Errorf("primitive: failed to run the primitive command, std=%s", std)
	}
	fmt.Println(std)
	var b bytes.Buffer
	_, err = io.Copy(&b, out)
	if err != nil {
		return nil, errors.New("primitive: failed to copy output file into byte buffer")
	}
	return &b, nil
}

func primitive(inputFile, outputFile string, numShapes int, mode Mode) (string, error) {
	inp := fmt.Sprintf("-i %s -o %s -n %d -m %d", inputFile, outputFile, numShapes, mode)
	cmd := exec.Command("primitive", strings.Fields(inp)...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}

func tempFile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("", "in_")
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary file")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}

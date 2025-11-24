package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type ElementID int
type NetID int

type Element struct {
	Name int `json:"name"`
	X    int `json:"x"`
	Y    int `json:"y"`
}

type ParseResult struct {
	ElemNames   []string            `json:"elem_names"`
	NetNames    []string            `json:"net_names"`
	Q           map[int]map[int]int `json:"Q"`
	R           map[int]map[int]int `json:"R"`
	Plot        []Element           `json:"plot"`
	OptPlot     []Element           `json:"opt_plot"`
	D           [][]float64         `json:"D"`
	NetElements map[int][]Element   `json:"NetElements"`
	FBefore     float64             `json:"F_before"`
	FAfter      float64             `json:"F_after"`
}

type Parser interface {
	Parse(r io.Reader) (*ParseResult, error)
	Name() string
}

var (
	// Calay: DD1(7)
	calayRe = regexp.MustCompile(`([A-Za-z0-9_]+)\(('?\d+)\)`)
	// Allegro: DD1.7
	allegroRe = regexp.MustCompile(`([Q-Za-z0-9_]+)\.(\d+)`)
	// Calay header: 107 or N00605
	calayHeaderRe = regexp.MustCompile(`^([^\s,;]+)`)
)

func DetectFormat(r io.Reader) (Parser, io.Reader, error) {
	const sniffLines = 20
	var lines []string
	br := bufio.NewReader(r)
	var buf bytes.Buffer

	for i := 0; i < sniffLines; i++ {
		l, err := br.ReadString('\n')
		if l != "" {
			lines = append(lines, l)
			buf.WriteString(l)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}
	}

	sample := strings.Join(lines, "\n")

	var parser Parser
	switch {
	case strings.Contains(sample, "$PACKAGES") || strings.Contains(sample, "$NETS"):
		parser = NewAllegroParser()
	case calayRe.MatchString(sample):
		parser = NewCalayParser()
	default:
		return nil, nil, fmt.Errorf("unknown format")
	}

	fullReader := io.MultiReader(bytes.NewReader(buf.Bytes()), br)
	return parser, fullReader, nil
}

func ParseAuto(r io.Reader) (*ParseResult, error) {
	parser, fullReader, err := DetectFormat(r)
	if err != nil {
		return nil, err
	}
	return parser.Parse(fullReader)

}

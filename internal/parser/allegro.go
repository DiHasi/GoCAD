package parser

import (
	"bufio"
	"io"
	"strings"
)

type AllegroParser struct {
	elemNames []string
	netNames  []string
	elemIndex map[string]int
	netIndex  map[string]int
	Q         map[int]map[int]int
}

func NewAllegroParser() *AllegroParser {
	return &AllegroParser{
		elemNames: []string{},
		netNames:  []string{},
		elemIndex: make(map[string]int),
		netIndex:  make(map[string]int),
		Q:         make(map[int]map[int]int),
	}
}

func (p *AllegroParser) Name() string { return "Allegro" }

func (p *AllegroParser) ensureElem(name string) int {
	if id, ok := p.elemIndex[name]; ok {
		return id
	}
	id := len(p.elemNames)
	p.elemIndex[name] = id
	p.elemNames = append(p.elemNames, name)
	if _, ok := p.Q[id]; !ok {
		p.Q[id] = make(map[int]int)
	}
	return id
}

func (p *AllegroParser) ensureNet(name string) int {
	if id, ok := p.netIndex[name]; ok {
		return id
	}
	id := len(p.netNames)
	p.netIndex[name] = id
	p.netNames = append(p.netNames, name)
	return id
}

func (p *AllegroParser) AddConn(elem, net string) {
	eid := p.ensureElem(elem)
	nid := p.ensureNet(net)
	p.Q[eid][nid]++
}

func (p *AllegroParser) Parse(r io.Reader) (*ParseResult, error) {
	sc := bufio.NewScanner(r)

	inNets := false
	var curNet string

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "$NETS") {
			inNets = true
			continue
		}

		if strings.HasPrefix(line, "$END") {
			break
		}

		if !inNets {
			continue
		}

		if strings.HasSuffix(line, ",") {
			line = strings.TrimSuffix(line, ",")
		}

		parts := strings.SplitN(line, ";", 2)
		if len(parts) < 2 {
			continue
		}

		netName := strings.TrimSpace(parts[0])
		rest := strings.TrimSpace(parts[1])

		curNet = netName
		p.processAllegroLine(rest, curNet)
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	R := ComputeR(p.Q)
	Plot := ComputePlot(p.elemNames)
	D := ComputeD(p.elemNames, Plot)

	FBefore := computeTotalFunctional(R, D)

	netElements := BuildNetElements(p.Q, Plot)
	plotCopy := make([]Element, len(Plot))
	copy(plotCopy, Plot)

	optPlot := Optimize(R, plotCopy, 5)

	DOpt := ComputeD(p.elemNames, optPlot)
	FAfter := computeTotalFunctional(R, DOpt)

	return &ParseResult{
		ElemNames:   p.elemNames,
		NetNames:    p.netNames,
		Q:           p.Q,
		R:           R,
		Plot:        Plot,
		OptPlot:     optPlot,
		D:           D,
		NetElements: netElements,
		FBefore:     FBefore,
		FAfter:      FAfter,
	}, nil
}

func (p *AllegroParser) processAllegroLine(line string, net string) {
	parts := strings.Fields(line)
	for _, t := range parts {
		t = strings.TrimSpace(t)
		if t == "" {
			continue
		}

		name := strings.Split(t, ".")[0]
		if name != "" {
			p.AddConn(name, net)
		}
	}
}

package parser

import (
	"bufio"
	"io"
	"strings"
)

type CalayParser struct {
	elemNames []string            // Индекс → имя
	netNames  []string            // Индекс → имя
	elemIndex map[string]int      // Имя → индекс
	netIndex  map[string]int      // Имя → индекс
	Q         map[int]map[int]int // [elemId][netId] = количество связей
}

func NewCalayParser() *CalayParser {
	return &CalayParser{
		elemNames: []string{},
		netNames:  []string{},
		elemIndex: make(map[string]int),
		netIndex:  make(map[string]int),
		Q:         make(map[int]map[int]int),
	}
}

func (p *CalayParser) Name() string { return "Calay" }

func (p *CalayParser) ensureElem(name string) int {
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

func (p *CalayParser) ensureNet(name string) int {
	if id, ok := p.netIndex[name]; ok {
		return id
	}
	id := len(p.netNames)
	p.netIndex[name] = id
	p.netNames = append(p.netNames, name)
	return id
}

func (p *CalayParser) AddConn(elem, net string) {
	eid := p.ensureElem(elem)
	nid := p.ensureNet(net)
	p.Q[eid][nid]++
}

func (p *CalayParser) Parse(r io.Reader) (*ParseResult, error) {
	sc := bufio.NewScanner(r)
	var curNet string

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		m := calayHeaderRe.FindStringSubmatch(line)
		if len(m) >= 2 {
			curNet = m[1]
			rest := strings.TrimSpace(line[len(m[0]):])
			p.processCalayLine(rest, curNet)
			continue
		}

		if curNet != "" {
			p.processCalayLine(line, curNet)
		}
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

func (p *CalayParser) processCalayLine(line, net string) {
	for _, m := range calayRe.FindAllStringSubmatch(line, -1) {
		if len(m) >= 3 {
			p.AddConn(m[1], net)
		}
	}
}

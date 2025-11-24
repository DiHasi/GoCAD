package parser

import (
	"math"
)

func ComputeR(Q map[int]map[int]int) map[int]map[int]int {
	R := make(map[int]map[int]int)

	for e1, nets1 := range Q {
		if _, ok := R[e1]; !ok {
			R[e1] = make(map[int]int)
		}
		for e2, nets2 := range Q {
			if e1 == e2 {
				continue
			}
			if _, ok := R[e2]; !ok {
				R[e2] = make(map[int]int)
			}
			count := 0
			for n := range nets1 {
				if nets2[n] != 0 {
					count++
				}
			}
			R[e1][e2] = count
		}
	}

	return R
}

func ComputePlot(elemNames []string) []Element {
	nElems := len(elemNames)
	if nElems == 0 {
		return nil
	}

	rows := int(math.Floor(math.Sqrt(float64(nElems))))
	if rows == 0 {
		rows = 1
	}
	cols := int(math.Ceil(float64(nElems) / float64(rows)))

	plot := make([]Element, 0, nElems)
	idx := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if idx >= nElems {
				break
			}
			elem := Element{
				Name: idx,
				X:    c, // X — колонка
				Y:    r, // Y — строка
			}
			plot = append(plot, elem)
			idx++
		}
	}
	return plot
}

func ComputeD(elemNames []string, plot []Element) [][]float64 {
	nElems := len(elemNames)
	if nElems == 0 {
		return nil
	}

	D := make([][]float64, nElems)
	for i := range D {
		D[i] = make([]float64, nElems)
	}

	for i := 0; i < nElems; i++ {
		for j := 0; j < nElems; j++ {
			if i == j {
				D[i][j] = 0
				continue
			}
			dx := float64(plot[i].X - plot[j].X)
			dy := float64(plot[i].Y - plot[j].Y)
			D[i][j] = math.Sqrt(dx*dx + dy*dy)
		}
	}

	return D
}

func BuildNetElements(Q map[int]map[int]int, plot []Element) map[int][]Element {
	X := make(map[int][]Element)

	elemPos := make(map[int]Element)
	for _, element := range plot {
		if element.Name >= 0 {
			elemPos[element.Name] = element
		}
	}

	for elemID, nets := range Q {
		for netID, count := range nets {
			if count > 0 {
				if pos, ok := elemPos[elemID]; ok {
					X[netID] = append(X[netID], pos)
				}
			}
		}
	}

	return X
}

func computeTotalFunctional(R map[int]map[int]int, D [][]float64) float64 {
	sum := 0.0
	for i, row := range R {
		for j, val := range row {
			sum += float64(val) * D[i][j]
		}
	}
	return sum
}

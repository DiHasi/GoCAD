package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

func (pr *ParseResult) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(pr)
}

func (pr *ParseResult) PrintQ() {
	netSet := map[int]struct{}{}
	for _, nets := range pr.Q {
		for n := range nets {
			netSet[n] = struct{}{}
		}
	}
	nets := make([]int, 0, len(netSet))
	for n := range netSet {
		nets = append(nets, n)
	}
	sort.Ints(nets)

	elems := make([]int, 0, len(pr.ElemNames))
	for i := range pr.ElemNames {
		elems = append(elems, i)
	}
	sort.Ints(elems)

	fmt.Printf("%15s", " ")
	for _, n := range nets {
		fmt.Printf("%10s", pr.NetNames[n])
	}
	fmt.Println()

	for _, e := range elems {
		fmt.Printf("%15s", pr.ElemNames[e])
		for _, n := range nets {
			fmt.Printf("%10d", pr.Q[e][n])
		}
		fmt.Println()
	}
}

func (pr *ParseResult) PrintR() {
	elems := make([]int, 0, len(pr.ElemNames))
	for i := range pr.ElemNames {
		elems = append(elems, i)
	}
	sort.Ints(elems)

	fmt.Printf("%15s", " ")
	for _, e := range elems {
		fmt.Printf("%10s", pr.ElemNames[e])
	}
	fmt.Println()

	for _, e1 := range elems {
		fmt.Printf("%15s", pr.ElemNames[e1])
		for _, e2 := range elems {
			val := 0
			if row, ok := pr.R[e1]; ok {
				if v, ok := row[e2]; ok {
					val = v
				}
			}
			fmt.Printf("%10d", val)
		}
		fmt.Println()
	}
}

func (pr *ParseResult) PrintPlot() {
	if pr.Plot == nil || len(pr.Plot) == 0 {
		fmt.Println("Plot is empty")
		return
	}

	for _, v := range pr.Plot {

		if v.Name >= 0 && v.Name < len(pr.ElemNames) {
			fmt.Printf("%10s %10d %10d", pr.ElemNames[v.Name], v.X, v.Y)
		} else {
			fmt.Printf("%10s", "")
		}

		fmt.Println()
	}
}

func (pr *ParseResult) PrintD() {
	if pr.D == nil || len(pr.D) == 0 {
		fmt.Println("Matrix D is empty")
		return
	}

	fmt.Printf("%15s", " ")
	for _, name := range pr.ElemNames {
		fmt.Printf("%10s", name)
	}
	fmt.Println()

	for i, row := range pr.D {
		fmt.Printf("%15s", pr.ElemNames[i])
		for _, val := range row {
			fmt.Printf("%10.2f", val)
		}
		fmt.Println()
	}
}

//func (pr *ParseResult) PrintX() {
//	// Собираем все сети
//	netSet := map[int]struct{}{}
//	for n := range pr.X {
//		netSet[n] = struct{}{}
//	}
//
//	nets := make([]int, 0, len(netSet))
//	for n := range netSet {
//		nets = append(nets, n)
//	}
//	sort.Ints(nets)
//
//	// Собираем все элементы
//	elems := make([]int, 0, len(pr.ElemNames))
//	for i := range pr.ElemNames {
//		elems = append(elems, i)
//	}
//	sort.Ints(elems)
//
//	// Заголовок: названия сетей
//	fmt.Printf("%15s", " ")
//	for _, n := range elems {
//		fmt.Printf("%10s", pr.ElemNames[n])
//	}
//	fmt.Println()
//
//	for _, n := range nets {
//		fmt.Printf("%15s", pr.NetNames[n])
//		for _, e := range elems {
//			found := 0
//			if elemsInNet, ok := pr.X[n]; ok {
//				for _, el := range elemsInNet {
//					if el == e {
//						found = 1
//						break
//					}
//				}
//			}
//			fmt.Printf("%10d", found)
//		}
//		fmt.Println()
//	}
//}

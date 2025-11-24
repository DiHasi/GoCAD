package parser

import (
	"math"
)

// Optimize — метод парных перестановок
// Для каждого элемента i подбирает наилучший обмен j, улучшающий функционал.
func Optimize(R map[int]map[int]int, plot []Element, maxIterations int) []Element {
	n := len(plot)
	if n == 0 {
		return plot
	}

	for iter := 0; iter < maxIterations; iter++ {
		D := computeDistanceMatrix(plot)
		W := computeW(R, D)
		//fmt.Println("W: ", W)
		improved := false

		// --- Для каждого элемента i ищем лучший обмен ---
		for i := 0; i < n; i++ {
			bestDelta := 0.0
			bestJ := -1

			for j := 0; j < n; j++ {
				if i == j {
					continue
				}

				// Пробуем обмен i <-> j
				swapElements(&plot[i], &plot[j])

				newD := computeDistanceMatrix(plot)
				newW := computeW(R, newD)

				delta := (W[i] + W[j]) - (newW[i] + newW[j])

				// Откатываем
				swapElements(&plot[i], &plot[j])

				if delta > bestDelta {
					bestDelta = delta
					bestJ = j
				}
			}

			// Если для i найдено улучшение — применяем
			if bestDelta > 0 && bestJ >= 0 {
				swapElements(&plot[i], &plot[bestJ])
				improved = true
			}
		}

		// Если за итерацию ни одно улучшение не найдено — завершаем
		if !improved {
			break
		}
	}

	return plot
}

// computeDistanceMatrix — матрица расстояний D[i][j]
func computeDistanceMatrix(plot []Element) [][]float64 {
	n := len(plot)
	D := make([][]float64, n)
	for i := range D {
		D[i] = make([]float64, n)
		for j := range D[i] {
			if i == j {
				continue
			}
			dx := float64(plot[i].X - plot[j].X)
			dy := float64(plot[i].Y - plot[j].Y)
			D[i][j] = math.Sqrt(dx*dx + dy*dy)
		}
	}
	return D
}

// computeW — W(i) = Σ R[i][k] * D[i][k]
func computeW(R map[int]map[int]int, D [][]float64) []float64 {
	n := len(D)
	W := make([]float64, n)

	for i := 0; i < n; i++ {
		row, ok := R[i]
		if !ok {
			continue
		}
		for k := 0; k < n; k++ {
			val := row[k]
			W[i] += float64(val) * D[i][k]
		}
	}
	return W
}

// swapElements — меняет координаты двух элементов
func swapElements(a, b *Element) {
	a.X, b.X = b.X, a.X
	a.Y, b.Y = b.Y, a.Y
}

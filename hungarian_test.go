package hungarian_test

import (
	"github.com/arthurkushman/go-hungarian"
	"testing"
)

var testsMax = []struct {
	m      [][]float64
	result map[int]map[int]float64
}{
	{[][]float64{
		{6, 2, 3, 4, 5},
		{3, 8, 2, 8, 1},
		{9, 9, 5, 4, 2},
		{6, 7, 3, 4, 3},
		{1, 2, 6, 4, 9},
	}, map[int]map[int]float64{
		0: {2: 3},
		1: {3: 8},
		2: {0: 9},
		3: {1: 7},
		4: {4: 9},
	}},
}

func TestSolveMax(t *testing.T) {
	for _, value := range testsMax {
		for key, val := range hungarian.SolveMax(value.m) {
			for k, v := range val {
				if v != value.result[key][k] {
					t.Fatalf("Want %f, got: %f", v, value.result[key][k])
				}
			}
		}
	}
}

var testsMin = []struct {
	m [][]float64
}{
	{[][]float64{
		{6, 2, 3, 4, 5, 11, 3, 8},
		{3, 8, 2, 8, 1, 12, 5, 4},
		{7, 9, 5, 10, 2, 11, 6, 8},
		{6, 7, 3, 4, 3, 5, 5, 3},
		{1, 2, 6, 13, 9, 11, 3, 6},
		{6, 2, 3, 4, 5, 11, 3, 8},
		{4, 6, 8, 9, 7, 1, 5, 3},
		{9, 1, 2, 5, 2, 7, 3, 8},
	}},
}

func TestSolveMin(t *testing.T) {
	data := make(map[int]float64)
	for _, value := range testsMin {
		for _, val := range hungarian.SolveMin(value.m) {
			for k, v := range val {
				if val, ok := data[k]; ok {
					t.Fatalf("Repeated column %d: %f", k, val)
				}
				data[k] = v
			}
		}
	}
}

var benchmarks = []struct {
	m [][]float64
}{
	{[][]float64{
		{6, 2, 3, 4, 5, 11, 3, 8},
		{3, 8, 2, 8, 1, 12, 5, 4},
		{7, 9, 5, 10, 2, 11, 6, 8},
		{6, 7, 3, 4, 3, 5, 5, 3},
		{1, 2, 6, 13, 9, 11, 3, 6},
		{6, 2, 3, 4, 5, 11, 3, 8},
		{4, 6, 8, 9, 7, 1, 5, 3},
		{9, 1, 2, 5, 2, 7, 3, 8},
	}},
	{[][]float64{
		{6, 2, 3, 4, 5, 11, 3, 8, 15, 18},
		{3, 8, 2, 12, 33, 8, 1, 12, 5, 4},
		{7, 9, 5, 11, 10, 2, 22, 11, 6, 8},
		{6, 7, 3, 4, 32, 3, 5, 5, 23, 3},
		{1, 2, 21, 6, 13, 9, 11, 3, 18, 6},
		{6, 2, 17, 3, 4, 41, 5, 11, 3, 8},
		{4, 6, 13, 8, 9, 7, 27, 1, 5, 3},
		{9, 1, 2, 16, 5, 2, 7, 31, 3, 8},
		{7, 1, 13, 8, 9, 4, 27, 6, 5, 3},
		{9, 2, 6, 16, 5, 1, 7, 31, 3, 8},
	}},
	{[][]float64{
		{6, 2, 72, 3, 4, 5, 11, 3, 19, 8, 15, 18},
		{3, 8, 2, 18, 12, 33, 8, 1, 34, 12, 5, 4},
		{7, 9, 5, 11, 10, 51, 2, 22, 11, 6, 15, 8},
		{6, 7, 3, 4, 32, 3, 5, 9, 5, 16, 23, 3},
		{1, 12, 2, 21, 6, 13, 9, 11, 17, 3, 18, 6},
		{6, 2, 16, 37, 17, 3, 4, 41, 5, 11, 3, 8},
		{4, 15, 6, 13, 8, 9, 7, 19, 27, 1, 5, 3},
		{9, 1, 73, 39, 2, 16, 5, 2, 7, 31, 3, 8},
		{6, 2, 72, 3, 4, 5, 11, 3, 19, 8, 15, 18},
		{3, 8, 2, 18, 12, 33, 8, 1, 34, 12, 5, 4},
		{7, 9, 5, 11, 10, 51, 2, 22, 11, 6, 15, 8},
		{6, 7, 3, 4, 32, 3, 5, 9, 5, 16, 23, 3},
	}},
}

func BenchmarkSolveMax(b *testing.B) {
	for _, v := range benchmarks {
		for i := 0; i < b.N; i++ {
			hungarian.SolveMax(v.m)
		}
	}
}

func BenchmarkSolveMin(b *testing.B) {
	for _, v := range benchmarks {
		for i := 0; i < b.N; i++ {
			hungarian.SolveMin(v.m)
		}
	}
}

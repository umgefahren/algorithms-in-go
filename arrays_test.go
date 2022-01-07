package algorithms_in_go

import (
	"math/rand"
	"sort"
	"testing"
)

func TestQuicksort(t *testing.T) {
	arr := make([]int, 100_000)
	for idx, _ := range arr {
		arr[idx] = rand.Int()
	}
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	t.Run("SingleC", func(t *testing.T) {
		singleArr := make([]int, len(arr))
		copy(singleArr, arr)
		quickSort(singleArr, 0, len(arr)-1)
		for idx, a := range singleArr {
			b := sorted[idx]
			if a != b {
				t.Errorf("The %v-idx elements (correct) %v (wrong) %v are not equal", idx, b, a)
			}
		}
	})

	t.Run("MultiC", func(t *testing.T) {
		multiArr := make([]int, len(arr))
		copy(multiArr, arr)
		parQuickSort(multiArr, 0, len(arr)-1)
		for idx, a := range multiArr {
			b := sorted[idx]
			if a != b {
				t.Errorf("The %v-idx elements (correct) %v (wrong) %v are not equal", idx, b, a)
			}
		}
	})

}

func TestInsertionSort(t *testing.T) {
	arr := make([]int, 100)
	for idx, _ := range arr {
		arr[idx] = rand.Int()
	}
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)
	singleArr := make([]int, len(arr))
	copy(singleArr, arr)
	insertionSort(singleArr, 0, len(arr)-1)
	for idx, a := range singleArr {
		b := sorted[idx]
		if a != b {
			t.Errorf("The %v-idx elements (correct) %v (wrong) %v are not equal", idx, b, a)
		}
	}
}

func TestIQuicksort(t *testing.T) {
	arr := make([]int, 100_000)
	for idx, _ := range arr {
		arr[idx] = rand.Int()
	}
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)
	singleArr := make([]int, len(arr))
	copy(singleArr, arr)
	iQuickSort(singleArr, 0, len(arr)-1)
	for idx, a := range singleArr {
		b := sorted[idx]
		if a != b {
			t.Errorf("The %v-idx elements (correct) %v (wrong) %v are not equal", idx, b, a)
		}
	}
}

func BenchmarkQuicksortSingle(b *testing.B) {
	testArr := make([]int, 1_000_000)

	for idx, _ := range testArr {
		testArr[idx] = rand.Int()
	}

	// b.ResetTimer()

	for i := 0; i < b.N; i++ {
		singleArr := make([]int, 1_000_000)
		copy(singleArr, testArr)
		quickSort(singleArr, 0, len(testArr)-1)
	}
}

func BenchmarkIQuicksortSingle(b *testing.B) {
	testArr := make([]int, 1_000_000)

	for idx, _ := range testArr {
		testArr[idx] = rand.Int()
	}

	// b.ResetTimer()

	for i := 0; i < b.N; i++ {
		singleArr := make([]int, 1_000_000)
		copy(singleArr, testArr)
		iQuickSort(singleArr, 0, len(testArr)-1)
	}
}

func BenchmarkQuicksortMulti(b *testing.B) {
	testArr := make([]int, 1_000_000)

	for idx, _ := range testArr {
		testArr[idx] = rand.Int()
	}

	for i := 0; i < b.N; i++ {
		multiArr := make([]int, 1_000_000)
		copy(multiArr, testArr)
		parQuickSort(multiArr, 0, len(multiArr)-1)
	}

}

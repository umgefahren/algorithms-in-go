package algorithms_in_go

import "sync"

const multiThreshold = 5_000
const singleThreshold = 10

func insertionSort(a []int, l, r int) {
	for i := r; i >= l; i-- {
		for j := i; j < r; j++ {
			if a[j] > a[j+1] {
				swap(a, j, j+1)
			} else {
				break
			}
		}
	}
}

func quickSort(a []int, l, r int) {
	if l < r {
		k := partition(a, l, r)
		quickSort(a, l, k-1)
		quickSort(a, k+1, r)
	}
}

func iQuickSort(a []int, l, r int) {
	if l < r {
		k := partition(a, l, r)
		if k-1-l >= singleThreshold {
			iQuickSort(a, l, k-1)
		} else {
			insertionSort(a, l, k-1)
		}
		if r-(k+1) >= singleThreshold {
			iQuickSort(a, k+1, r)
		} else {
			insertionSort(a, k+1, r)
		}
	}
}

func parQuickSort(a []int, l, r int) {
	if l < r {
		k := partition(a, l, r)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			if k-1-l >= multiThreshold {
				parQuickSort(a, l, k-1)
			} else {
				iQuickSort(a, l, k-1)
			}

		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			if r-(k+1) >= multiThreshold {
				parQuickSort(a, k+1, r)
			} else {
				iQuickSort(a, k+1, r)
			}

		}()

		wg.Wait()
	}
}

func partition(a []int, l, r int) int {
	i := l
	j := r - 1
	p := a[r]
	for {
		for i < r && a[i] < p {
			i = i + 1
		}
		for j > l && a[j] > p {
			j = j - 1
		}
		if i < j {
			swap(a, i, j)
		}
		if i >= j {
			break
		}
	}
	swap(a, i, r)
	return i
}

func swap(a []int, i, j int) {
	b := a[i]
	a[i] = a[j]
	a[j] = b

}

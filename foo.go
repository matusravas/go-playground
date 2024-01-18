package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	result := []int{}
	mid := int16(len(arr) / 2)
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	lIndex, rIndex := 0, 0
	for lIndex < len(left) && rIndex < len(right) {
		if left[lIndex] < right[rIndex] {
			result = append(result, left[lIndex])
			lIndex += 1
		} else {
			result = append(result, right[rIndex])
			rIndex += 1
		}
	}
	result = append(result, left[lIndex:]...)
	result = append(result, right[rIndex:]...)
	return result
}

func bubbleSort(arr []int) []int {
	didSort := false
	for true {
		didSort = false
		for i, e := range arr {
			if i < len(arr)-1 && e > arr[i+1] {
				tmp := arr[i+1]
				arr[i+1] = e
				arr[i] = tmp
				didSort = true
			}
		}
		if !didSort {
			break
		}
	}
	return arr
}

func bubbleSortRecursive(arr []int) []int {
	n := len(arr)

	if n == 1 {
		return arr
	}
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			tmp := arr[i+1]
			arr[i+1] = arr[i]
			arr[i] = tmp
		}
	}
	var new []int
	return bubbleSortRecursive(append(new, arr[:n-1]...))
}

type Result struct {
	Index int
	found bool
}

func findNumber(arr []int, n int, agrs ...int) Result {
	var idx int
	if len(arr) == 1 && arr[0] != n {
		return Result{
			Index: -1,
			found: false,
		}
	}

	if len(agrs) == 1 {
		idx = agrs[0]
	}

	if arr[0] == n {
		return Result{
			Index: idx,
			found: true,
		}
	}
	return findNumber(arr[1:], n, idx+1)
}

func binarySearchRecursive(arr []int, n int) int {
	fmt.Println(arr)
	if len(arr) == 1 {
		return -1
	}

	idx := len(arr) / 2
	pivot := arr[idx]
	var left int
	right := len(arr)
	if n == pivot {
		return 1
	} else if n < pivot {
		right = idx
	} else {
		left = idx
	}
	fmt.Printf("left: %d right: %d\n", left, right)
	return binarySearchRecursive(arr[left:right], n)
}

func binarySearch(arr []int, n int) int {
	if len(arr) == 1 && arr[0] != n {
		return -1
	}
	// var left, right int
	// right = len(arr) - 1
	var left, right int

	for left <= right {

		fmt.Println(arr)

		mid := len(arr) / 2
		pivot := arr[mid]
		fmt.Println(pivot)
		if n == pivot {
			return 1
		} else if n < pivot {
			right = mid
			left = 0
		} else {
			left = mid + 1
			right = len(arr)
		}
		arr = arr[left:right]
	}

	return -1
}

// 65/2 32 1
// 32/2 16 0
// 16/2 8 0
// 8/2 4 0
// 4/2 2 0
// 2/2 1 0
func decToBin(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{1}
	}
	return append(decToBin(n/2), n%2)
}

// func mirror(arr []int, varagrs ...int) []int {
// 	var idx int
// 	if len(varagrs) == 1 {
// 		idx = varagrs[0]
// 	}

// 	n := len(arr)
// 	mid := len(arr) / 2

// 	tmp := arr[idx]
// 	arr[idx] = arr[n-idx-1]
// 	arr[n-idx-1] = tmp
// 	// fmt.Println(arr)
// 	if idx == mid {
// 		return arr
// 	}
// 	return mirror(arr, idx+1)
// }

func mirror(arr []int, varagrs ...int) []int {
	var left int
	right := len(arr) - 1
	if len(varagrs) == 2 {
		left = varagrs[0]
		right = varagrs[1]
	}
	if left > right {
		return arr
	}
	// tmp := arr[left]
	// arr[left] = arr[right]
	// arr[right] = tmp
	arr[left], arr[right] = arr[right], arr[left]
	return mirror(arr, left+1, right-1)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	pivot := arr[mid]
	var lt []int
	var gt []int
	var eq []int
	for _, e := range arr {
		if e < pivot {
			lt = append(lt, e)
		} else if e > pivot {
			gt = append(gt, e)
		} else {
			eq = append(eq, e)
		}
	}
	return append(quickSort(lt), append(eq, quickSort(gt)...)...)

}

func main() {
	x := []int{6, 5, 9, 8, 1, 7, 2, 4, 3, 2, 1}
	fmt.Println(x)
	sortedMerge := mergeSort(x)
	fmt.Println("Merge sort")
	fmt.Println(sortedMerge)
	// sortedBubble := bubbleSortRecursive(x)
	// fmt.Println("Bubble sort")
	// fmt.Println(sortedBubble)
	// sortedBubble := bubbleSort(x)
	// fmt.Println("Bubble sort")
	// fmt.Println(sortedBubble)
	quickSorted := quickSort(x)
	fmt.Println("Quick sort")
	fmt.Println(quickSorted)

	// fmt.Println("Finding...")
	// result := findNumber(x, 12)
	// fmt.Println(result)

	fmt.Println("Binary search...")
	result := binarySearchRecursive(quickSorted, 8)
	fmt.Println(result)

	// fmt.Println("Mirror...")
	// mirrored := mirror(x)
	// fmt.Println(mirrored)

	fmt.Printf("Decimal to Binary: %v\n", decToBin(65))
}

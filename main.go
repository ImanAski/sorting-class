package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("Please enter the number of elements -> ")
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	// fmt.Println("Please enter elements ->")
	for i := 0; i < n; i++ {
		arr[i] = rand.Int()
		// fmt.Scan(&arr[i])
	}

	bubblesortExecTime := BubbleSort(arr)
	fmt.Println("Array Sorted using BubbleSort in " + bubblesortExecTime.String())
	for i := 0; i < n; i++ {
		// fmt.Printf("%v, ", arr[i])
		if i == n-1 {
			fmt.Printf("\n")
		}
	}

	shellSortExecTime := BubbleSort(arr)
	fmt.Println("Array Sorted using ShellSort in " + shellSortExecTime.String())
	for i := 0; i < n; i++ {
		// fmt.Printf("%v, ", arr[i])
		if i == n-1 {
			fmt.Printf("\n")
		}
	}
	gnomeSortExecTime := BubbleSort(arr)
	fmt.Println("Array Sorted using GnomeSort in " + gnomeSortExecTime.String())
	for i := 0; i < n; i++ {
		// fmt.Printf("%v, ", arr[i])
		if i == n-1 {
			fmt.Printf("\n")
		}
	}
}

func BubbleSort(arr []int) time.Duration {
	startTime := time.Now()
	n := len(arr)
	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	endTime := time.Now()
	return endTime.Sub(startTime)
}

func ShellShort(arr []int) time.Duration {
	startTime := time.Now()
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			var j int

			for j = i; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
			}

			arr[j] = temp
		}
	}
	endTime := time.Now()
	return endTime.Sub(startTime)
}

func GnomeSort(arr []int) time.Duration {
	startTime := time.Now()
	n := len(arr)
	i := 0
	for i < n-1 {
		if i == 0 {
			i += 1
		}

		if arr[i] > arr[i-1] {
			i += 1
		}

		if arr[i] < arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i -= 1
		}
	}
	endTime := time.Now()
	return endTime.Sub(startTime)
}

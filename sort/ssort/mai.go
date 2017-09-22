package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
}

func getRandomIntSlice(length int) []int {
	s := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		s = append(s, rand.Intn(100))
	}

	return s
}

func testSort() {
	a := getRandomIntSlice(10)
	fmt.Println("unsorted : ", a)
	sSort(a)
	fmt.Println("sorted : ", a)
}

func main() {
	for i := 0; i < 10; i++ {
		testSort()
		fmt.Println("------------------我是漂亮的分隔线---------------------------")
		time.Sleep(time.Microsecond)
	}
}

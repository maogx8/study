package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
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
	bSort(a)
	fmt.Println("sorted : ", a)
}

func main() {
	for i := 0; i < 10; i++ {
		testSort()
		fmt.Println("------------------我是漂亮的分隔线---------------------------")
		time.Sleep(time.Microsecond)
	}
}

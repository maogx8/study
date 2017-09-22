package main

import (
	"fmt"
	"math/rand"
	"time"
)

func qSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	base := left
	leftPointer, rightPointer := left, right
	for leftPointer != rightPointer {
		for ; rightPointer > leftPointer; rightPointer-- {
			if arr[rightPointer] < arr[base] {
				break
			}
		}

		for ; leftPointer < rightPointer; leftPointer++ {
			if arr[leftPointer] > arr[base] {
				break
			}
		}

		if leftPointer < rightPointer {
			arr[leftPointer], arr[rightPointer] = arr[rightPointer], arr[leftPointer]
		}

		//fmt.Println(leftPointer, rightPointer, arr)
	}
	arr[base], arr[leftPointer] = arr[leftPointer], arr[base]
	qSort(arr, left, leftPointer-1)
	qSort(arr, leftPointer+1, right)
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
	//a := []int{54, 1, 47, 63, 12, 98, 74, 33, 66, 3}
	fmt.Println("unsorted : ", a)
	qSort(a, 0, 9)
	fmt.Println("sorted : ", a)
}

func main() {
	for i := 0; i < 10; i++ {
		testSort()
		fmt.Println("------------------我是漂亮的分隔线---------------------------")
		time.Sleep(time.Microsecond)
	}
}

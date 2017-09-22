package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insert(arr []int, a int) []int {
	length := len(arr)
	tmp := make([]int, length+1)
	postion := 0

	// find the postion to insert
	for ; postion < length; postion++ {
		if a < arr[postion] {
			break
		}
		tmp[postion] = arr[postion]
	}

	// right shift the value behind postion one by one
	for j := length; j > postion; j-- {
		tmp[j] = arr[j-1]
	}

	// set the insert postion value
	tmp[postion] = a
	return tmp
}

func iSort(arr []int) []int {
	s := make([]int, 0)
	length := len(arr)
	for i := 0; i < length; i++ {
		s = insert(s, arr[i])
	}

	return s
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
	a = iSort(a)
	fmt.Println("sorted : ", a)
}

func main() {
	for i := 0; i < 10; i++ {
		testSort()
		fmt.Println("------------------我是漂亮的分隔线---------------------------")
		time.Sleep(time.Microsecond)
	}
}

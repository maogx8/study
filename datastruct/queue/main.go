/*
go语言实现的队列，使用指定长度的切片实现，
切片最后一个位置不存储数据，用来判读队列是否已满。
head指向第一个元素，tail指向最后一个元素的下一个位置。
*/

package main

import "errors"
import "fmt"

type queue struct {
	data   []interface{}
	length int
	head   int
	tail   int
}

// Newqueue 创建队列
func Newqueue(length int) (*queue, error) {
	if length <= 0 {
		return nil, errors.New("queue length need greater than 0")
	}

	return &queue{
		data:   make([]interface{}, length+1),
		length: length + 1,
		head:   0,
		tail:   0,
	}, nil
}

func (s *queue) isEmpty() bool {
	if s.head == s.tail {
		return true
	}

	return false
}

func (s *queue) isFull() bool {
	if (s.tail+1)%s.length == s.head {
		return true
	}

	return false
}

func (s *queue) get() (interface{}, error) {

	if !s.isEmpty() {
		data := s.data[s.head]
		s.head = (s.head + 1) % s.length
		return data, nil
	}

	return nil, errors.New("empty queue")
}

func (s *queue) put(data interface{}) error {

	if !s.isFull() {
		s.data[s.tail] = data
		s.tail = (s.tail + 1) % s.length
		return nil
	}

	return errors.New("full queue")
}

func test(s *queue) {
	testIntData := []int{1, 2, 3, 4, 5, 6, 7}
	for _, data := range testIntData {
		err := s.put(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("put data: ", data)
	}

	for _ = range testIntData {
		data, err := s.get()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("get data: ", data)
	}

	testStringData := []string{"a", "b", "c", "d", "e", "f", "g"}
	for _, data := range testStringData {
		err := s.put(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("put data: ", data)
	}

	for _ = range testStringData {
		data, err := s.get()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("get data: ", data)
	}

}

func main() {
	s, err := Newqueue(5)
	if err != nil {
		fmt.Println("create queue fail")
	}
	//test(s)

	s.put(1)
	s.put("a")
	fmt.Println(s.get())
	fmt.Println(s.get())
	s.put(2)
	fmt.Println(s.get())
	s.put("b")
	fmt.Println(s.get())
}

package main

import "errors"
import "fmt"

type stack struct {
	data   []interface{}
	length int
	top    int
}

// Newstack 创建栈
func Newstack(length int) (*stack, error) {
	if length <= 0 {
		return nil, errors.New("stack length need greater than 0")
	}

	return &stack{
		data:   make([]interface{}, length),
		length: length,
		top:    -1,
	}, nil
}

func (s *stack) get() (interface{}, error) {
	if s.top >= 0 {
		data := s.data[s.top]
		s.top--
		return data, nil
	}
	return nil, errors.New("empty stack")
}

func (s *stack) put(data interface{}) error {
	if s.top < s.length-1 {
		s.top++
		s.data[s.top] = data
		return nil
	}

	return errors.New("full stack")
}

func main() {
	s, err := Newstack(5)
	if err != nil {
		fmt.Println("create stack fail")
	}
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

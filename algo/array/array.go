package array

import "errors"

type Array struct {
	data []int
}

func NewArray(len int) *Array {
	return &Array{
		data: make([]int, len),
	}
}

func (a *Array) Len() int {
	return len(a.data)
}

func (a *Array) InRange(i int) bool {
	return 0 <= i && i < len(a.data)
}

func (a *Array) Get(i int) (int, error) {
	if a.InRange(i) {
		return a.data[i], nil
	} else {
		return 0, errors.New("out of index range")
	}
}

func (a *Array) Set(i int, v int) error {
	if a.InRange(i) {
		a.data[i] = v
		return nil
	} else {
		return errors.New("out of index range")
	}
}

func (a *Array) Append(v int) error {
	a.data = append(a.data, v)
	return nil
}
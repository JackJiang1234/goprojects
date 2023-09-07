package array

import "testing"

func TestNewArr(t *testing.T) {
	a := NewArray(1)
	if a == nil {
		t.Errorf("expect not nil, actual %v", a)
	}
}

func TestGet(t *testing.T){
	a := NewArray(1)
	a.Set(0, 2)
	v, err := a.Get(0)
	if (err != nil){
		t.Errorf("expect not error")
	}
	if v != 2 {
		t.Errorf("expect value as 2, actual is %d", v)
	}
}
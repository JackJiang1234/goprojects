package singleton

import "testing"

func TestSingletonInstance(t *testing.T) {
	var obj1 = GetInstance()
	var obj2 = GetInstance()
	if obj1 != obj2 {
		t.Error("singleton test failed!")
	}
}

func TestSingletonAdd(t *testing.T){
	var obj1 = GetInstance()
	obj1.AddOne()

	var obj2 = GetInstance()
	result := obj2.AddOne()

	if result != 2 {
		t.Errorf("expected:%d, got: %d", 2, result)
	}
}
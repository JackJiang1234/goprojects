package factory

import "testing"

func TestCountryCreate(t *testing.T) {
	var c = NewCountry("china")
	if c != nil {
		c.Greet()
		t.Log("test passed")
	} else {
		t.Error("fail")
	}
}

package linklist

import "testing"

func TestNewSingleLinklist(t *testing.T) {
	l := NewLinkList()
	if l == nil {
		t.Error("new link list should not be nil")
	}
}

func TestAppend(t *testing.T){
	l := NewLinkList()
	l.Append(1)
	l.Append(2)

	if l.Count() != 2 {
		t.Error("append two element, lenght should be 2")
	}
}
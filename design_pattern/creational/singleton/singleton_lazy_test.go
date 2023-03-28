package singleton

import "testing"

func TestGetLazyInstance(t *testing.T){
	if GetLazyInstance() != GetLazyInstance(){
		t.Error("lazy singleton test fail")
	}
}
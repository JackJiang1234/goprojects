package main

import "github.com/JackJiang1234/goprojects/instrument_trace"

func foo() {
	defer trace.Trace()()
	bar()
}

func bar() {
	defer trace.Trace()()

}

func main() {
	defer trace.Trace()()
	foo()
}

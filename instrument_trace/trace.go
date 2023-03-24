package trace

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var goroutineSpace = []byte("goroutine")

func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	//fmt.Println("origin:", string(b))
	b = bytes.TrimPrefix(b, goroutineSpace)
	//fmt.Println("after trim:", string(b))
	i := bytes.LastIndexByte(b, ' ')
	//fmt.Println("index:", i)
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))	
	}
	b = b[:i]
	n, err := strconv.ParseUint(strings.TrimSpace(string(b)), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}

	return n
}

/*
func Trace(name string) func() {
	println("enter:", name)
	return func() {
		println("exit:", name)
	}
}
*/

/*
func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	gid := curGoroutineID()
	fmt.Printf("g[%05d]: enter: [%s]\n", gid, name)
	return func() { fmt.Printf("g[%05d]: exit: [%s]\n", gid, name) }
}

*/

func printTrace(id uint64, name, arrow string, indent int){
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "	"
	}
	fmt.Printf("g[%05d]:%s%s%s\n", id, indents, arrow, name)
}

var mu sync.Mutex
var m = make(map[uint64]int)

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	gid := curGoroutineID()

	mu.Lock()
	indents := m[gid]
	m[gid] = indents + 1
	mu.Unlock()
	printTrace(gid, name, "->", indents + 1)
	return func() {
		mu.Lock()
		indents := m[gid]
		m[gid] = indents - 1
		mu.Unlock()
		printTrace(gid, name, "<-", indents)
	}
}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
}

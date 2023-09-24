package main

import (
	"fmt"

	"github.com/jackjiang/crawler/collect"
)

func main() {
	f := collect.NewBrowserFetch()
	body, err := f.Get("https://book.douban.com/subject/1007305/")
	if err != nil {
		fmt.Printf("get url error:%v\n", err)
	}

	fmt.Println(string(body))
}

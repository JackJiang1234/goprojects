package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.thepaper.cn"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("fetch url error:%v", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%v", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read content failed:%v\n", err)
	}

	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has %d links\n", numLinks)

	exits := strings.Contains(string(body), "疫情")
	fmt.Printf("是否存疫情%v \n", exits)

	bcount := bytes.Count(body, []byte("<a"))
	fmt.Printf("homepage has %d links\n", bcount)
}

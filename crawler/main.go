package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func main() {
	url := "https://www.thepaper.cn"
	resp, err := Fetch(url)
	if err != nil {
		fmt.Printf("fetch url error: %v\n", err)
	}
	fmt.Printf("%s", string(resp))
}

func Fetch(url string)([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Read := transform.NewReader(bodyReader, e.NewDecoder())

	return io.ReadAll(utf8Read)
}

func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Printf("fetch error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
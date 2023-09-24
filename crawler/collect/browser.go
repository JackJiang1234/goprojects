package collect

import (
	"bufio"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/text/transform"
)

type BrowserFetch struct {
}

func (BrowserFetch) Get(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("get url failed:%v", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return io.ReadAll(utf8Reader)
}

func NewBrowserFetch() Fetcher {
	return &BrowserFetch{}
}
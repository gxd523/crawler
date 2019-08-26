package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(200 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	log.Printf("Fetching url=%s\n", url)
	<-rateLimiter
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add(
		"User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.2 Safari/605.1.15",
	)
	request.Header.Add(
		"timestamp",
		time.StampMilli,
	)
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code=%d,msg=%s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	reader := bufio.NewReader(resp.Body) // io.Reader读过之后指针一定会移动，bufio.Reader使用Peek()后，指针不会移动
	encode := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader, encode.NewDecoder()) // gopm get -g -v golang.org/x/text
	bytes, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	encode, _, _ := charset.DetermineEncoding(bytes, "") // gopm get -g -v golang.org/x/net/html
	return encode
}

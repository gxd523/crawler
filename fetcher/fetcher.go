package fetcher

import (
	"bufio"
	"demos/util"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add(
		"User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.2 Safari/605.1.15",
	)
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			util.PanicWrapper(err)
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
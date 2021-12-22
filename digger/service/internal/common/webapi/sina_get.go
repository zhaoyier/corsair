package webapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	// log "github.com/Sirupsen/logrus"
)

// data[3] current p
func GetSinaDayDetail(code string) ([]string, error) {
	url := fmt.Sprintf("http://hq.sinajs.cn/list=%s", strings.ToLower(code))
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, fmt.Errorf("create http request failed: %v", err)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do http request failed: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %v", err)
	}
	data, start, end := string(body), 0, 0
	if start = strings.IndexAny(data, `"`); start == -1 {
		return nil, fmt.Errorf("not found prefix: %s", code)
	}
	if end = strings.LastIndexAny(data, `"`); end == -1 {
		return nil, fmt.Errorf("not found suffix: %s", code)
	}

	data = data[start:end]
	return strings.Split(data, ","), nil
}

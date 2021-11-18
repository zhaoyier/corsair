package webapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	maxRetry = 3
)

var client = &http.Client{
	Timeout: time.Minute * 3,
}

func RetryPost(url string, req, resp interface{}) error {
	retryTime := 0
retry:
	err := postRequest(url, req, resp)
	if err != nil {
		if retryTime >= maxRetry {
			return err
		}
		retryTime++
		goto retry
	}

	return nil
}

func postRequest(url string, req, resp interface{}) error {
	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal request \"%s\" failed: %v", string(data), err)
	}

	reader := bytes.NewReader(data)
	httpReq, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return fmt.Errorf("create http request failed: %v", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("do http request failed: %v", err)
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("webapi return bad code: %d", httpResp.StatusCode)
	}

	respData, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return fmt.Errorf("read response body failed: %v", err)
	}

	err = json.Unmarshal(respData, resp)
	if err != nil {
		return fmt.Errorf("unmarshal response failed: %v", err)
	}

	return nil
}

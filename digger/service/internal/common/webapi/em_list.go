package webapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetEastmoneyCode(page, size int32, resp interface{}) error {
	var retryTime int32
retry:
	err := eastmoneyCode(page, size, resp)
	if err != nil {
		if retryTime >= maxRetry {
			return err
		}
		retryTime++
		goto retry
	}

	return nil

}

func eastmoneyCode(page, size int32, resp interface{}) error {
	url := fmt.Sprintf("http://26.push2.eastmoney.com/api/qt/clist/get?cb=jQuery11240682916251377502_1637505893462&pn=%d&pz=%d&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048&fields=f2,f3,f4,f5,f6,f10,f12,f15,f17,f20,f21,f23", page, size)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return fmt.Errorf("create http request failed: %v", err)
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Referer", "http://quote.eastmoney.com/")
	req.Header.Add("Accept-Language", "zh,zh-CN;q=0.9")
	req.Header.Add("Cookie", "em_hq_fls=js; qgqp_b_id=852f40f543bbed655b5642a128a6b0d8; st_si=35481744959379; st_asi=delete; HAList=a-sz-300059-%u4E1C%u65B9%u8D22%u5BCC%2Ca-sz-002202-%u91D1%u98CE%u79D1%u6280%2Ca-sh-601808-%u4E2D%u6D77%u6CB9%u670D%2Ca-sz-000792-%u76D0%u6E56%u80A1%u4EFD; cowCookie=true; intellpositionL=1012px; intellpositionT=455px; st_pvi=29102764427919; st_sp=2021-08-10%2010%3A30%3A12; st_inirUrl=http%3A%2F%2F3.3.3.3%2F; st_sn=59; st_psi=20211121225449851-113200301321-1359493152")

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("do http request failed: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("read response body failed: %v", err)
	}

	body = filterPrefix(body)
	err = json.Unmarshal(body, resp)
	if err != nil {
		return fmt.Errorf("unmarshal response failed: %v", err)
	}

	return nil
}

func filterPrefix(body []byte) []byte {
	var start, end int
	for idx, val := range body {
		if val == '(' {
			start = idx
		}
		if val == ')' {
			end = idx
		}
	}

	body = body[start+1 : end]
	return body
}

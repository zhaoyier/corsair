package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
)

func GetTXDayDetail(code string) ([]string, error) {
	url := fmt.Sprintf("http://qt.gtimg.cn/q=%s", strings.ToLower(code))
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
	return strings.Split(data, "~"), nil
}

func GetTXDayPrice(secucode string) float64 {
	codes := strings.Split(secucode, ".")
	secucode = strings.Join(codes, "")
	secucode = strings.ToLower(secucode)
	ts := time.Now().Unix()
	result, ok := presentPriceMap.Load(secucode)
	if ok {
		data := result.(*trpc.PresentPrice)
		if data.Timestamp > ts {
			return data.Price
		} else {
			presentPriceMap.Delete(secucode)
		}
	}

	results, err := GetTXDayDetail(secucode)
	if err != nil {
		log.Errorf("get present price failed: %s|%q", secucode, err)
		return 0
	}

	if len(results) < 3 {
		log.Errorf("get present price invalid: %s|%+v", secucode, results)
		return 0
	}
	price := results[3]
	pp := utils.String2Float64(price)
	presentPriceMap.Store(secucode, &trpc.PresentPrice{
		Timestamp: ts + 15*60,
		Price:     pp,
	})
	return pp
}

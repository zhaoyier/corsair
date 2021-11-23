package webapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.ezbuy.me/ezbuy/corsair/digger/rpc/digger"
)

func GetEastmoneyData(typ digger.EastMoneyType, code string, resp interface{}) error {
	url, retryTime := "", 0
	switch typ {
	case digger.EastMoneyType_EastMoneyTypeHolder:
		url = fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/ShareholderResearch/PageAjax?code=%s", code)
	case digger.EastMoneyType_EastMoneyTypeNews:
		url = fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewsBulletin/PageAjax?code=%s", code)
	case digger.EastMoneyType_EastMoneyTypeOperations:
		url = fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/OperationsRequired/OperationsRequiredAjax?times=1&code=%s", code)
	default:
		return fmt.Errorf("invalid type: %+v", typ.Short())
	}

retry:
	err := eastmoneyGet(url, resp)
	if err != nil {
		if retryTime >= maxRetry {
			return err
		}
		retryTime++
		goto retry
	}

	return nil

}

func eastmoneyGet(url string, resp interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return fmt.Errorf("create http request failed: %v", err)
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("do http request failed: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("read response body failed: %v", err)
	}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return fmt.Errorf("unmarshal response failed: %v", err)
	}

	return nil
}

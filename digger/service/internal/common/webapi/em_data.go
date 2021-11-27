package webapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
	case digger.EastMoneyType_EastMoneyTypeGPList:
		pageNum, _ := strconv.Atoi(code)
		url = fmt.Sprintf("http://26.push2.eastmoney.com/api/qt/clist/get?cb=jQuery11240682916251377502_1637505893462&pn=%d&pz=40&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048&fields=f2,f3,f4,f5,f6,f10,f12,f15,f17,f20,f21,f23", pageNum)
	default:
		return fmt.Errorf("invalid type: %+v", typ.String())
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

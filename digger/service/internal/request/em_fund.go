package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func GetEastmoneyFundFlow(secucode int64, resp interface{}) error {
	var retryTime int32
retry:
	err := eastmoneyFoudFlow(secucode, resp)
	log.Infof("==>>TODO 871: %+v", err)
	if err != nil {
		log.Errorf("east get code failed: %s|%q", secucode, err)
		if retryTime >= maxRetry {
			return err
		}
		retryTime++
		goto retry
	}

	return nil

}

func eastmoneyFoudFlow(secucode int64, resp interface{}) error {
	var prefix int32 = 1
	if secucode < 600000 {
		prefix = 0
	}

	url := fmt.Sprintf("https://push2.eastmoney.com/api/qt/ulist.np/get?fltt=2&secids=%d.%d&fields=f164,f174,f252&ut=b2884a393a59ad64002292a3e90d46a5", prefix, secucode)
	log.Infof("==>>TODO 881: %+v", url)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("Sec-Fetch-Mode", "no-cors")
	req.Header.Add("Sec-Fetch-Dest", "script")
	req.Header.Add("Referer", "https://data.eastmoney.com/")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Cookie", "qgqp_b_id=c3e233a777a7a2ed2872326c57200a01; em_hq_fls=js; st_si=42383655869778; cowCookie=true; intellpositionL=1080px; cowminicookie=true; st_asi=delete; intellpositionT=455px; HAList=a-sh-600941-%u4E2D%u56FD%u79FB%u52A8%2Ca-sh-600519-%u8D35%u5DDE%u8305%u53F0; st_pvi=81631919063174; st_sp=2022-01-12%2022%3A40%3A09; st_inirUrl=https%3A%2F%2Fcn.bing.com%2F; st_sn=14; st_psi=20220112231314210-113300300815-0441393438")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	// body = filterPrefix(body)
	log.Infof("==>>TODO 886: %+v", string(body))
	err = json.Unmarshal(replacChar(body), resp)
	if err != nil {
		log.Infof("==>>TODO 887: %+v", err)
		return fmt.Errorf("unmarshal response failed: %v", err)
	}
	log.Infof("==>>TODO 888: %+v", resp)
	return nil
}

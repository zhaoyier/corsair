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

	// url := "https://push2his.eastmoney.com/api/qt/stock/fflow/daykline/get?lmt=0&klt=101&fields1=f1%252Cf2%252Cf3%252Cf7&fields2=f51%252Cf52%252Cf53%252Cf54%252Cf55%252Cf56%252Cf57%252Cf58%252Cf59%252Cf60%252Cf61%252Cf62%252Cf63%252Cf64%252Cf65&secid=0.300204&_=1642245260116"
	url := fmt.Sprintf("https://push2his.eastmoney.com/api/qt/stock/fflow/daykline/get?lmt=0&klt=101&fields1=f1,f2,f3,f7&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63,f64,f65&secid=%d.%d", prefix, secucode)

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
	req.Header.Add("Accept-Language", "zh,zh-CN;q=0.9")
	req.Header.Add("Cookie", "em_hq_fls=js; qgqp_b_id=852f40f543bbed655b5642a128a6b0d8; intellpositionL=1012px; em-quote-version=topspeed; emshistory=%5B%22%E9%95%87%E6%B4%8B%E5%8F%91%E5%B1%95%22%5D; st_si=88286280226808; cowminicookie=true; st_asi=delete; HAList=a-sz-300204-%u8212%u6CF0%u795E%2Ca-sz-301138-%u534E%u7814%u7CBE%u673A%2Ca-sh-600031-%u4E09%u4E00%u91CD%u5DE5%2Ca-sz-301055-%u5F20%u5C0F%u6CC9%2Ca-sz-002304-%u6D0B%u6CB3%u80A1%u4EFD%2Ca-sz-000564-*ST%u5927%u96C6%2Ca-sz-001218-%u4E3D%u81E3%u5B9E%u4E1A%2Ca-sz-301090-%u534E%u6DA6%u6750%u6599%2Ca-sz-002667-%u978D%u91CD%u80A1%u4EFD%2Ca-sz-000858-%u4E94%20%u7CAE%20%u6DB2%2Ca-sz-002897-%u610F%u534E%u80A1%u4EFD; cowCookie=true; st_pvi=54508668225105; st_sp=2022-01-02%2011%3A40%3A23; st_inirUrl=https%3A%2F%2Fwww.baidu.com%2Flink; st_sn=317; st_psi=20220115191227735-113300300815-1252750790; intellpositionT=1147px")

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
	// fmt.Println(string(body))

	// url := fmt.Sprintf("https://push2.eastmoney.com/api/qt/ulist.np/get?fltt=2&secids=%d.%d&fields=f164,f174,f252&ut=b2884a393a59ad64002292a3e90d46a5", prefix, secucode)
	// log.Infof("==>>TODO 881: %+v", url)
	// method := "GET"

	// client := &http.Client{}
	// req, err := http.NewRequest(method, url, nil)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// req.Header.Add("Connection", "keep-alive")
	// req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"")
	// req.Header.Add("sec-ch-ua-mobile", "?0")
	// req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	// req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	// req.Header.Add("Accept", "*/*")
	// req.Header.Add("Sec-Fetch-Site", "same-site")
	// req.Header.Add("Sec-Fetch-Mode", "no-cors")
	// req.Header.Add("Sec-Fetch-Dest", "script")
	// req.Header.Add("Referer", "https://data.eastmoney.com/")
	// req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	// req.Header.Add("Cookie", "qgqp_b_id=c3e233a777a7a2ed2872326c57200a01; em_hq_fls=js; st_si=42383655869778; cowCookie=true; intellpositionL=1080px; cowminicookie=true; st_asi=delete; intellpositionT=455px; HAList=a-sh-600941-%u4E2D%u56FD%u79FB%u52A8%2Ca-sh-600519-%u8D35%u5DDE%u8305%u53F0; st_pvi=81631919063174; st_sp=2022-01-12%2022%3A40%3A09; st_inirUrl=https%3A%2F%2Fcn.bing.com%2F; st_sn=14; st_psi=20220112231314210-113300300815-0441393438")

	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// fmt.Println(string(body))

	// body = filterPrefix(body)
	// log.Infof("==>>TODO 886: %+v", string(body))
	err = json.Unmarshal(replacChar(body), resp)
	if err != nil {
		// log.Infof("==>>TODO 887: %+v", err)
		return fmt.Errorf("unmarshal response failed: %v", err)
	}
	// log.Infof("==>>TODO 888: %+v", resp)
	return nil
}

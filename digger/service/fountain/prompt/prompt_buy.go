package prompt

import (
	"net/http"
	"sync"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/request"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func PromptBuyList(in *gin.Context) {
	var req trpc.PromptBuyListReq
	resp := &trpc.PromptBuyListResp{
		Code: 21000,
		Data: &trpc.PromptBuyData{
			Items: make([]*trpc.PromptBuyItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{
		"Disabled": false,
	}
	// 查询提示买入列表
	results, err := orm.GPPromptBuyMgr.Find(query, int(req.GetLimit()), int(req.GetOffset()))
	if err != nil {
		log.Errorf("get prompt buy failed: %q", err)
		in.JSON(http.StatusForbidden, resp)
		return
	}
	total := orm.GPPromptBuyMgr.Count(query)
	resp.Data.Total = int32(total)

	wg := sync.WaitGroup{}
	wg.Add(len(results))
	for _, result := range results {
		result := result
		go func(wg *sync.WaitGroup) {
			item := getPromptBuyItem(result)
			resp.Data.Items = append(resp.Data.Items, item)
			wg.Done()
		}(&wg)
	}
	wg.Wait()

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func getPromptBuyItem(data *orm.GPPromptBuy) *trpc.PromptBuyItem {
	result := &trpc.PromptBuyItem{}
	result.Secucode = data.Secucode
	result.Name = data.Name
	result.MinPrice = data.MinPrice
	result.PresentPrice = getPresentPrice(data)
	result.PriceDiff = utils.Decimal(result.PresentPrice - result.MinPrice)
	result.CreateDate = time.Unix(data.Start, 0).Format("2006-01-02 15:04:05")
	return result
}

func getPresentPrice(data *orm.GPPromptBuy) float64 {
	now := time.Now()
	if (now.Hour() > 9 && now.Hour() < 15) || data.PresentPrice == 0 {
		price := request.GetSinaDayPrice(data.Secucode)
		data.PresentPrice = utils.Decimal(price)
		data.UpdateDate = time.Now().Unix()
		data.Save()
	}

	return utils.Decimal(data.PresentPrice)
}

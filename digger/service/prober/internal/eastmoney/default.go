package eastmoney

// 股东研究
type ShareholderResearch struct {
	Gdrs       []Holder      `json:"gdrs"`
	SdltgdDate []SdltgdDate  `json:"sdltgd_date"`
	SdgdDate   []SdgdDate    `json:"sdgd_date"`
	Sjkzr      []Sjkzr       `json:"sjkzr"`
	JgccDate   []JgccDate    `json:"jgcc_date"`
	Sdgdcgbd   []Sdgdcgbd    `json:"sdgdcgbd"`
	JjcgDate   []JjcgDate    `json:"jjcg_date"`
	Xsjj       []interface{} `json:"xsjj"`
	Sdltgd     []Sdltgd      `json:"sdltgd"`
	Ltgf       []Ltgf        `json:"ltgf"`
	Sdgd       []Sdgd        `json:"sdgd"`
	Jgcc       []Jgcc        `json:"jgcc"`
	Jjcg       []Jjcg        `json:"jjcg"`
}

// 公司公告
type NewsBulletin struct {
	Gszx Gszx   `json:"gszx"`
	Gsgg []Gsgg `json:"gsgg"`
}

// 操盘必读
type OperationsRequired struct {
	Test     interface{}   `json:"test"`
	Dstx     []Dstx        `json:"dstx"`
	Xsjj     []interface{} `json:"xsjj"`
	Hxtc     []Hxtc        `json:"hxtc"`
	Gszx     Gszx          `json:"gszx"`
	Gsgg     []Gsgg        `json:"gsgg"`
	Dzjy     []Dzjy        `json:"dzjy"`
	Rzrq     []Rzrq        `json:"rzrq"`
	Lhbd     []Lhbd        `json:"lhbd"`
	Gdrs     []Gdrs        `json:"gdrs"`
	ZyzbAbgq []ZyzbAbgq    `json:"zyzb_abgq"`
	ZyzbAnd  []ZyzbAnd     `json:"zyzb_and"`
	ZyzbAdjd []ZyzbAdjd    `json:"zyzb_adjd"`
	Zxzb1    string        `json:"zxzb1"`
	Zxzb2    string        `json:"zxzb2"`
	Jgyc     []Jgyc        `json:"jgyc"`
	BaseYear int           `json:"BaseYear"`
	JgycPic  []JgycPic     `json:"jgyc_pic"`
	Cym      string        `json:"cym"`
	Ybzy     []Ybzy        `json:"ybzy"`
}

// 股市列表
type StockList struct {
	Rc   int       `json:"rc"`
	Rt   int       `json:"rt"`
	Svr  int64     `json:"svr"`
	Lt   int       `json:"lt"`
	Full int       `json:"full"`
	Data *CodeData `json:"data"`
}
type Holder struct {
	SECUCODE           string  `json:"SECUCODE"`
	SECURITYCODE       string  `json:"SECURITY_CODE"`
	ENDDATE            string  `json:"END_DATE"`
	HOLDERTOTALNUM     float64 `json:"HOLDER_TOTAL_NUM"`
	TOTALNUMRATIO      float64 `json:"TOTAL_NUM_RATIO"`
	AVGFREESHARES      float64 `json:"AVG_FREE_SHARES"`
	AVGFREESHARESRATIO float64 `json:"AVG_FREESHARES_RATIO"`
	HOLDFOCUS          string  `json:"HOLD_FOCUS"`
	PRICE              float64 `json:"PRICE"`
	AVGHOLDAMT         float64 `json:"AVG_HOLD_AMT"`
	HOLDRATIOTOTAL     float64 `json:"HOLD_RATIO_TOTAL"`
	FREEHOLDRATIOTOTAL float64 `json:"FREEHOLD_RATIO_TOTAL"`
}
type SdltgdDate struct {
	SECUCODE string `json:"SECUCODE"`
	ENDDATE  string `json:"END_DATE"`
}
type SdgdDate struct {
	SECUCODE string `json:"SECUCODE"`
	ENDDATE  string `json:"END_DATE"`
}
type Sjkzr struct {
	SECUCODE     string      `json:"SECUCODE"`
	SECURITYCODE string      `json:"SECURITY_CODE"`
	HOLDERNAME   interface{} `json:"HOLDER_NAME"`
	HOLDRATIO    interface{} `json:"HOLD_RATIO"`
}
type JgccDate struct {
	SECUCODE   string `json:"SECUCODE"`
	REPORTDATE string `json:"REPORT_DATE"`
}
type Sdgdcgbd struct {
	ENDDATE      string  `json:"END_DATE"`
	HOLDERRANK   float64 `json:"HOLDER_RANK"`
	HOLDERNAME   string  `json:"HOLDER_NAME"`
	SHARESTYPE   string  `json:"SHARES_TYPE"`
	HOLDNUM      float64 `json:"HOLD_NUM"`
	HOLDNUMRATIO float64 `json:"HOLD_NUM_RATIO"`
	HOLDCHANGE   string  `json:"HOLD_CHANGE"`
	CHANGERATIO  float64 `json:"CHANGE_RATIO"`
	CHANGEREASON string  `json:"CHANGE_REASON"`
}
type JjcgDate struct {
	SECUCODE   string `json:"SECUCODE"`
	REPORTDATE string `json:"REPORT_DATE"`
}
type Sdltgd struct {
	SECUCODE         string  `json:"SECUCODE"`
	SECURITYCODE     string  `json:"SECURITY_CODE"`
	ENDDATE          string  `json:"END_DATE"`
	HOLDERRANK       float64 `json:"HOLDER_RANK"`
	HOLDERNAME       string  `json:"HOLDER_NAME"`
	HOLDERTYPE       string  `json:"HOLDER_TYPE"`
	SHARESTYPE       string  `json:"SHARES_TYPE"`
	HOLDNUM          float64 `json:"HOLD_NUM"`
	FREEHOLDNUMRATIO float64 `json:"FREE_HOLDNUM_RATIO"`
	HOLDNUMCHANGE    string  `json:"HOLD_NUM_CHANGE"`
	CHANGERATIO      float64 `json:"CHANGE_RATIO"`
}
type Ltgf struct {
	SECUCODE             string  `json:"SECUCODE"`
	SECURITYCODE         string  `json:"SECURITY_CODE"`
	ENDDATE              string  `json:"END_DATE"`
	HOLDNUMCOUNT         int64   `json:"HOLD_NUM_COUNT"`
	LIMITEDSHARES        float64 `json:"LIMITED_SHARES"`
	OTHERUNLIMITEDSHARES float64 `json:"OTHER_UNLIMITED_SHARES"`
	HOLDNUMRATIO         float64 `json:"HOLD_NUM_RATIO"`
	LIMITEDSHARESRATIO   float64 `json:"LIMITED_SHARES_RATIO"`
	UNLIMITEDSHARESRATIO float64 `json:"UNLIMITED_SHARES_RATIO"`
}
type Sdgd struct {
	SECUCODE      string  `json:"SECUCODE"`
	SECURITYCODE  string  `json:"SECURITY_CODE"`
	ENDDATE       string  `json:"END_DATE"`
	HOLDERRANK    float64 `json:"HOLDER_RANK"`
	HOLDERNAME    string  `json:"HOLDER_NAME"`
	SHARESTYPE    string  `json:"SHARES_TYPE"`
	HOLDNUM       float64 `json:"HOLD_NUM"`
	HOLDNUMRATIO  float64 `json:"HOLD_NUM_RATIO"`
	HOLDNUMCHANGE string  `json:"HOLD_NUM_CHANGE"`
	CHANGERATIO   float64 `json:"CHANGE_RATIO"`
}
type Jgcc struct {
	SECUCODE         string  `json:"SECUCODE"`
	SECURITYCODE     string  `json:"SECURITY_CODE"`
	REPORTDATE       string  `json:"REPORT_DATE"`
	ORGTYPE          string  `json:"ORG_TYPE"`
	TOTALORGNUM      float64 `json:"TOTAL_ORG_NUM"`
	TOTALFREESHARES  float64 `json:"TOTAL_FREE_SHARES"`
	TOTALSHARESRATIO float64 `json:"TOTAL_SHARES_RATIO"`
	ALLSHARESRATIO   float64 `json:"ALL_SHARES_RATIO"`
}
type Jjcg struct {
	ORGTYPE          string  `json:"ORG_TYPE"`
	SECUCODE         string  `json:"SECUCODE"`
	REPORTDATE       string  `json:"REPORT_DATE"`
	HOLDERCODE       string  `json:"HOLDER_CODE"`
	HOLDERNAME       string  `json:"HOLDER_NAME"`
	TOTALSHARES      float64 `json:"TOTAL_SHARES"`
	HOLDVALUE        float64 `json:"HOLD_VALUE"`
	TOTALSHARESRATIO float64 `json:"TOTALSHARES_RATIO"`
	FREESHARESRATIO  float64 `json:"FREESHARES_RATIO"`
	FREEMARKETCAP    float64 `json:"FREE_MARKET_CAP"`
	FREESHARES       float64 `json:"FREE_SHARES"`
	SECURITYCODE     string  `json:"SECURITY_CODE"`
	FUNDCODE         string  `json:"FUND_CODE"`
	FUNDDERIVECODE   string  `json:"FUND_DERIVECODE"`
	NETVALUERATIO    float64 `json:"NETVALUE_RATIO"`
}

type Items struct {
	UniqueURL    string      `json:"uniqueUrl"`
	URL          string      `json:"url"`
	Code         string      `json:"code"`
	InfoCode     string      `json:"infoCode"`
	RecordID     interface{} `json:"recordId"`
	Source       interface{} `json:"source"`
	Title        string      `json:"title"`
	UpdateTime   int         `json:"updateTime"`
	ShowDateTime int64       `json:"showDateTime"`
	PublishDate  int         `json:"publishDate"`
	SRatingName  interface{} `json:"sRatingName"`
	Summary      string      `json:"summary"`
}
type Data struct {
	Count     int     `json:"count"`
	Condition string  `json:"condition"`
	Total     int     `json:"total"`
	Items     []Items `json:"items"`
}
type Gszx struct {
	Code       int    `json:"code"`
	CostTime   int    `json:"costTime"`
	Message    string `json:"message"`
	Reserve    string `json:"reserve"`
	Data       Data   `json:"data"`
	OStr       string `json:"OStr"`
	CacheTime  string `json:"CacheTime"`
	APIStatus  int    `json:"ApiStatus"`
	APIMessage string `json:"ApiMessage"`
}
type Gsgg struct {
	ArtCode     string `json:"art_code"`
	DisplayTime string `json:"display_time"`
	NoticeDate  string `json:"notice_date"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}

type Dstx struct {
	DateNow                             string `json:"DateNow"`
	DisclosureAnnualReport              string `json:"DisclosureAnnualReport"`
	DisclosureTime                      string `json:"DisclosureTime"`
	BonusTimeFinancing                  string `json:"BonusTimeFinancing"`
	DividendInstructions                string `json:"DividendInstructions"`
	AnnouncementDate                    string `json:"AnnouncementDate"`
	DateOfMeeting                       string `json:"DateOfMeeting"`
	ShareholdersCategory                string `json:"ShareholdersCategory"`
	ShareholderAnnouncementDate         string `json:"ShareholderAnnouncementDate"`
	BanTime                             string `json:"BanTime"`
	TheNumberOfLifting                  string `json:"TheNumberOfLifting"`
	DistributionPlan                    string `json:"DistributionPlan"`
	AnnouncementTime                    string `json:"AnnouncementTime"`
	NoticeOfTheContentsOfThePerformance string `json:"NoticeOfTheContentsOfThePerformance"`
	ReasonsForChangesInPerformance      string `json:"ReasonsForChangesInPerformance"`
	TotalEquity                         string `json:"TotalEquity"`
	SuspensionDate                      string `json:"suspensionDate"`
	SuspensionReason                    string `json:"suspensionReason"`
	DateOfRecordA                       string `json:"DateOfRecordA"`
	DateOfRecordB                       string `json:"DateOfRecordB"`
	SharesDividendDate                  string `json:"SharesDividendDate"`
	DividendDateA                       string `json:"DividendDateA"`
	DividendDateB                       string `json:"DividendDateB"`
	Allocation                          string `json:"Allocation"`
	AllocationDate                      string `json:"AllocationDate"`
	StockType                           string `json:"StockType"`
	Bbpl                                string `json:"bbpl"`
	RecDate                             string `json:"RecDate"`
	LatestDividendReportDate            string `json:"LatestDividendReportDate"`
	LatestDistributionReportDate        string `json:"LatestDistributionReportDate"`
	YGND                                string `json:"YGND"`
	ReportDate                          string `json:"ReportDate"`
	DisclosureAnnualReport2             string `json:"DisclosureAnnualReport2"`
	DisclosureTime2                     string `json:"DisclosureTime2"`
	Bbpl2                               string `json:"bbpl2"`
	RecDate2                            string `json:"RecDate2"`
	ImplementDate                       string `json:"ImplementDate"`
}
type Hxtc struct {
	Zqnm   string `json:"zqnm"`
	Zqdm   string `json:"zqdm"`
	Zqjc   string `json:"zqjc"`
	Jyscbm string `json:"jyscbm"`
	Gjc    string `json:"gjc"`
	Yd     string `json:"yd"`
	Ydnr   string `json:"ydnr"`
}

type Dzjy struct {
	Jyrq  string `json:"jyrq"`
	Cjj   string `json:"cjj"`
	Cjl   string `json:"cjl"`
	Cjje  string `json:"cjje"`
	Mryyb string `json:"mryyb"`
	Mcyyb string `json:"mcyyb"`
}
type Rzrq struct {
	Sj    string `json:"sj"`
	Rzmre string `json:"rzmre"`
	Rzche string `json:"rzche"`
	Rzye  string `json:"rzye"`
	Rqmcl string `json:"rqmcl"`
	Rqchl string `json:"rqchl"`
	Rqye  string `json:"rqye"`
}
type GroupMr struct {
	Yybmc  string `json:"yybmc"`
	Mrje   string `json:"mrje"`
	ZjebMr string `json:"zjeb_mr"`
	Mcje   string `json:"mcje"`
	ZjebMc string `json:"zjeb_mc"`
}
type GroupMc struct {
	Yybmc  string `json:"yybmc"`
	Mrje   string `json:"mrje"`
	ZjebMr string `json:"zjeb_mr"`
	Mcje   string `json:"mcje"`
	ZjebMc string `json:"zjeb_mc"`
}
type Lhbd struct {
	GroupMr  []GroupMr `json:"group_mr"`
	GroupMc  []GroupMc `json:"group_mc"`
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Rq       string    `json:"rq"`
	Zl       string    `json:"zl"`
	Mrzj     string    `json:"mrzj"`
	Mrzjzjeb string    `json:"mrzjzjeb"`
	Mczj     string    `json:"mczj"`
	Mczjzjeb string    `json:"mczjzjeb"`
}
type Gdrs struct {
	Rq          string `json:"rq"`
	Gdrs        string `json:"gdrs"`
	GdrsJsqbh   string `json:"gdrs_jsqbh"`
	Rjltg       string `json:"rjltg"`
	RjltgJsqbh  string `json:"rjltg_jsqbh"`
	Cmjzd       string `json:"cmjzd"`
	Gj          string `json:"gj"`
	Rjcgje      string `json:"rjcgje"`
	Qsdgdcghj   string `json:"qsdgdcghj"`
	Qsdltgdcghj string `json:"qsdltgdcghj"`
}
type ZyzbAbgq struct {
	Rq          string `json:"rq"`
	Jbmgsy      string `json:"jbmgsy"`
	Kfmgsy      string `json:"kfmgsy"`
	Xsmgsy      string `json:"xsmgsy"`
	Gsjlr       string `json:"gsjlr"`
	Gsjlrtbzz   string `json:"gsjlrtbzz"`
	Gsjlrgdhbzz string `json:"gsjlrgdhbzz"`
	Jqjzcsyl    string `json:"jqjzcsyl"`
	Tbjzcsyl    string `json:"tbjzcsyl"`
	Mll         string `json:"mll"`
	Sjsl        string `json:"sjsl"`
	Ysk         string `json:"ysk"`
	Xsxjl       string `json:"xsxjl"`
	Zzczzl      string `json:"zzczzl"`
	Zcfzl       string `json:"zcfzl"`
	Ldfz        string `json:"ldfz"`
}
type ZyzbAnd struct {
	Rq          string `json:"rq"`
	Jbmgsy      string `json:"jbmgsy"`
	Kfmgsy      string `json:"kfmgsy"`
	Xsmgsy      string `json:"xsmgsy"`
	Gsjlr       string `json:"gsjlr"`
	Gsjlrtbzz   string `json:"gsjlrtbzz"`
	Gsjlrgdhbzz string `json:"gsjlrgdhbzz"`
	Jqjzcsyl    string `json:"jqjzcsyl"`
	Tbjzcsyl    string `json:"tbjzcsyl"`
	Mll         string `json:"mll"`
	Sjsl        string `json:"sjsl"`
	Ysk         string `json:"ysk"`
	Xsxjl       string `json:"xsxjl"`
	Zzczzl      string `json:"zzczzl"`
	Zcfzl       string `json:"zcfzl"`
	Ldfz        string `json:"ldfz"`
}
type ZyzbAdjd struct {
	Rq        string `json:"rq"`
	Jbmgsy    string `json:"jbmgsy"`
	Gsjlr     string `json:"gsjlr"`
	Gsjlrtbzz string `json:"gsjlrtbzz"`
	Gsjlrhbzz string `json:"gsjlrhbzz"`
	Tbzcsyl   string `json:"tbzcsyl"`
	Mll       string `json:"mll"`
}
type Jgyc struct {
	Jgmc string `json:"jgmc"`
	Sy1  string `json:"sy1"`
	Syl1 string `json:"syl1"`
	Sy2  string `json:"sy2"`
	Syl2 string `json:"syl2"`
	Sy3  string `json:"sy3"`
	Syl3 string `json:"syl3"`
	Sy4  string `json:"sy4"`
	Syl4 string `json:"syl4"`
}
type JgycPic struct {
	Nf     string `json:"nf"`
	Mgsy   string `json:"mgsy"`
	Mgsyzz string `json:"mgsyzz"`
}
type Ybzy struct {
	AimPrice            string `json:"aim_price"`
	ArtCode             string `json:"art_code"`
	EmRatingCode        string `json:"em_rating_code"`
	EmRatingName        string `json:"em_rating_name"`
	InduOldIndustryCode string `json:"indu_old_industry_code"`
	InduOldIndustryName string `json:"indu_old_industry_name"`
	PublishTime         string `json:"publish_time"`
	Rating              string `json:"rating"`
	RatingChange        string `json:"rating_change"`
	ReportType          string `json:"report_type"`
	SRatingName         string `json:"s_rating_name"`
	Source              string `json:"source"`
	Title               string `json:"title"`
	Content             string `json:"content"`
}

type CodeBase struct {
	Code string `json:"f12"`
	Name string `json:"f14"`
}
type CodeData struct {
	Total int         `json:"total"`
	Diff  []*CodeBase `json:"diff"`
}

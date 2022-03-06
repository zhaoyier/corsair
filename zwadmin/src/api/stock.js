import request from '@/utils/request'

// 查询配置参数
export function getCNConfig(data) {
  return request({
    url: '/api/backend/GetCNConfig',
    method: 'post',
    data
  })
}

// 配置参数
export function updateCNConfig(data) {
  return request({
    url: '/api/backend/updateCNConfig',
    method: 'post',
    data
  })
}

export function getRecommendList(data) {
  return request({
    url: '/api/stock/GetRecommend',
    method: 'post',
    data
  })
}

// 更新
export function updateRecommend(data) {
  return request({
    url: '/api/stock/UpdateRecommend',
    method: 'post',
    data
  })
}

// 更新
export function promptBuyList(data) {
  return request({
    url: '/api/stock/PromptBuyList',
    method: 'post',
    data
  })
}

// 查询股东人数变化
export function getLongLineList(data) {
  return request({
    url: '/api/stock/GetLongLineList',
    method: 'post',
    data
  })
}

// 查询每日数据
export function getDailyList(data) {
  return request({
    url: '/api/stock/GetDailyList',
    method: 'post',
    data
  })
}

// 查询手动调整幅度列表
export function manualDecreaseList(data) {
  return request({
    url: '/api/stock/ManualDecreaseList',
    method: 'post',
    data
  })
}

// 关注确认
export function confirmFocus(data) {
  return request({
    url: '/api/stock/ConfirmFocus',
    method: 'post',
    data
  })
}

// 关注列表
export function getFocusList(data) {
  return request({
    url: '/api/stock/GetFocusList',
    method: 'post',
    data
  })
}

// 取消关注
export function cancelFocus(data) {
  return request({
    url: '/api/stock/CancelFocus',
    method: 'post',
    data
  })
}

// 更新关注
export function updateFocus(data) {
  return request({
    url: '/api/stock/updateFocus',
    method: 'post',
    data
  })
}

// 股东人数列表
export function gdrenshuList(data) {
  return request({
    url: '/api/stock/GDRenshuList',
    method: 'post',
    data
  })
}

// 个股股东列表
export function gdrenshuDetail(data) {
  return request({
    url: '/api/stock/GDRenshuDetail',
    method: 'post',
    data
  })
}

// 新增周期
export function addGPZhouQi(data) {
  return request({
    url: '/api/stock/AddGPZhouQi',
    method: 'post',
    data
  })
}

// 更新周期
export function updateGPZhouQi(data) {
  return request({
    url: '/api/stock/UpdateGPZhouQi',
    method: 'post',
    data
  })
}

// 更新周期备注
export function addGPRemark(data) {
  return request({
    url: '/api/stock/AddGPRemark',
    method: 'post',
    data
  })
}

// 周期列表
export function gpzhouQiList(data) {
  return request({
    url: '/api/stock/GPZhouQiList',
    method: 'post',
    data
  })
}

// 股东聚合设置
export function gdaggregationReset(data) {
  return request({
    url: '/api/stock/GDAggregationReset',
    method: 'post',
    data
  })
}

// 股东聚合设置
export function gdaggregationList(data) {
  return request({
    url: '/api/stock/GDAggregationList',
    method: 'post',
    data
  })
}

// 资金流入
export function getFundFlowList(data) {
  return request({
    url: '/api/stock/GetFundFlowList',
    method: 'post',
    data
  })
}
// GetFundDetailList
export function getFundDetailList(data) {
  return request({
    url: '/api/stock/GetFundDetailList',
    method: 'post',
    data
  })
}


// 查询十大股东
export function getGDSDLT(data) {
  return request({
    url: '/api/stock/GetGDSDLT',
    method: 'post',
    data
  })
}

// 查询十大股东详情
export function getGDSDLTDetail(data) {
  return request({
    url: '/api/stock/GetGDSDLTDetail',
    method: 'post',
    data
  })
}

// 查询大变的数据
export function getWaterfallList(data) {
  return request({
    url: '/api/stock/GetWaterfallList',
    method: 'post',
    data
  })
}

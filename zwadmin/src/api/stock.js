import request from '@/utils/request'

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

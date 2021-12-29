import request from '@/utils/request'

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
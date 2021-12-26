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
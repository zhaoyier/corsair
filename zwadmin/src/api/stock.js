import request from '@/utils/request'

export function getRecommendList(data) {
  return request({
    url: '/api/stock/GetRecommend',
    method: 'post',
    data
  })
}

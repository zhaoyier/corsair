import request from '@/utils/request'

export function login(data) {
  return request({
    // url: '/vue-admin-template/user/login',
    url: '/api/backend/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    // url: '/vue-admin-template/user/info',
    url: '/api/backend/info',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    // url: '/vue-admin-template/user/logout',
    url: '/api/backend/logout',
    method: 'post'
  })
}

// 更新配置
export function updateCNConfig() {
  return request({
    url: '/api/backend/UpdateCNConfig',
    method: 'post'
  })
}
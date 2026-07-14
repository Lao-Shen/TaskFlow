/**
 * API 请求工具
 *
 * 用法：
 *   import api from '@/api'
 *   const res = await api.login({ username, password })
 */

// 开发环境走 Vite proxy，生产环境与后端同域部署
const BASE_URL = '/api'

const TOKEN_KEY = 'token'

// ============================================================
//  基础封装
// ============================================================

async function request(url, options = {}) {
  const token = localStorage.getItem(TOKEN_KEY)

  const headers = {
    'Content-Type': 'application/json',
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
    ...options.headers,
  }

  const res = await fetch(`${BASE_URL}${url}`, { ...options, headers })
  const data = await res.json()

  if (!res.ok || data.code !== 0) {
    throw new ApiError(data.msg || '请求失败', data.code || res.status, data)
  }

  return data
}

class ApiError extends Error {
  constructor(message, code, data) {
    super(message)
    this.name = 'ApiError'
    this.code = code
    this.data = data
  }
}

// ============================================================
//  授权
// ============================================================

function getToken() {
  return localStorage.getItem(TOKEN_KEY)
}

function setToken(token) {
  localStorage.setItem(TOKEN_KEY, token)
}

function removeToken() {
  localStorage.removeItem(TOKEN_KEY)
}

function isLoggedIn() {
  return !!getToken()
}

// ============================================================
//  接口方法
// ============================================================

const api = {
  /** POST /api/register — 注册 */
  register(data) {
    return request('/register', {
      method: 'POST',
      body: JSON.stringify(data),
    })
  },

  /** POST /api/login — 登录 */
  login(data) {
    return request('/login', {
      method: 'POST',
      body: JSON.stringify(data),
    })
  },

  /** POST /api/logout — 登出 */
  logout() {
    return request('/logout', { method: 'POST' })
  },

  /** GET /api/user/me — 获取当前用户 */
  getMe() {
    return request('/user/me')
  },

  /** POST /api/tasks — 创建任务 */
  createTask(data) {
    return request('/tasks', {
      method: 'POST',
      body: JSON.stringify(data),
    })
  },

  /** GET /api/tasks — 获取任务列表（含统计） */
  listTasks(status) {
    const qs = status ? `?status=${encodeURIComponent(status)}` : ''
    return request(`/tasks${qs}`)
  },
}

export { getToken, setToken, removeToken, isLoggedIn, ApiError }
export default api
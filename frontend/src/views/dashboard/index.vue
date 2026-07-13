<template>
  <div class="dashboard">
    <!-- ===== 顶部导航栏 ===== -->
    <header class="topbar">
      <div class="topbar__inner">
        <!-- 左侧：Logo -->
        <div class="topbar__left">
          <span class="topbar__logo">⚡</span>
          <span class="topbar__brand">TaskFlow</span>
        </div>

        <!-- 右侧：通知 + 用户信息 -->
        <div class="topbar__right">
          <button class="topbar__notify" title="消息通知">
            <span class="topbar__notify-dot"></span>
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"/>
              <path d="M13.73 21a2 2 0 0 1-3.46 0"/>
            </svg>
          </button>

          <!-- 用户信息区 -->
          <div class="topbar__user" @click="showUserMenu = !showUserMenu">
            <span v-if="loadingUser" class="topbar__avatar-fallback">…</span>
            <template v-else>
              <img
                v-if="user.avatar"
                class="topbar__avatar"
                :src="user.avatar"
                :alt="user.name"
                @error="onAvatarError"
              />
              <span v-else class="topbar__avatar-fallback">{{ user.name.charAt(0) }}</span>
            </template>
            <div class="topbar__user-meta">
              <span class="topbar__username">{{ loadingUser ? '加载中…' : user.name }}</span>
              <span class="topbar__role">{{ user.email }}</span>
            </div>
            <svg class="topbar__chevron" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
              <polyline points="6 9 12 15 18 9"/>
            </svg>

            <!-- 下拉菜单 -->
            <Transition name="menu">
              <div v-if="showUserMenu" class="topbar__dropdown" @click.stop>
                <div class="topbar__dropdown-item" @click="handleLogout">退出登录</div>
              </div>
            </Transition>
          </div>
        </div>
      </div>
    </header>

    <!-- ===== Hero 横条 ===== -->
    <section class="hero">
      <div class="hero__inner">
        <div class="hero__text">
          <h1 class="hero__greeting">{{ greeting }}，{{ user.name }}</h1>
          <p class="hero__subtitle">欢迎回到 TaskFlow，这是你的任务总览</p>
        </div>
        <div class="hero__date">{{ today }}</div>
      </div>
    </section>

    <!-- ===== 统计卡片 ===== -->
    <section class="stats">
      <div
        class="stat-card"
        v-for="s in stats"
        :key="s.label"
      >
        <span class="stat-card__icon">{{ s.icon }}</span>
        <div class="stat-card__body">
          <span class="stat-card__value">{{ s.value }}</span>
          <span class="stat-card__label">{{ s.label }}</span>
        </div>
      </div>
    </section>

    <!-- ===== 任务面板 ===== -->
    <section class="section">
      <div class="section__header">
        <h2 class="section__title">📌 任务面板</h2>
        <span class="section__count">{{ cards.length }} 个分组</span>
      </div>

      <div class="card-grid">
        <div
          class="card-grid__item"
          v-for="item in cards"
          :key="item.id"
          @click="openCard(item)"
        >
          <card v-bind="item" />
        </div>
      </div>
    </section>

    <!-- ===== 弹窗 ===== -->
    <Transition name="modal">
      <div v-if="showModal" class="modal-backdrop" @click.self="closeModal">
        <div class="modal">
          <button class="modal__close" @click="closeModal" title="关闭">✕</button>
          <div class="modal__icon">{{ activeCard?.icon }}</div>
          <h3 class="modal__title">{{ activeCard?.title }}</h3>
          <p class="modal__desc">{{ activeCard?.description }}</p>
          <div class="modal__actions">
            <button class="modal__btn modal__btn--primary" @click="closeModal">进入查看</button>
            <button class="modal__btn modal__btn--secondary" @click="closeModal">取消</button>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import card from '@/components/task/card.vue'
import api, { removeToken } from '@/api'

const router = useRouter()

// ============================================================
//  用户数据  ——  从后端加载
// ============================================================

const user = ref({
  name: '',
  email: '',
  avatar: '',
})

const loadingUser = ref(true)
const showUserMenu = ref(false)

async function loadUser() {
  try {
    const res = await api.getMe()
    user.value.name = res.data.username
    user.value.email = res.data.email
  } catch {
    // token 无效，跳转登录
    removeToken()
    router.push('/login')
  } finally {
    loadingUser.value = false
  }
}

/** 头像加载失败时回退到首字母 */
function onAvatarError() {
  user.value.avatar = ''
}

/** 登出 */
async function handleLogout() {
  try {
    await api.logout()
  } catch {
    // 即使服务端登出失败也清除本地 token
  }
  removeToken()
  router.push('/login')
}

/** 根据当前时间段生成问候语 */
const greeting = computed(() => {
  const h = new Date().getHours()
  if (h < 6) return '夜深了'
  if (h < 12) return '早上好'
  if (h < 14) return '中午好'
  if (h < 18) return '下午好'
  return '晚上好'
})

const today = new Date().toLocaleDateString('zh-CN', {
  year: 'numeric',
  month: 'long',
  day: 'numeric',
  weekday: 'long',
})

/** 统计面板（假数据） */
const stats = ref([
  { icon: '📊', value: 12, label: '全部任务' },
  { icon: '⏳', value: 5,  label: '待办' },
  { icon: '🔄', value: 4,  label: '进行中' },
  { icon: '✅', value: 3,  label: '已完成' },
])

/** 任务卡片列表（假数据） */
const cards = ref([
  { id: 1, title: '待办任务', description: '查看所有待处理的任务项，合理安排优先级',   icon: '📋', badge: '5',  variant: 'todo' },
  { id: 2, title: '进行中',   description: '跟踪当前正在推进的任务，掌握最新进度',       icon: '🔄', badge: '4',  variant: 'progress' },
  { id: 3, title: '已完成',   description: '回顾已经完成的任务，总结经验与收获',         icon: '✅', badge: '3',  variant: 'done' },
  { id: 4, title: '归档',     description: '浏览历史归档记录，便于日后检索与复盘',       icon: '📦', badge: '12', variant: 'archive' },
])

// ============================================================
//  数据加载  ——  接入后端后放开此函数并替换假数据
// ============================================================
// import { fetchDashboard, fetchUser } from '@/api/dashboard'
//
// async function loadDashboard() {
//   const [userRes, dashRes] = await Promise.all([
//     fetchUser(),          // GET /api/user/me
//     fetchDashboard(),     // GET /api/dashboard
//   ])
//   user.value  = userRes.data
//   stats.value = dashRes.data.stats
//   cards.value = dashRes.data.cards
// }

/** 点击外部关闭下拉 */
function onClickOutside(e) {
  if (!e.target.closest('.topbar__user')) {
    showUserMenu.value = false
  }
}

onMounted(() => {
  loadUser()
  document.addEventListener('click', onClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', onClickOutside)
})

// ============================================================
//  弹窗交互
// ============================================================
const showModal = ref(false)
const activeCard = ref(null)

function openCard(item) {
  activeCard.value = item
  showModal.value = true
}

function closeModal() {
  showModal.value = false
  activeCard.value = null
}
</script>

<style scoped>
/* ============================================================
   CSS 变量
   ============================================================ */
.dashboard {
  --color-bg:        #f0f3f8;
  --color-surface:   #ffffff;
  --color-primary:   #4f6ef7;
  --color-text:      #1d1d1f;
  --color-secondary: #86868b;
  --color-border:    #e8eaef;
  --radius-lg:       12px;
  --radius-xl:       20px;
  --shadow-sm:       0 2px 12px rgba(0,0,0,.05);
  --shadow-md:       0 4px 20px rgba(0,0,0,.06);
  --shadow-lg:       0 24px 80px rgba(0,0,0,.18);

  min-height: 100vh;
  background: linear-gradient(160deg, #f0f3f8 0%, #f7f9fc 40%, #f5f6fa 100%);
  color: var(--color-text);
}

/* ============================================================
   顶部导航栏
   ============================================================ */
.topbar {
  position: sticky;
  top: 0;
  z-index: 100;
  background: rgba(255,255,255,.78);
  backdrop-filter: blur(12px);
  border-bottom: 1px solid var(--color-border);
}

.topbar__inner {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 32px;
  height: 56px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

/* 左侧品牌 */
.topbar__left {
  display: flex;
  align-items: center;
  gap: 8px;
}

.topbar__logo {
  font-size: 24px;
}

.topbar__brand {
  font-size: 18px;
  font-weight: 700;
  letter-spacing: -0.3px;
}

/* 右侧 */
.topbar__right {
  display: flex;
  align-items: center;
  gap: 12px;
}

/* 通知按钮 */
.topbar__notify {
  position: relative;
  width: 36px;
  height: 36px;
  border: none;
  border-radius: 10px;
  background: transparent;
  color: #555;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background .2s;
}

.topbar__notify:hover {
  background: #f0f2f5;
}

.topbar__notify-dot {
  position: absolute;
  top: 7px;
  right: 8px;
  width: 7px;
  height: 7px;
  border-radius: 50%;
  background: #f56c6c;
  border: 2px solid #fff;
}

/* 用户区 */
.topbar__user {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 4px 8px 4px 4px;
  border-radius: 28px;
  cursor: pointer;
  transition: background .2s;
}

.topbar__user:hover {
  background: #f0f2f5;
}

.topbar__avatar {
  width: 34px;
  height: 34px;
  border-radius: 50%;
  object-fit: cover;
  background: linear-gradient(135deg, #4f6ef7, #7c5cfc);
  flex-shrink: 0;
}

/* 头像加载失败时显示的占位首字母 */
.topbar__avatar-fallback {
  width: 34px;
  height: 34px;
  border-radius: 50%;
  background: linear-gradient(135deg, #4f6ef7, #7c5cfc);
  color: #fff;
  font-size: 15px;
  font-weight: 600;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.topbar__user-meta {
  display: flex;
  flex-direction: column;
  line-height: 1.3;
}

.topbar__username {
  font-size: 14px;
  font-weight: 600;
}

.topbar__role {
  font-size: 11px;
  color: var(--color-secondary);
}

.topbar__chevron {
  color: #aaa;
  margin-left: 2px;
  transition: transform .2s;
}

/* 下拉菜单 */
.topbar__user {
  position: relative;
}

.topbar__dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 8px;
  min-width: 120px;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 8px 30px rgba(0,0,0,.12);
  overflow: hidden;
  z-index: 200;
}

.topbar__dropdown-item {
  padding: 10px 16px;
  font-size: 13px;
  color: #606266;
  cursor: pointer;
  transition: background .15s, color .15s;
}

.topbar__dropdown-item:hover {
  background: #f0f2f5;
  color: #f56c6c;
}

/* 下拉动画 */
.menu-enter-active,
.menu-leave-active {
  transition: opacity .15s, transform .15s;
}

.menu-enter-from,
.menu-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}

/* ============================================================
   Hero 横条
   ============================================================ */
.hero {
  background: linear-gradient(135deg, #2c3e50 0%, #3b5998 50%, #4a6fa5 100%);
  color: #fff;
}

.hero__inner {
  max-width: 1200px;
  margin: 0 auto;
  padding: 28px 32px;
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  gap: 16px;
}

.hero__greeting {
  margin: 0;
  font-size: 24px;
  font-weight: 700;
  letter-spacing: -0.3px;
}

.hero__subtitle {
  margin: 4px 0 0;
  font-size: 14px;
  opacity: .78;
}

.hero__date {
  font-size: 13px;
  opacity: .7;
  white-space: nowrap;
}

/* ============================================================
   统计卡片
   ============================================================ */
.stats {
  max-width: 1200px;
  margin: -20px auto 0;
  padding: 0 32px;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
  position: relative;
  z-index: 1;
}

.stat-card {
  background: var(--color-surface);
  border-radius: var(--radius-lg);
  padding: 18px 20px;
  display: flex;
  align-items: center;
  gap: 14px;
  box-shadow: var(--shadow-md);
  transition: transform .2s, box-shadow .2s;
}

.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 28px rgba(0,0,0,.1);
}

.stat-card__icon  { font-size: 30px; line-height: 1; }
.stat-card__body  { display: flex; flex-direction: column; }
.stat-card__value { font-size: 24px; font-weight: 700; line-height: 1.2; }
.stat-card__label { font-size: 13px; color: var(--color-secondary); margin-top: 2px; }

/* ============================================================
   任务面板
   ============================================================ */
.section {
  max-width: 1200px;
  margin: 32px auto 0;
  padding: 0 32px 40px;
}

.section__header {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  margin-bottom: 20px;
}

.section__title {
  margin: 0;
  font-size: 19px;
  font-weight: 700;
}

.section__count {
  font-size: 13px;
  color: #909399;
}

/* ============================================================
   卡片网格
   ============================================================ */
.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

.card-grid__item { cursor: pointer; }

/* ============================================================
   弹窗
   ============================================================ */
.modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,.45);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  backdrop-filter: blur(4px);
}

.modal {
  background: #fff;
  border-radius: var(--radius-xl);
  padding: 36px 32px 28px;
  width: min(480px, 90%);
  box-shadow: var(--shadow-lg);
  text-align: center;
  position: relative;
}

.modal__close {
  position: absolute;
  top: 14px;
  right: 16px;
  width: 32px;
  height: 32px;
  border: none;
  border-radius: 50%;
  background: #f2f3f5;
  color: #606266;
  font-size: 16px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background .2s, color .2s;
}

.modal__close:hover { background: #e0e1e5; color: #303133; }

.modal__icon  { font-size: 48px; margin-bottom: 12px; }
.modal__title { margin: 0 0 10px; font-size: 22px; font-weight: 700; }
.modal__desc  { margin: 0; color: var(--color-secondary); font-size: 15px; line-height: 1.7; }

.modal__actions {
  display: flex;
  gap: 12px;
  justify-content: center;
  margin-top: 28px;
}

.modal__btn {
  padding: 10px 28px;
  border: none;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
  transition: opacity .2s, transform .15s;
}

.modal__btn:active { transform: scale(.97); }

.modal__btn--primary {
  background: linear-gradient(135deg, #409eff, #337ecc);
  color: #fff;
}

.modal__btn--primary:hover { opacity: .9; }

.modal__btn--secondary {
  background: #f2f3f5;
  color: #606266;
}

.modal__btn--secondary:hover { background: #e5e6eb; }

/* 弹窗动画 */
.modal-enter-active,
.modal-leave-active { transition: opacity .25s; }

.modal-enter-active .modal,
.modal-leave-active .modal { transition: transform .25s, opacity .25s; }

.modal-enter-from,
.modal-leave-to { opacity: 0; }

.modal-enter-from .modal { transform: scale(.92) translateY(20px); opacity: 0; }
.modal-leave-to   .modal { transform: scale(.92) translateY(20px); opacity: 0; }

/* ============================================================
   响应式
   ============================================================ */
@media (max-width: 768px) {
  .topbar__inner {
    padding: 0 20px;
  }

  .topbar__user-meta,
  .topbar__chevron {
    display: none;
  }

  .stats {
    grid-template-columns: repeat(2, 1fr);
  }

  .hero__inner {
    flex-direction: column;
    align-items: flex-start;
    padding: 20px;
  }

  .hero__greeting { font-size: 20px; }

  .section,
  .stats {
    padding-left: 20px;
    padding-right: 20px;
  }

  .card-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 480px) {
  .stats {
    grid-template-columns: 1fr;
  }
}
</style>

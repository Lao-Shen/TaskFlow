<template>
  <div class="dashboard">
    <!-- ===== 顶部导航栏 ===== -->
    <header class="topbar">
      <div class="topbar__inner">
        <div class="topbar__left">
          <span class="topbar__logo">⚡</span>
          <span class="topbar__brand">TaskFlow</span>
        </div>

        <div class="topbar__right">
          <button class="topbar__notify" title="消息通知">
            <span class="topbar__notify-dot"></span>
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"/>
              <path d="M13.73 21a2 2 0 0 1-3.46 0"/>
            </svg>
          </button>

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
        :class="{ 'stat-card--active': statModal.key === s.key }"
        v-for="s in stats"
        :key="s.key"
        @click="openStatModal(s)"
      >
        <span class="stat-card__icon">{{ s.icon }}</span>
        <div class="stat-card__body">
          <span class="stat-card__value">{{ s.value }}</span>
          <span class="stat-card__label">{{ s.label }}</span>
        </div>
      </div>
    </section>

    <!-- ===== 平台功能面板 ===== -->
    <section class="section">
      <div class="section__header">
        <h2 class="section__title">⚙️ 平台功能</h2>
        <span class="section__count">{{ cards.length }} 个工具</span>
      </div>

      <div class="card-grid">
        <div
          class="card-grid__item"
          v-for="item in cards"
          :key="item.id"
          @click="openCardModal(item)"
        >
          <card v-bind="item" />
        </div>
      </div>
    </section>

    <!-- ===== 统计卡片弹窗（含任务列表） ===== -->
    <Transition name="modal">
      <div v-if="showStatModal" class="modal-backdrop" @click.self="closeStatModal">
        <div class="modal modal--wide">
          <button class="modal__close" @click="closeStatModal" title="关闭">✕</button>
          <div class="modal__icon">{{ statModal.icon }}</div>
          <h3 class="modal__title">{{ statModal.label }}</h3>

          <!-- 任务表格 -->
          <table class="task-table" v-if="statModal.tasks.length">
            <thead>
              <tr>
                <th>任务名称</th>
                <th>任务类型</th>
                <th>状态</th>
                <th>创建时间</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="t in statModal.tasks" :key="t.id">
                <td class="task-table__name">{{ t.name }}</td>
                <td>{{ t.type }}</td>
                <td>
                  <span class="task-table__status" :class="`task-table__status--${t.statusClass}`">
                    {{ t.status }}
                  </span>
                </td>
                <td class="task-table__time">{{ t.time }}</td>
              </tr>
            </tbody>
          </table>

          <p v-else class="modal__empty">暂无任务</p>

          <div class="modal__actions">
            <button class="modal__btn modal__btn--secondary" @click="closeStatModal">关闭</button>
          </div>
        </div>
      </div>
    </Transition>

    <!-- ===== 功能卡片弹窗 ===== -->
    <Transition name="modal">
      <div v-if="showCardModal" class="modal-backdrop" @click.self="closeCardModal">
        <div class="modal">
          <button class="modal__close" @click="closeCardModal" title="关闭">✕</button>
          <div class="modal__icon">{{ activeCard?.icon }}</div>
          <h3 class="modal__title">{{ activeCard?.title }}</h3>
          <p class="modal__desc">{{ activeCard?.description }}</p>
          <div class="modal__actions">
            <button class="modal__btn modal__btn--primary" @click="closeCardModal">进入功能</button>
            <button class="modal__btn modal__btn--secondary" @click="closeCardModal">取消</button>
          </div>
        </div>
      </div>
    </Transition>

    <!-- ===== 临时：随机创建任务按钮 ===== -->
    <button class="fab" title="随机创建任务" @click="createRandomTask" :disabled="creatingTask">
      <span v-if="!creatingTask">🎲</span>
      <span v-else class="fab__spinner"></span>
    </button>
    <Transition name="toast">
      <div v-if="toast.show" class="toast" :class="`toast--${toast.type}`">
        {{ toast.msg }}
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import card from '@/components/task/card.vue'
import api, { removeToken, ApiError } from '@/api'

const router = useRouter()

// ============================================================
//  用户数据
// ============================================================
const user = ref({ name: '', email: '', avatar: '' })
const loadingUser = ref(true)
const showUserMenu = ref(false)

async function loadUser() {
  try {
    const res = await api.getMe()
    user.value.name = res.data.username
    user.value.email = res.data.email
  } catch {
    removeToken()
    router.push('/login')
  } finally {
    loadingUser.value = false
  }
}

function onAvatarError() { user.value.avatar = '' }

async function handleLogout() {
  try { await api.logout() } catch { /* ignore */ }
  removeToken()
  router.push('/login')
}

const greeting = computed(() => {
  const h = new Date().getHours()
  if (h < 6) return '夜深了'
  if (h < 12) return '早上好'
  if (h < 14) return '中午好'
  if (h < 18) return '下午好'
  return '晚上好'
})

const today = new Date().toLocaleDateString('zh-CN', {
  year: 'numeric', month: 'long', day: 'numeric', weekday: 'long',
})

// ============================================================
//  任务数据  ——  从后端加载
// ============================================================
const tasks = ref([])                  // 已映射的显示任务列表
const loadingTasks = ref(false)

// 后端状态 → 显示名
const STATUS_LABEL = {
  CREATED: '已创建', QUEUED: '等待中', RUNNING: '进行中',
  SUCCESS: '已完成', FAILED: '失败',
}
// 后端状态 → CSS class
const STATUS_CLASS = {
  CREATED: 'created', QUEUED: 'waiting', RUNNING: 'progress',
  SUCCESS: 'done',    FAILED: 'failed',
}

function mapTask(t) {
  return {
    id:      t.id,
    name:    t.name,
    type:    t.type,
    status:  STATUS_LABEL[t.status] || t.status,
    statusClass: STATUS_CLASS[t.status] || 'waiting',
    time:    t.created_at,
  }
}

async function loadTasks() {
  loadingTasks.value = true
  try {
    const res = await api.listTasks()
    tasks.value = (res.data.tasks || []).map(mapTask)
  } catch {
    tasks.value = []
  } finally {
    loadingTasks.value = false
  }
}

// ============================================================
//  统计卡片
// ============================================================
const stats = computed(() => [
  {
    key: 'all', icon: '📊', label: '全部任务',
    value: tasks.value.length,
    tasks: tasks.value,
  },
  {
    key: 'waiting', icon: '⏳', label: '等待中',
    value: tasks.value.filter(t => t.statusClass === 'waiting').length,
    tasks: tasks.value.filter(t => t.statusClass === 'waiting'),
  },
  {
    key: 'progress', icon: '🔄', label: '进行中',
    value: tasks.value.filter(t => t.statusClass === 'progress').length,
    tasks: tasks.value.filter(t => t.statusClass === 'progress'),
  },
  {
    key: 'done', icon: '✅', label: '已完成',
    value: tasks.value.filter(t => t.statusClass === 'done').length,
    tasks: tasks.value.filter(t => t.statusClass === 'done'),
  },
])

// ============================================================
//  平台功能卡片
// ============================================================
const cards = ref([
  { id: 1,  title: '图片压缩',   icon: '🖼️', description: '支持 JPG / PNG / WebP 格式压缩，自定义压缩比例与输出质量', variant: 'progress', badge: '热门' },
  { id: 2,  title: '进制转换',   icon: '🔢', description: '二进制、八进制、十进制、十六进制互转，支持批量输入',       variant: 'todo',     badge: '' },
  { id: 3,  title: '编码转换',   icon: '🔤', description: 'URL 编码、Base64 编解码、Unicode 转义一键完成',           variant: 'progress', badge: '' },
  { id: 4,  title: 'JSON 格式化', icon: '📋', description: 'JSON 美化、压缩、校验，支持树形折叠查看',                 variant: 'todo',     badge: '新' },
  { id: 5,  title: '文本对比',   icon: '📝', description: '逐行对比两段文本差异，高亮显示新增/删除/修改内容',           variant: 'todo',     badge: '' },
  { id: 6,  title: '二维码生成', icon: '📱', description: '输入文本或链接生成二维码，支持调整尺寸并下载 PNG',         variant: 'done',     badge: '' },
  { id: 7,  title: '时间戳转换', icon: '🕐', description: 'Unix 时间戳与日期字符串互转，支持多时区',                 variant: 'todo',     badge: '' },
  { id: 8,  title: '颜色转换',   icon: '🎨', description: 'Hex / RGB / HSL 颜色格式互转，实时预览色值',             variant: 'archive',  badge: 'Beta' },
])

// ============================================================
//  弹窗 —— 统计卡片
// ============================================================
const showStatModal = ref(false)
const statModal = reactive({ key: '', icon: '', label: '', tasks: [] })

function openStatModal(s) {
  statModal.key = s.key
  statModal.icon = s.icon
  statModal.label = s.label
  statModal.tasks = s.tasks
  showStatModal.value = true
}

function closeStatModal() {
  showStatModal.value = false
}

// ============================================================
//  弹窗 —— 功能卡片
// ============================================================
const showCardModal = ref(false)
const activeCard = ref(null)

function openCardModal(item) {
  activeCard.value = item
  showCardModal.value = true
}

function closeCardModal() {
  showCardModal.value = false
  activeCard.value = null
}

// ============================================================
//  临时：随机创建任务
// ============================================================
const creatingTask = ref(false)
const toast = reactive({ show: false, type: 'success', msg: '' })

const TASK_POOL = [
  { name: '产品宣传图压缩',        type: '图片压缩' },
  { name: 'API 参数 Base64 编码',   type: '编码转换' },
  { name: '数据库连接串转义',      type: '编码转换' },
  { name: '日志分析进制计算',      type: '进制转换' },
  { name: '配置文件 JSON 校验',    type: 'JSON格式化' },
  { name: '接口响应体格式化',      type: 'JSON格式化' },
  { name: '新旧配置文本对比',      type: '文本对比' },
  { name: '微信小程序码生成',      type: '二维码生成' },
  { name: 'App 启动页图片压缩',    type: '图片压缩' },
  { name: 'Hex to ASCII 转换',     type: '进制转换' },
  { name: 'URL 参数编码处理',      type: '编码转换' },
  { name: '支付回调 JSON 美化',    type: 'JSON格式化' },
  { name: '用户头像批量压缩',      type: '图片压缩' },
  { name: '时间戳调试工具',        type: '时间戳转换' },
  { name: '品牌色值转换',          type: '颜色转换' },
]

function randomPick(arr) {
  return arr[Math.floor(Math.random() * arr.length)]
}

async function createRandomTask() {
  creatingTask.value = true
  const task = randomPick(TASK_POOL)
  try {
    await api.createTask({ name: task.name, type: task.type })
    loadTasks()  // 刷新任务总览
    toast.msg = `已创建：${task.name}`
    toast.type = 'success'
    toast.show = true
    setTimeout(() => { toast.show = false }, 2000)
  } catch (e) {
    toast.msg = e instanceof ApiError ? e.message : '创建失败'
    toast.type = 'error'
    toast.show = true
    setTimeout(() => { toast.show = false }, 2500)
  } finally {
    creatingTask.value = false
  }
}

// ============================================================
//  生命周期
// ============================================================
function onClickOutside(e) {
  if (!e.target.closest('.topbar__user')) {
    showUserMenu.value = false
  }
}

onMounted(() => {
  loadUser()
  loadTasks()
  document.addEventListener('click', onClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', onClickOutside)
})
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

.topbar__left  { display: flex; align-items: center; gap: 8px; }
.topbar__logo  { font-size: 24px; }
.topbar__brand { font-size: 18px; font-weight: 700; letter-spacing: -0.3px; }
.topbar__right { display: flex; align-items: center; gap: 12px; }

.topbar__notify {
  position: relative;
  width: 36px; height: 36px;
  border: none; border-radius: 10px;
  background: transparent;
  color: #555;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background .2s;
}
.topbar__notify:hover { background: #f0f2f5; }

.topbar__notify-dot {
  position: absolute;
  top: 7px; right: 8px;
  width: 7px; height: 7px;
  border-radius: 50%;
  background: #f56c6c;
  border: 2px solid #fff;
}

.topbar__user {
  display: flex; align-items: center; gap: 10px;
  padding: 4px 8px 4px 4px;
  border-radius: 28px;
  cursor: pointer;
  transition: background .2s;
  position: relative;
}
.topbar__user:hover { background: #f0f2f5; }

.topbar__avatar {
  width: 34px; height: 34px;
  border-radius: 50%;
  object-fit: cover;
  background: linear-gradient(135deg, #4f6ef7, #7c5cfc);
  flex-shrink: 0;
}

.topbar__avatar-fallback {
  width: 34px; height: 34px;
  border-radius: 50%;
  background: linear-gradient(135deg, #4f6ef7, #7c5cfc);
  color: #fff; font-size: 15px; font-weight: 600;
  display: flex; align-items: center; justify-content: center;
  flex-shrink: 0;
}

.topbar__user-meta { display: flex; flex-direction: column; line-height: 1.3; }
.topbar__username   { font-size: 14px; font-weight: 600; }
.topbar__role       { font-size: 11px; color: var(--color-secondary); }
.topbar__chevron    { color: #aaa; margin-left: 2px; transition: transform .2s; }

.topbar__dropdown {
  position: absolute; top: 100%; right: 0;
  margin-top: 8px; min-width: 120px;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 8px 30px rgba(0,0,0,.12);
  overflow: hidden;
  z-index: 200;
}

.topbar__dropdown-item {
  padding: 10px 16px;
  font-size: 13px; color: #606266;
  cursor: pointer;
  transition: background .15s, color .15s;
}
.topbar__dropdown-item:hover { background: #f0f2f5; color: #f56c6c; }

.menu-enter-active, .menu-leave-active { transition: opacity .15s, transform .15s; }
.menu-enter-from,   .menu-leave-to      { opacity: 0; transform: translateY(-6px); }

/* ============================================================
   Hero 横条
   ============================================================ */
.hero {
  background: linear-gradient(135deg, #2c3e50 0%, #3b5998 50%, #4a6fa5 100%);
  color: #fff;
}

.hero__inner {
  max-width: 1200px; margin: 0 auto;
  padding: 28px 32px;
  display: flex; align-items: flex-end; justify-content: space-between;
  gap: 16px;
}

.hero__greeting { margin: 0; font-size: 24px; font-weight: 700; letter-spacing: -0.3px; }
.hero__subtitle { margin: 4px 0 0; font-size: 14px; opacity: .78; }
.hero__date     { font-size: 13px; opacity: .7; white-space: nowrap; }

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
  display: flex; align-items: center; gap: 14px;
  box-shadow: var(--shadow-md);
  cursor: pointer;
  transition: transform .2s, box-shadow .2s, outline-color .2s;
  outline: 2px solid transparent;
  outline-offset: -2px;
}

.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 28px rgba(0,0,0,.1);
}

.stat-card--active {
  outline-color: #4f6ef7;
}

.stat-card__icon  { font-size: 30px; line-height: 1; }
.stat-card__body  { display: flex; flex-direction: column; }
.stat-card__value { font-size: 24px; font-weight: 700; line-height: 1.2; }
.stat-card__label { font-size: 13px; color: var(--color-secondary); margin-top: 2px; }

/* ============================================================
   平台功能面板
   ============================================================ */
.section {
  max-width: 1200px;
  margin: 32px auto 0;
  padding: 0 32px 40px;
}

.section__header {
  display: flex; align-items: baseline; justify-content: space-between;
  margin-bottom: 20px;
}

.section__title { margin: 0; font-size: 19px; font-weight: 700; }
.section__count { font-size: 13px; color: #909399; }

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
  position: fixed; inset: 0;
  background: rgba(0,0,0,.45);
  display: flex; align-items: center; justify-content: center;
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

.modal--wide {
  width: min(720px, 92%);
  text-align: left;
}

.modal--wide .modal__title {
  text-align: center;
}

.modal__close {
  position: absolute; top: 14px; right: 16px;
  width: 32px; height: 32px;
  border: none; border-radius: 50%;
  background: #f2f3f5;
  color: #606266; font-size: 16px;
  cursor: pointer;
  display: flex; align-items: center; justify-content: center;
  transition: background .2s, color .2s;
}
.modal__close:hover { background: #e0e1e5; color: #303133; }

.modal__icon  { font-size: 48px; margin-bottom: 12px; }
.modal__title { margin: 0 0 10px; font-size: 22px; font-weight: 700; }
.modal__desc  { margin: 0; color: var(--color-secondary); font-size: 15px; line-height: 1.7; }
.modal__empty { text-align: center; color: #909399; padding: 24px 0; font-size: 14px; }

.modal__actions {
  display: flex; gap: 12px; justify-content: center;
  margin-top: 28px;
}

.modal__btn {
  padding: 10px 28px;
  border: none; border-radius: 10px;
  font-size: 15px; font-weight: 500;
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

/* ============================================================
   任务表格
   ============================================================ */
.task-table {
  width: 100%;
  margin-top: 20px;
  border-collapse: collapse;
  font-size: 14px;
}

.task-table th {
  text-align: left;
  padding: 10px 14px;
  background: #f7f8fa;
  color: #606266;
  font-weight: 600;
  font-size: 13px;
  border-bottom: 2px solid #e8eaef;
}

.task-table td {
  padding: 12px 14px;
  border-bottom: 1px solid #f0f2f5;
  color: #303133;
}

.task-table tbody tr:hover {
  background: #f9fafc;
}

.task-table__name {
  font-weight: 500;
}

.task-table__time {
  color: #909399;
  font-size: 13px;
  white-space: nowrap;
}

/* 状态标签 */
.task-table__status {
  display: inline-block;
  padding: 3px 10px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.task-table__status--waiting  { background: #fdf6ec; color: #e6a23c; }
.task-table__status--progress { background: #ecf5ff; color: #409eff; }
.task-table__status--done     { background: #f0f9eb; color: #67c23a; }
.task-table__status--created  { background: #f4f4f5; color: #909399; }
.task-table__status--failed   { background: #fef0f0; color: #f56c6c; }

/* ============================================================
   弹窗动画
   ============================================================ */
.modal-enter-active, .modal-leave-active { transition: opacity .25s; }
.modal-enter-active .modal, .modal-leave-active .modal { transition: transform .25s, opacity .25s; }
.modal-enter-from, .modal-leave-to { opacity: 0; }
.modal-enter-from .modal { transform: scale(.92) translateY(20px); opacity: 0; }
.modal-leave-to   .modal { transform: scale(.92) translateY(20px); opacity: 0; }

/* ============================================================
   浮动按钮（临时）
   ============================================================ */
.fab {
  position: fixed;
  bottom: 32px;
  right: 32px;
  width: 52px;
  height: 52px;
  border: none;
  border-radius: 16px;
  background: linear-gradient(135deg, #4f6ef7, #3b5998);
  color: #fff;
  font-size: 24px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 6px 24px rgba(79,110,247,.35);
  transition: transform .2s, box-shadow .2s, opacity .2s;
  z-index: 500;
}

.fab:hover:not(:disabled) {
  transform: translateY(-2px) scale(1.04);
  box-shadow: 0 8px 32px rgba(79,110,247,.5);
}

.fab:active:not(:disabled) { transform: scale(.96); }
.fab:disabled { opacity: .7; cursor: not-allowed; }

.fab__spinner {
  width: 20px; height: 20px;
  border: 2px solid rgba(255,255,255,.3);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin .7s linear infinite;
}

@keyframes spin { to { transform: rotate(360deg); } }

/* Toast 通知 */
.toast {
  position: fixed;
  bottom: 36px;
  left: 50%;
  transform: translateX(-50%);
  padding: 10px 24px;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 500;
  z-index: 600;
  pointer-events: none;
}

.toast--success { background: #f0f9eb; color: #67c23a; border: 1px solid #e1f3d8; }
.toast--error   { background: #fef0f0; color: #f56c6c; border: 1px solid #fde2e2; }

.toast-enter-active,
.toast-leave-active { transition: opacity .25s, transform .25s; }
.toast-enter-from,
.toast-leave-to       { opacity: 0; transform: translateX(-50%) translateY(8px); }

/* ============================================================
   响应式
   ============================================================ */
@media (max-width: 768px) {
  .topbar__inner { padding: 0 20px; }
  .topbar__user-meta, .topbar__chevron { display: none; }

  .stats { grid-template-columns: repeat(2, 1fr); }

  .hero__inner {
    flex-direction: column; align-items: flex-start;
    padding: 20px;
  }
  .hero__greeting { font-size: 20px; }

  .section, .stats { padding-left: 20px; padding-right: 20px; }

  .card-grid { grid-template-columns: 1fr; }

  .task-table { font-size: 12px; }
  .task-table th, .task-table td { padding: 8px 10px; }
}

@media (max-width: 480px) {
  .stats { grid-template-columns: 1fr; }
}
</style>

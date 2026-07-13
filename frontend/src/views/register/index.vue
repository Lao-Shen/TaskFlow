<template>
  <div class="register">
    <!-- 左侧装饰区 -->
    <div class="register__hero">
      <div class="register__hero-inner">
        <h1 class="register__brand">
          <span class="register__logo">⚡</span>TaskFlow
        </h1>
        <p class="register__tagline">加入我们，开始高效协作<br />注册一个账号即可免费使用</p>
        <div class="register__deco">
          <div class="register__deco-shape"></div>
          <div class="register__deco-shape"></div>
          <div class="register__deco-shape"></div>
        </div>
      </div>
    </div>

    <!-- 右侧注册卡片 -->
    <div class="register__main">
      <div class="register__card">
        <h2 class="register__title">创建账号</h2>
        <p class="register__subtitle">填写以下信息完成注册</p>

        <form class="register__form" @submit.prevent="handleRegister">
          <!-- 错误提示 -->
          <div v-if="errorMsg" class="form-error">{{ errorMsg }}</div>

          <!-- 用户名 -->
          <div class="field">
            <label class="field__label" for="username">用户名</label>
            <div class="field__wrap">
              <svg class="field__icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                <circle cx="12" cy="7" r="4"/>
              </svg>
              <input
                id="username"
                v-model="form.username"
                class="field__input"
                type="text"
                placeholder="请输入用户名"
                autocomplete="username"
              />
            </div>
          </div>

          <!-- 邮箱 -->
          <div class="field">
            <label class="field__label" for="email">邮箱</label>
            <div class="field__wrap">
              <svg class="field__icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="2" y="4" width="20" height="16" rx="2"/>
                <path d="m22 4-10 8L2 4"/>
              </svg>
              <input
                id="email"
                v-model="form.email"
                class="field__input"
                type="email"
                placeholder="请输入邮箱地址"
                autocomplete="email"
              />
            </div>
          </div>

          <!-- 密码 -->
          <div class="field">
            <label class="field__label" for="password">密码</label>
            <div class="field__wrap">
              <svg class="field__icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
                <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
              </svg>
              <input
                id="password"
                v-model="form.password"
                class="field__input"
                :type="showPwd ? 'text' : 'password'"
                placeholder="请输入密码"
                autocomplete="new-password"
              />
              <button type="button" class="field__toggle" @click="showPwd = !showPwd" tabindex="-1">
                <svg v-if="!showPwd" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94"/>
                  <path d="M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19"/>
                  <path d="M14.12 14.12a3 3 0 1 1-4.24-4.24"/>
                  <line x1="1" y1="1" x2="23" y2="23"/>
                </svg>
                <svg v-else width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/>
                  <circle cx="12" cy="12" r="3"/>
                </svg>
              </button>
            </div>
          </div>

          <!-- 确认密码 -->
          <div class="field">
            <label class="field__label" for="confirmPwd">确认密码</label>
            <div class="field__wrap">
              <svg class="field__icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
                <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
                <path d="m9 16 2 2 4-4"/>
              </svg>
              <input
                id="confirmPwd"
                v-model="form.confirmPwd"
                class="field__input"
                :class="{ 'field__input--error': pwdMismatch }"
                type="password"
                placeholder="请再次输入密码"
                autocomplete="new-password"
              />
              <span v-if="pwdMismatch" class="field__hint">两次密码不一致</span>
            </div>
          </div>

          <!-- 滑动验证 -->
          <div class="field">
            <label class="field__label">安全验证</label>
            <div
              class="slider"
              :class="{
                'slider--passed': slider.passed,
                'slider--failed': slider.status === 'failed',
                'slider--dragging': slider.dragging,
              }"
              @mousedown.prevent="onSliderStart"
              @touchstart.prevent="onSliderStart"
            >
              <span class="slider__hint">
                <template v-if="slider.passed">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"/></svg>
                  验证通过
                </template>
                <template v-else>
                  向右滑动完成验证
                </template>
              </span>

              <div
                ref="sliderBtn"
                class="slider__btn"
                :style="{ left: slider.offset + 'px' }"
                @mousedown.prevent="onSliderStart"
                @touchstart.prevent="onSliderStart"
              >
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                  <polyline points="9 18 15 12 9 6"/>
                </svg>
              </div>

              <div class="slider__track" :style="{ width: slider.trackWidth + 'px' }"></div>
            </div>
          </div>

          <!-- 注册按钮 -->
          <button
            class="register__btn"
            type="submit"
            :disabled="!canSubmit || loading"
            :class="{ 'register__btn--loading': loading }"
          >
            <span v-if="loading" class="register__btn-spinner"></span>
            <span>{{ loading ? '注册中…' : '注 册' }}</span>
          </button>
        </form>

        <p class="register__footer">
          已有账号？<router-link to="/login" class="register__link">立即登录</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import api, { setToken, ApiError } from '@/api'

// ============================================================
//  表单数据
// ============================================================
const form = reactive({
  username: '',
  email: '',
  password: '',
  confirmPwd: '',
})

const showPwd = ref(false)
const loading = ref(false)
const errorMsg = ref('')

/** 两次密码是否不一致 */
const pwdMismatch = computed(() => {
  return form.confirmPwd.length > 0 && form.password !== form.confirmPwd
})

// ============================================================
//  滑动验证
// ============================================================
const sliderBtn = ref(null)

const slider = reactive({
  offset: 0,
  trackWidth: 0,
  passed: false,
  dragging: false,
  status: '',
  startX: 0,
  maxOffset: 0,
})

/** 校验是否可提交 */
const canSubmit = computed(() => {
  return form.username.trim() !== ''
    && form.email.trim() !== ''
    && form.password.trim() !== ''
    && form.confirmPwd.trim() !== ''
    && !pwdMismatch.value
    && slider.passed
    && !loading.value
})

// ---------- 滑块拖拽逻辑 ----------
function onSliderStart(e) {
  if (slider.passed || loading.value) return

  const clientX = e.touches ? e.touches[0].clientX : e.clientX
  slider.dragging = true
  slider.startX = clientX
  slider.status = ''

  const wrap = e.currentTarget
  const btnW = sliderBtn.value ? sliderBtn.value.offsetWidth : 40
  slider.maxOffset = wrap.offsetWidth - btnW - 4

  window.addEventListener('mousemove', onMove)
  window.addEventListener('mouseup', onEnd)
  window.addEventListener('touchmove', onMove, { passive: false })
  window.addEventListener('touchend', onEnd)
}

function onMove(e) {
  if (!slider.dragging) return
  const clientX = e.touches ? e.touches[0].clientX : e.clientX
  const dx = clientX - slider.startX
  slider.offset = Math.max(0, Math.min(dx, slider.maxOffset))
  slider.trackWidth = Math.max(0, slider.offset)
}

function onEnd() {
  if (!slider.dragging) return
  slider.dragging = false

  if (slider.offset >= slider.maxOffset - 4) {
    slider.offset = slider.maxOffset
    slider.trackWidth = slider.maxOffset
    slider.passed = true
    slider.status = ''
  } else {
    slider.offset = 0
    slider.trackWidth = 0
    slider.status = 'failed'
    setTimeout(() => {
      if (!slider.passed) slider.status = ''
    }, 600)
  }

  window.removeEventListener('mousemove', onMove)
  window.removeEventListener('mouseup', onEnd)
  window.removeEventListener('touchmove', onMove)
  window.removeEventListener('touchend', onEnd)
}

// ============================================================
//  注册
// ============================================================
const router = useRouter()

async function handleRegister() {
  if (!canSubmit.value) return

  errorMsg.value = ''
  loading.value = true
  try {
    const res = await api.register({
      username: form.username,
      email: form.email,
      password: form.password,
    })
    setToken(res.data.token)
    router.push('/dashboard')
  } catch (e) {
    errorMsg.value = e instanceof ApiError ? e.message : '网络错误，请稍后重试'
    // 重置滑块
    slider.passed = false
    slider.offset = 0
    slider.trackWidth = 0
    slider.status = 'failed'
    setTimeout(() => { if (!slider.passed) slider.status = '' }, 600)
  } finally {
    loading.value = false
  }
}

// ---------- 清理全局事件 ----------
onUnmounted(() => {
  window.removeEventListener('mousemove', onMove)
  window.removeEventListener('mouseup', onEnd)
  window.removeEventListener('touchmove', onMove)
  window.removeEventListener('touchend', onEnd)
})
</script>

<style scoped>
/* ============================================================
   整体布局
   ============================================================ */
.register {
  min-height: 100vh;
  display: flex;
  background: linear-gradient(160deg, #f0f3f8 0%, #f7f9fc 40%, #f5f6fa 100%);
  font-family: Inter, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}

/* ============================================================
   左侧装饰区
   ============================================================ */
.register__hero {
  flex: 1;
  display: none;
  background: linear-gradient(160deg, #2c3e50 0%, #3b5998 45%, #4a6fa5 100%);
  color: #fff;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow: hidden;
}

.register__hero-inner {
  position: relative;
  z-index: 1;
  text-align: center;
  padding: 48px;
}

.register__brand {
  margin: 0;
  font-size: 38px;
  font-weight: 700;
  letter-spacing: -0.5px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
}

.register__logo { font-size: 42px; }

.register__tagline {
  margin: 16px 0 0;
  font-size: 16px;
  line-height: 1.8;
  opacity: .78;
}

.register__deco {
  margin-top: 48px;
  display: flex;
  gap: 16px;
  justify-content: center;
}

.register__deco-shape {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: rgba(255,255,255,.25);
  animation: decoPulse 2.4s ease-in-out infinite;
}

.register__deco-shape:nth-child(2) { animation-delay: .3s; }
.register__deco-shape:nth-child(3) { animation-delay: .6s; }

@keyframes decoPulse {
  0%, 100% { transform: scale(1); opacity: .25; }
  50%      { transform: scale(1.6); opacity: .55; }
}

/* ============================================================
   右侧注册卡片
   ============================================================ */
.register__main {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
}

.register__card {
  width: 420px;
  max-width: 100%;
  background: #fff;
  border-radius: 20px;
  padding: 36px 36px 32px;
  box-shadow: 0 8px 40px rgba(0,0,0,.06);
}

.register__title {
  margin: 0;
  font-size: 26px;
  font-weight: 700;
  color: #1d1d1f;
  letter-spacing: -0.3px;
}

.register__subtitle {
  margin: 6px 0 0;
  font-size: 14px;
  color: #86868b;
}

/* ============================================================
   表单
   ============================================================ */
.register__form {
  margin-top: 28px;
  display: flex;
  flex-direction: column;
  gap: 18px;
}

/* 字段 */
.field { display: flex; flex-direction: column; gap: 6px; }

.field__label {
  font-size: 13px;
  font-weight: 600;
  color: #1d1d1f;
}

.field__wrap {
  position: relative;
  display: flex;
  align-items: center;
}

.field__icon {
  position: absolute;
  left: 14px;
  color: #aaa;
  pointer-events: none;
}

.field__input {
  width: 100%;
  padding: 12px 14px 12px 42px;
  border: 1.5px solid #e4e7ed;
  border-radius: 10px;
  font-size: 14px;
  color: #1d1d1f;
  background: #fafbfc;
  transition: border-color .2s, box-shadow .2s;
  outline: none;
}

.field__input:focus {
  border-color: #4f6ef7;
  box-shadow: 0 0 0 3px rgba(79,110,247,.1);
  background: #fff;
}

.field__input::placeholder { color: #c0c4cc; }

.field__input--error {
  border-color: #f56c6c;
}

.field__input--error:focus {
  border-color: #f56c6c;
  box-shadow: 0 0 0 3px rgba(245,108,108,.1);
}

.field__hint {
  position: absolute;
  right: 12px;
  font-size: 12px;
  color: #f56c6c;
  pointer-events: none;
}

/* 密码可见切换 */
.field__toggle {
  position: absolute;
  right: 10px;
  width: 32px;
  height: 32px;
  border: none;
  border-radius: 6px;
  background: transparent;
  color: #aaa;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: color .2s, background .2s;
}

.field__toggle:hover { color: #606266; background: #f0f2f5; }

/* 错误提示 */
.form-error {
  padding: 10px 14px;
  border-radius: 8px;
  background: #fef0f0;
  color: #f56c6c;
  font-size: 13px;
  border: 1px solid #fde2e2;
}

/* ============================================================
   滑动验证
   ============================================================ */
.slider {
  position: relative;
  height: 42px;
  border-radius: 10px;
  background: #f2f4f7;
  display: flex;
  align-items: center;
  justify-content: center;
  user-select: none;
  overflow: hidden;
  transition: background .3s;
}

.slider--passed  { background: #e8f5e9; }
.slider--failed  { background: #fef0f0; }

.slider__track {
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  border-radius: 10px 0 0 10px;
  background: linear-gradient(135deg, #67c23a, #85ce61);
  transition: width .06s linear;
  pointer-events: none;
}

.slider--passed .slider__track { border-radius: 10px; }

.slider__hint {
  position: relative;
  z-index: 1;
  font-size: 13px;
  color: #909399;
  display: flex;
  align-items: center;
  gap: 6px;
  pointer-events: none;
  transition: color .3s;
}

.slider--passed .slider__hint { color: #67c23a; }
.slider--failed .slider__hint { color: #f56c6c; }

.slider__btn {
  position: absolute;
  left: 2px;
  top: 2px;
  width: 38px;
  height: 38px;
  border-radius: 9px;
  background: #fff;
  box-shadow: 0 2px 8px rgba(0,0,0,.12);
  display: flex;
  align-items: center;
  justify-content: center;
  color: #606266;
  cursor: grab;
  z-index: 2;
  transition: box-shadow .2s;
}

.slider__btn:active { cursor: grabbing; }

.slider--dragging .slider__btn { box-shadow: 0 4px 16px rgba(0,0,0,.16); }

.slider--passed .slider__btn {
  color: #67c23a;
  cursor: default;
}

.slider--failed .slider__btn {
  animation: sliderShake .4s ease;
}

@keyframes sliderShake {
  0%, 100%   { transform: translateX(0); }
  20%        { transform: translateX(-6px); }
  40%        { transform: translateX(6px); }
  60%        { transform: translateX(-4px); }
  80%        { transform: translateX(4px); }
}

/* ============================================================
   注册按钮
   ============================================================ */
.register__btn {
  width: 100%;
  padding: 13px 0;
  margin-top: 4px;
  border: none;
  border-radius: 10px;
  background: linear-gradient(135deg, #4f6ef7, #3b5998);
  color: #fff;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  transition: opacity .2s, transform .15s, box-shadow .2s;
  box-shadow: 0 4px 16px rgba(79,110,247,.3);
}

.register__btn:hover:not(:disabled) { opacity: .92; box-shadow: 0 6px 22px rgba(79,110,247,.4); }
.register__btn:active:not(:disabled) { transform: scale(.98); }
.register__btn:disabled { opacity: .5; cursor: not-allowed; box-shadow: none; }

.register__btn-spinner {
  width: 16px;
  height: 16px;
  border: 2px solid rgba(255,255,255,.3);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin .7s linear infinite;
}

@keyframes spin { to { transform: rotate(360deg); } }

/* ============================================================
   底部链接
   ============================================================ */
.register__footer {
  margin: 24px 0 0;
  text-align: center;
  font-size: 13px;
  color: #909399;
}

.register__link {
  color: #4f6ef7;
  text-decoration: none;
  font-weight: 500;
}

.register__link:hover { text-decoration: underline; }

/* ============================================================
   响应式
   ============================================================ */
@media (min-width: 768px) {
  .register__hero {
    display: flex;
  }
}

@media (max-width: 767px) {
  .register__card {
    padding: 28px 24px 24px;
  }
}
</style>
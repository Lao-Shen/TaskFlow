<template>
  <div class="task-card" :class="[`task-card--${variant}`]">
    <div class="task-card__accent"></div>
    <div class="task-card__body">
      <header class="task-card__header">
        <span class="task-card__icon">{{ icon }}</span>
        <h3 class="task-card__title">{{ title }}</h3>
        <span v-if="badge" class="task-card__badge">{{ badge }}</span>
      </header>
      <div class="task-card__content">
        <p class="task-card__description">{{ description }}</p>
        <slot />
      </div>
      <footer class="task-card__footer">
        <span class="task-card__arrow">→</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps({
  title: { type: String, default: '任务名称' },
  description: { type: String, default: '任务详情描述' },
  icon: { type: String, default: '📋' },
  badge: { type: String, default: '' },
  variant: { type: String, default: 'default' },
})
</script>

<style scoped>
.task-card {
  position: relative;
  display: flex;
  border-radius: 12px;
  background: #fff;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.05);
  overflow: hidden;
  transition: transform 0.25s ease, box-shadow 0.25s ease;
}

.task-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.1);
}

.task-card__accent {
  width: 4px;
  flex-shrink: 0;
  background: #e0e0e0;
  border-radius: 4px 0 0 4px;
  transition: background 0.25s ease;
}

.task-card--todo .task-card__accent { background: #e6a23c; }
.task-card--progress .task-card__accent { background: #409eff; }
.task-card--done .task-card__accent { background: #67c23a; }
.task-card--archive .task-card__accent { background: #909399; }

.task-card__body {
  flex: 1;
  padding: 20px 20px 16px;
  display: flex;
  flex-direction: column;
}

.task-card__header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 10px;
}

.task-card__icon {
  font-size: 22px;
  line-height: 1;
}

.task-card__title {
  margin: 0;
  font-size: 17px;
  font-weight: 600;
  color: #1d1d1f;
  flex: 1;
}

.task-card__badge {
  font-size: 12px;
  padding: 2px 10px;
  border-radius: 20px;
  background: #f0f2f5;
  color: #606266;
  font-weight: 500;
}

.task-card--todo .task-card__badge { background: #fdf6ec; color: #e6a23c; }
.task-card--progress .task-card__badge { background: #ecf5ff; color: #409eff; }
.task-card--done .task-card__badge { background: #f0f9eb; color: #67c23a; }
.task-card--archive .task-card__badge { background: #f4f4f5; color: #909399; }

.task-card__content {
  color: #86868b;
  line-height: 1.6;
  font-size: 14px;
  flex: 1;
}

.task-card__description {
  margin: 0;
}

.task-card__footer {
  margin-top: 14px;
  display: flex;
  justify-content: flex-end;
}

.task-card__arrow {
  font-size: 16px;
  color: #c0c4cc;
  transition: transform 0.25s ease, color 0.25s ease;
}

.task-card:hover .task-card__arrow {
  transform: translateX(4px);
  color: #409eff;
}
</style>

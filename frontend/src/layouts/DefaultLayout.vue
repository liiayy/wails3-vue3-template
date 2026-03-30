<script setup lang="ts">
import { computed, ref, onMounted, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useSettingsStore } from '@/stores/settings'
import { useI18n } from 'vue-i18n'
import { useWindowControl } from '@/composables'
import { RectangleIcon, MinusIcon, CloseIcon } from 'tdesign-icons-vue-next'
import { Window } from '@wailsio/runtime'

const router = useRouter()
const route = useRoute()
const settings = useSettingsStore()
const { t } = useI18n()
const { minimise } = useWindowControl() // Only minimise is needed from useWindowControl

// Custom window control functions
const close = () => {
  // 按照需求：点击关闭按钮仅隐藏窗口到托盘，不退出程序
  Window.Hide()
}
const toggleMaximise = () => {
  Window.ToggleMaximise()
}

// 当前激活的菜单项
const activeMenu = computed(() => {
  const name = route.name as string
  if (name?.startsWith('demo')) return 'demo'
  if (name?.startsWith('settings')) return 'settings'
  return name || 'home'
})

// 动态通过路由配置生成菜单
const allRoutes = router.options.routes.find((r) => r.path === '/')?.children || []

// 顶部主菜单
const topMenuItems = computed(() => {
  return allRoutes
    .filter((r) => r.meta?.showInMenu && r.meta?.menuSection === 'top')
    .map((r) => ({
      value: r.name as string,
      label: t(r.meta?.title as string),
      icon: r.meta?.icon, // 直接透传组件
      path: r.path === '' ? '/' : `/${r.path}`,
    }))
})

// 底部功能菜单
const bottomMenuItems = computed(() => {
  return allRoutes
    .filter((r) => r.meta?.showInMenu && r.meta?.menuSection === 'bottom')
    .map((r) => ({
      value: r.name as string,
      label: t(r.meta?.title as string),
      icon: r.meta?.icon, // 直接透传组件
      path: `/${r.path}`,
    }))
})

function onMenuChange(value: string) {
  const allItems = [...topMenuItems.value, ...bottomMenuItems.value]
  const item = allItems.find((m) => m.value === value)
  if (item) {
    router.push(item.path)
  }
}

function toggleSidebar() {
  settings.updateSetting('isSidebarCollapsed', !settings.isSidebarCollapsed)
}

// 底部状态栏时间
const currentTime = ref('')
let timer: number | null = null

const updateTime = () => {
  const now = new Date()
  const hours = String(now.getHours()).padStart(2, '0')
  const minutes = String(now.getMinutes()).padStart(2, '0')
  const seconds = String(now.getSeconds()).padStart(2, '0')
  currentTime.value = `${hours}:${minutes}:${seconds}`
}

onMounted(() => {
  updateTime()
  timer = window.setInterval(updateTime, 1000)
})

onUnmounted(() => {
  if (timer) clearInterval(timer)
})
</script>

<template>
  <div class="h-full w-full flex overflow-hidden bg-[var(--td-bg-color-page)]">
    <!-- ========== 侧边栏 ========== -->
    <aside
      class="sidebar flex flex-col h-full transition-all duration-300 bg-[var(--td-bg-color-container)] border-[var(--td-border-level-1-color)]"
      :style="{ width: settings.isSidebarCollapsed ? '55px' : '200px' }"
    >
      <!-- Logo 区域 (可拖拽) -->
      <div
        class="h-[50px] flex items-center gap-2 px-4 shrink-0"
        style="--wails-draggable: drag; -webkit-app-region: drag; user-select: none"
      >
        <div
          class="w-7 h-7 rounded-md bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center text-white text-xs font-bold shrink-0 shadow"
        >
          M
        </div>
        <span
          v-show="!settings.isSidebarCollapsed"
          class="text-sm font-semibold text-[var(--td-text-color-primary)] whitespace-nowrap overflow-hidden"
        >
          MyApp2
        </span>
      </div>

      <!-- 菜单导航 -->
      <nav
        class="flex-1 flex flex-col overflow-x-hidden overflow-y-hidden"
        style="-webkit-app-region: no-drag"
      >
        <!-- 顶部主菜单 -->
        <div class="flex-1 overflow-y-auto overflow-x-hidden">
          <t-menu
            :value="activeMenu"
            :collapsed="settings.isSidebarCollapsed"
            @change="onMenuChange"
            style="width: 100%"
          >
            <t-menu-item v-for="item in topMenuItems" :key="item.value" :value="item.value">
              <template #icon>
                <component :is="item.icon" />
              </template>
              {{ item.label }}
            </t-menu-item>
          </t-menu>
        </div>

        <!-- 底部功能菜单 -->
        <div class="shrink-0 border-t border-[var(--td-border-level-1-color)]/50 overflow-x-hidden">
          <t-menu
            :value="activeMenu"
            :collapsed="settings.isSidebarCollapsed"
            @change="onMenuChange"
            style="width: 100%"
          >
            <t-menu-item v-for="item in bottomMenuItems" :key="item.value" :value="item.value">
              <template #icon>
                <component :is="item.icon" />
              </template>
              {{ item.label }}
            </t-menu-item>
          </t-menu>
        </div>
      </nav>
    </aside>

    <!-- ========== 右侧主区域 ========== -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <!-- 顶部标题栏 (可拖拽) -->
      <header
        class="h-[40px] flex items-center justify-between pl-5 shrink-0 border-[var(--td-border-level-1-color)] bg-[var(--td-bg-color-container)]"
        style="--wails-draggable: drag; -webkit-app-region: drag; user-select: none"
      >
        <div class="flex items-center gap-3">
          <div
            class="h-5 w-1 rounded-full bg-gradient-to-b from-blue-500 to-purple-600 shadow-sm"
          ></div>
          <h1 class="text-base font-semibold text-[var(--td-text-color-primary)] tracking-wide">
            {{ $t(`menu.${activeMenu}`) }}
          </h1>
        </div>

        <!-- 右侧窗口控制按钮 -->
        <div class="flex h-full items-stretch" style="-webkit-app-region: no-drag">
          <div class="window-control-btn" @click="minimise">
            <MinusIcon size="16" />
          </div>
          <div class="window-control-btn" @click="toggleMaximise">
            <RectangleIcon size="14" />
          </div>
          <div class="window-control-btn hover:bg-[#e81123] hover:text-white" @click="close">
            <CloseIcon size="16" />
          </div>
        </div>
      </header>

      <!-- 主内容区 -->
      <main class="flex-1 overflow-auto p-1">
        <router-view v-slot="{ Component }">
          <Transition name="fade" mode="out-in">
            <keep-alive>
              <component :is="Component" />
            </keep-alive>
          </Transition>
        </router-view>
      </main>

      <!-- ========== 底部状态栏 ========== -->
      <footer
        class="h-6 flex items-center justify-between px-3 shrink-0 border-t border-[var(--td-border-level-1-color)] bg-[var(--td-bg-color-container)] text-[11px] text-[var(--td-text-color-secondary)] uppercase tracking-wider"
      >
        <div class="flex items-center gap-3">
          <span class="flex items-center gap-1.5">
            <div
              class="w-1.5 h-1.5 rounded-full bg-green-500 shadow-[0_0_4px_rgba(34,197,94,0.6)]"
            ></div>
            Connected
          </span>
          <span class="opacity-40">|</span>
          <span>Ready</span>
        </div>

        <div class="flex-1 flex justify-center font-medium">MyApp2 v1.0.0</div>

        <div class="flex items-center gap-3">
          <span>UTF-8</span>
          <span class="opacity-40">|</span>
          <span class="font-mono tabular-nums">{{ currentTime }}</span>
        </div>
      </footer>
    </div>
  </div>
</template>

<style scoped>
/* 页面切换动画 */
.fade-enter-active,
.fade-leave-active {
  transition: all 0.11s cubic-bezier(0.16, 1, 0.3, 1);
}
.fade-enter-from {
  opacity: 0;
  transform: scale(0.99);
}
.fade-leave-to {
  opacity: 0;
  transform: scale(1);
}

/* 侧边栏菜单样式微调 */
.sidebar :deep(.t-default-menu) {
  border-right: none;
  background: transparent;
}

/* 窗口控制按钮样式 */
.window-control-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 46px;
  height: 100%;
  cursor: pointer;
  transition: all 0.2s;
  color: var(--td-text-color-primary);
}

.window-control-btn:hover {
  background-color: var(--td-bg-color-secondarycontainer);
}

.window-control-btn :deep(svg) {
  display: block;
  margin: auto;
}
:deep(.t-menu) {
  padding: var(--td-comp-paddingTB-s) var(--td-comp-paddingLR-s);
}
</style>

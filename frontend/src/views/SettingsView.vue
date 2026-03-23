<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import TitleBar from '../components/TitleBar.vue'
import { PaletteIcon, NotificationIcon } from 'tdesign-icons-vue-next'

const route = useRoute()
const router = useRouter()
const isStandalone = route.path.startsWith('/standalone')

// 基础路由路径 (处理独立窗口和主窗口场景)
const baseUrl = isStandalone ? '/standalone/settings' : '/settings'

// 根据当前路由名称判断激活分类
const activeCategory = computed(() => {
  // 匹配子路由名
  const routeName = route.name?.toString() || ''
  if (routeName.includes('notifications')) return 'notifications'
  return 'personalization'
})

const categories = [
  { id: 'personalization', label: 'settings.personalization', icon: PaletteIcon, path: 'personalization' },
  { id: 'notifications', label: 'settings.notifications', icon: NotificationIcon, path: 'notifications' },
]

function navigateTo(path: string) {
  router.push(`${baseUrl}/${path}`)
}
</script>

<template>
  <div class="h-full flex flex-col bg-[var(--td-bg-color-container)] overflow-hidden">
    <!-- 无边框窗口标题栏 -->
    <TitleBar v-if="isStandalone" :title="$t('settings.title')" no-minimize no-maximize />

    <div class="flex-1 flex overflow-hidden">
      <!-- 第二栏：分类导航 (Middle Column) -->
      <aside
        class="w-44 border-r border-[var(--td-border-level-1-color)] bg-[var(--td-bg-color-container)] shrink-0 flex flex-col"
      >
        <div class="p-2">
          <nav class="space-y-1">
            <div
              v-for="cat in categories"
              :key="cat.id"
              @click="navigateTo(cat.path)"
              class="flex items-center gap-2.5 px-3 py-2 rounded-lg cursor-pointer transition-all duration-200"
              :class="
                activeCategory === cat.id
                  ? 'bg-[var(--td-brand-color-light)] text-[var(--td-brand-color)] font-medium shadow-sm'
                  : 'text-[var(--td-text-color-secondary)] hover:bg-[var(--td-bg-color-secondarycontainer)]'
              "
            >
              <component :is="cat.icon" size="16" />
              <span class="text-sm">
                {{ $t(cat.label) }}
              </span>
            </div>
          </nav>
        </div>
      </aside>

      <!-- 第三栏：主内容区 (Content Area) -->
      <main class="flex-1 overflow-auto bg-[var(--td-bg-color-page)] relative">
        <div class="max-w-4xl p-6 space-y-6">
          <!-- 头部标题 -->
          <div class="space-y-1 animate-fade-in" :key="activeCategory">
            <h1 class="text-xl font-bold text-[var(--td-text-color-primary)] tracking-tight">
              {{ $t(categories.find((c) => c.id === activeCategory)?.label || '') }}
            </h1>
            <p class="text-[var(--td-text-color-secondary)] text-sm">
              {{ $t('settings.subtitle') }}
            </p>
          </div>

          <!-- 子路由展示区 -->
          <router-view v-slot="{ Component }">
            <transition name="fade-sub" mode="out-in">
              <keep-alive>
                <component :is="Component" />
              </keep-alive>
            </transition>
          </router-view>
        </div>
      </main>
    </div>
  </div>
</template>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 子路由切换动画 */
.fade-sub-enter-active,
.fade-sub-leave-active {
  transition: all 0.1s ease;
}
.fade-sub-enter-from {
  opacity: 0;
  transform: translateX(10px);
}
.fade-sub-leave-to {
  opacity: 0;
  transform: translateX(-10px);
}
</style>

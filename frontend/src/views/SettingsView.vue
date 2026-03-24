<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import TitleBar from '../components/TitleBar.vue'
import CategorySidebar from '../components/CategorySidebar.vue'

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
  { id: 'personalization', label: 'settings.personalization', icon: 'palette', path: 'personalization' },
  { id: 'notifications', label: 'settings.notifications', icon: 'notification', path: 'notifications' },
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
      <!-- 第二栏：分类导航 (Independent Component) -->
      <CategorySidebar
        :categories="categories"
        :active-value="activeCategory"
        @change="navigateTo"
      />

      <!-- 第三栏：主内容区 (Content Area) -->
      <main class="flex-1 overflow-auto bg-[var(--td-bg-color-page)] relative">
        <div class="max-w-4xl p-6 space-y-6">
          <!-- 头部标题 -->
          <div class="space-y-1">
            <h1 class="text-xl font-bold text-[var(--td-text-color-primary)] tracking-tight">
              {{ $t(categories.find((c) => c.id === activeCategory)?.label || '') }}
            </h1>
            <p class="text-[var(--td-text-color-secondary)] text-sm">
              {{ $t('settings.subtitle') }}
            </p>
          </div>

          <!-- 子路由展示区 -->
          <router-view v-slot="{ Component }">
            <keep-alive>
              <component :is="Component" />
            </keep-alive>
          </router-view>
        </div>
      </main>
    </div>
  </div>
</template>

<style scoped>
/* 侧边栏样式微调 */
</style>

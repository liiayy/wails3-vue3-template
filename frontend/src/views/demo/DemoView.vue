<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import CategorySidebar from '../../components/CategorySidebar.vue'

const route = useRoute()
const router = useRouter()
const baseUrl = '/demo'

// 根据当前路由名称判断激活分类
const activeCategory = computed(() => {
  const routeName = route.name?.toString() || ''
  if (routeName.includes('demo-file')) return 'file'
  if (routeName.includes('demo-clipboard')) return 'clipboard'
  if (routeName.includes('demo-dragdrop')) return 'dragdrop'
  if (routeName.includes('demo-notifications')) return 'notifications'
  return 'file' // 默认为文件演示
})

const categories = [
  { id: 'file', label: 'demo.fileOps', icon: 'file', path: 'file' },
  { id: 'clipboard', label: 'demo.clipboardOps', icon: 'copy', path: 'clipboard' },
  { id: 'dragdrop', label: 'demo.dragDropOps', icon: 'cloud-upload', path: 'dragdrop' },
  { id: 'notifications', label: 'demo.notificationOps', icon: 'notification', path: 'notifications' },
]

function navigateTo(path: string) {
  router.push(`${baseUrl}/${path}`)
}
</script>

<template>
  <div class="h-full flex flex-col bg-[var(--td-bg-color-container)] overflow-hidden">
    <div class="flex-1 flex overflow-hidden">
      <!-- 第二栏：功能导航 -->
      <CategorySidebar
        :categories="categories"
        :active-value="activeCategory"
        @change="navigateTo"
      />

      <!-- 第三栏：主演示区 -->
      <main class="flex-1 overflow-auto bg-[var(--td-bg-color-page)] relative">
        <div class="max-w-4xl p-6 space-y-6">
          <!-- 头部标题 -->
          <div class="space-y-1 animate-fade-in" :key="activeCategory">
            <h1 class="text-xl font-bold text-[var(--td-text-color-primary)] tracking-tight">
              {{ $t(categories.find((c) => c.id === activeCategory)?.label || '') }}
            </h1>
            <p class="text-[var(--td-text-color-secondary)] text-sm">
              {{ $t('demo.subtitle') }}
            </p>
          </div>

          <!-- 功能演示展示区 -->
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

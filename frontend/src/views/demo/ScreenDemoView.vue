<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { ScreenService } from '#/myapp2/internal/binding'
import { 
  DesktopIcon, 
  InfoCircleIcon, 
  CheckCircleFilledIcon,
  RefreshIcon
} from 'tdesign-icons-vue-next'
import { MessagePlugin } from 'tdesign-vue-next'

const screens = ref<any[]>([])
const loading = ref(false)
let timer: any = null

const fetchScreens = async () => {
  loading.value = true
  try {
    const data = await ScreenService.GetAllScreens()
    screens.value = data
  } catch (err) {
    MessagePlugin.error(`获取屏幕信息失败: ${err}`)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchScreens()
  // 每 3 秒自动刷新一次，模拟屏幕变化检测
  timer = setInterval(fetchScreens, 3000)
})

onUnmounted(() => {
  if (timer) clearInterval(timer)
})

// 计算所有屏幕的包围盒，用于图形化预览缩放
const boundingBox = computed(() => {
  if (screens.value.length === 0) return { minX: 0, minY: 0, maxX: 1, maxY: 1, width: 1, height: 1 }
  
  let minX = Infinity, minY = Infinity, maxX = -Infinity, maxY = -Infinity
  screens.value.forEach(s => {
    minX = Math.min(minX, s.x)
    minY = Math.min(minY, s.y)
    maxX = Math.max(maxX, s.x + s.width)
    maxY = Math.max(maxY, s.y + s.height)
  })
  
  return { minX, minY, maxX, maxY, width: maxX - minX, height: maxY - minY }
})

// 图形化比例计算
const getScreenStyle = (s: any) => {
  const containerWidth = 400 
  const containerHeight = 240
  const bb = boundingBox.value
  
  // 留出边距
  const padding = 20
  const availableWidth = containerWidth - padding * 2
  const availableHeight = containerHeight - padding * 2
  
  const scale = Math.min(availableWidth / bb.width, availableHeight / bb.height)
  
  return {
    left: `${(s.x - bb.minX) * scale + padding}px`,
    top: `${(s.y - bb.minY) * scale + padding}px`,
    width: `${s.width * scale}px`,
    height: `${s.height * scale}px`,
  }
}
</script>

<template>
  <div class="space-y-6">
    <!-- 说明 -->
    <div class="p-4 rounded-lg bg-[var(--td-brand-color-light)] text-[var(--td-brand-color)] text-sm leading-relaxed">
      {{ $t('demo.screenDesc') }}
      <p class="mt-2 text-xs opacity-70">支持实时监测：断开或接入显示器时会动态更新。Wails 3 自动处理 DPI 逻辑，您使用的都是逻辑像素。</p>
    </div>

    <!-- 图形化预览 -->
    <div class="relative h-[280px] rounded-2xl bg-[var(--td-bg-color-secondarycontainer)] border border-[var(--td-border-level-1-color)] overflow-hidden">
      <div class="absolute inset-0 opacity-10 pointer-events-none p-4 font-mono text-[100px] font-bold select-none overflow-hidden">
        MONITORS
      </div>
      
      <div 
        v-for="(s, idx) in screens" 
        :key="s.id"
        class="absolute border-2 rounded-lg transition-all duration-500 flex flex-col items-center justify-center p-2 shadow-lg group overflow-hidden"
        :class="[
          s.isPrimary ? 'border-[var(--td-brand-color)] bg-[var(--td-bg-color-container)]' : 'border-[var(--td-component-border)] bg-[var(--td-bg-color-secondarycontainer)]'
        ]"
        :style="getScreenStyle(s)"
      >
        <DesktopIcon :size="s.width < 1000 ? '24' : '48'" :class="s.isPrimary ? 'text-[var(--td-brand-color)]' : 'text-[var(--td-text-color-placeholder)]'" />
        <p class="mt-2 text-[10px] font-bold truncate max-w-full" :class="s.isPrimary ? 'text-[var(--td-brand-color)]' : 'text-[var(--td-text-color-secondary)]'">
           {{ s.name || `Screen ${idx+1}` }}
        </p>
        <span v-if="s.isPrimary" class="absolute top-1 right-1 text-[8px] px-1 bg-[var(--td-brand-color)] text-white rounded">PRIMARY</span>
        
        <!-- 悬浮详情 -->
        <div class="absolute inset-0 bg-black/80 text-white opacity-0 group-hover:opacity-100 transition-opacity duration-200 flex flex-col items-center justify-center text-[8px] space-y-1">
           <p>{{ s.width }}x{{ s.height }}</p>
           <p>DPI: {{ s.scaleFactor }}x</p>
           <p>Pos: ({{ s.x }}, {{ s.y }})</p>
        </div>
      </div>
    </div>

    <!-- 数据列表 -->
    <div class="space-y-4">
      <div class="flex items-center justify-between">
        <h3 class="text-sm font-medium text-[var(--td-text-color-primary)]">
          {{ $t('demo.screenList') }} ({{ screens.length }})
        </h3>
        <t-button theme="primary" variant="text" size="small" @click="fetchScreens" :loading="loading">
          <template #icon><RefreshIcon /></template>
          刷新
        </t-button>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div 
          v-for="s in screens" 
          :key="s.id"
          class="p-4 rounded-xl bg-[var(--td-bg-color-container)] border border-[var(--td-border-level-1-color)] hover:shadow-md transition-all duration-200"
        >
          <div class="flex items-center justify-between mb-3">
            <div class="flex items-center gap-2">
              <DesktopIcon class="text-[var(--td-brand-color)]" v-if="s.isPrimary" />
              <DesktopIcon class="text-[var(--td-text-color-placeholder)]" v-else />
              <span class="text-sm font-bold">{{ s.name }}</span>
            </div>
            <t-tag v-if="s.isPrimary" theme="primary" variant="light" size="small">主显示器</t-tag>
          </div>
          
          <div class="space-y-2 text-xs">
            <div class="flex justify-between">
              <span class="text-[var(--td-text-color-placeholder)]">逻辑坐标 (DIP)</span>
              <span class="font-mono text-[var(--td-brand-color)]">({{ s.x }}, {{ s.y }})</span>
            </div>
            <div class="flex justify-between">
              <span class="text-[var(--td-text-color-placeholder)]">物理坐标 (PX)</span>
              <span class="font-mono text-[var(--td-success-color)]">({{ s.px }}, {{ s.py }})</span>
            </div>
            <div class="flex justify-between">
              <span class="text-[var(--td-text-color-placeholder)]">DPI 缩放倍率</span>
              <span class="font-mono">{{ s.scaleFactor }}x</span>
            </div>
            <div class="flex justify-between">
              <span class="text-[var(--td-text-color-placeholder)]">物理像素尺寸</span>
              <span class="font-mono opacity-60">{{ s.pWidth }} x {{ s.pHeight }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.drop-shadow-glow {
  filter: drop-shadow(0 0 8px var(--td-brand-color-light));
}
</style>

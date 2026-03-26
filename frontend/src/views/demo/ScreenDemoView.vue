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
import { handleResult } from '@/api/base'

const screens = ref<any[]>([])
const loading = ref(false)
const previewContainer = ref<HTMLElement | null>(null)
const containerWidth = ref(800)
let timer: any = null

const updateContainerSize = () => {
  if (previewContainer.value) {
    containerWidth.value = previewContainer.value.clientWidth
  }
}

const fetchScreens = async () => {
  loading.value = true
  try {
    const res = await ScreenService.GetAllScreens()
    screens.value = handleResult<any[]>(res, true) || []
  } catch (err) {
    // 静默失败，已有 handleResult 弹出或忽略
    console.error(err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchScreens()
  updateContainerSize()
  window.addEventListener('resize', updateContainerSize)
  timer = setInterval(fetchScreens, 3000)
})

onUnmounted(() => {
  if (timer) clearInterval(timer)
  window.removeEventListener('resize', updateContainerSize)
})

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

const getScreenStyle = (s: any) => {
  const cWidth = containerWidth.value
  const cHeight = 280
  const bb = boundingBox.value
  
  const padding = 40
  const availableWidth = cWidth - padding * 2
  const availableHeight = cHeight - padding * 2
  
  const scale = Math.min(availableWidth / bb.width, availableHeight / bb.height)
  
  return {
    left: `${(s.x - bb.minX) * scale + (cWidth - bb.width * scale) / 2}px`,
    top: `${(s.y - bb.minY) * scale + (cHeight - bb.height * scale) / 2}px`,
    width: `${s.width * scale}px`,
    height: `${s.height * scale}px`,
  }
}
</script>

<template>
  <div class="space-y-6">
    <div class="p-4 rounded-lg bg-[var(--td-brand-color-light)] text-[var(--td-brand-color)] text-sm leading-relaxed">
      {{ $t('demo.screenDesc') }}
      <p class="mt-2 text-xs opacity-70">响应式预览：拖动窗口会自动调整图形比例。宽屏模式下可显示 4 列数据卡片。</p>
    </div>

    <!-- 响应式图形化预览 -->
    <div 
      ref="previewContainer"
      class="relative h-[280px] rounded-2xl bg-[var(--td-bg-color-secondarycontainer)] border border-[var(--td-border-level-1-color)] overflow-hidden shadow-inner"
    >
      <div class="absolute inset-0 opacity-5 pointer-events-none p-4 font-mono text-[120px] font-bold select-none overflow-hidden whitespace-nowrap">
        SYSTEM TOPOLOGY
      </div>
      
      <div 
        v-for="(s, idx) in screens" 
        :key="s.id"
        class="absolute border-2 rounded-lg transition-all duration-300 flex flex-col items-center justify-center p-2 shadow-lg group overflow-hidden"
        :class="[
          s.isPrimary ? 'border-[var(--td-brand-color)] bg-[var(--td-bg-color-container)]' : 'border-[var(--td-component-border)] bg-[var(--td-bg-color-secondarycontainer)]'
        ]"
        :style="getScreenStyle(s)"
      >
        <DesktopIcon :size="s.width < 1000 ? '24' : '48'" :class="s.isPrimary ? 'text-[var(--td-brand-color)]' : 'text-[var(--td-text-color-placeholder)]'" />
        <p class="mt-2 text-[10px] font-bold truncate max-w-full px-1" :class="s.isPrimary ? 'text-[var(--td-brand-color)]' : 'text-[var(--td-text-color-secondary)]'">
           {{ s.name || `Screen ${idx+1}` }}
        </p>
        <span v-if="s.isPrimary" class="absolute top-1 right-1 text-[8px] px-1 bg-[var(--td-brand-color)] text-white rounded">PRIMARY</span>
        
        <div class="absolute inset-0 bg-black/80 text-white opacity-0 group-hover:opacity-100 transition-opacity duration-200 flex flex-col items-center justify-center text-[10px] space-y-1">
           <p class="font-bold">{{ s.width }}x{{ s.height }}</p>
           <p class="opacity-70">Pos: ({{ s.x }}, {{ s.y }})</p>
        </div>
      </div>
    </div>

    <!-- 数据列表 (高度响应式栅格) -->
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

      <t-row :gutter="[16, 16]">
        <t-col 
          v-for="s in screens" 
          :key="s.id"
          :xs="12" 
          :sm="6"
          :lg="4"
          :xl="3"
        >
          <div class="p-4 rounded-xl bg-[var(--td-bg-color-container)] border border-[var(--td-border-level-1-color)] hover:shadow-md transition-all duration-200 h-full">
            <div class="flex items-center justify-between mb-3">
              <div class="flex items-center gap-2">
                <DesktopIcon class="text-[var(--td-brand-color)]" v-if="s.isPrimary" />
                <DesktopIcon class="text-[var(--td-text-color-placeholder)]" v-else />
                <span class="text-sm font-bold truncate">{{ s.name }}</span>
              </div>
              <t-tag v-if="s.isPrimary" theme="primary" variant="light" size="small">主屏</t-tag>
            </div>
            
            <div class="space-y-2 text-xs">
              <div class="flex justify-between">
                <span class="text-[var(--td-text-color-placeholder)]">逻辑坐标</span>
                <span class="font-mono text-[var(--td-brand-color)] text-[10px]">({{ s.x }}, {{ s.y }})</span>
              </div>
              <div class="flex justify-between">
                <span class="text-[var(--td-text-color-placeholder)]">DPI 缩放</span>
                <span class="font-mono text-[10px]">{{ s.scaleFactor }}x</span>
              </div>
              <div class="flex justify-between">
                <span class="text-[var(--td-text-color-placeholder)]">物理尺寸</span>
                <span class="font-mono opacity-60 text-[10px]">{{ s.pWidth }} x {{ s.pHeight }}</span>
              </div>
            </div>
          </div>
        </t-col>
      </t-row>
    </div>
  </div>
</template>

<style scoped>
.drop-shadow-glow {
  filter: drop-shadow(0 0 8px var(--td-brand-color-light));
}
</style>

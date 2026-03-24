<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { Events } from '@wailsio/runtime'
import {
  CloudUploadIcon,
  FileIcon,
  DeleteIcon,
  CheckCircleFilledIcon
} from 'tdesign-icons-vue-next'
import { MessagePlugin } from 'tdesign-vue-next'

const droppedItems = ref<string[]>([])
let unsubscribeDrops: (() => void) | null = null

onMounted(() => {
  // 监听从 Go 后端转发过来的原生拖放事件
  // Wails JS runtime 的 Events.On 产生的回调参数是一个事件对象 { name, data, sender }
  unsubscribeDrops = Events.On('files-dropped', (event: any) => {
    console.log('[DragDrop] Got files-dropped event:', event)
    
    // 实际数据在 event.data 中
    const data = event.data
    let paths: string[] = []
    
    if (Array.isArray(data)) {
      paths = data
    } else if (data && typeof data === 'object' && Array.isArray(data.paths)) {
      paths = data.paths
    }

    if (paths.length > 0) {
      // 去重并限制数量
      const newItems = [...new Set([...paths, ...droppedItems.value])]
      droppedItems.value = newItems.slice(0, 50)
      MessagePlugin.success(`接收到 ${paths.length} 个资源`)
    }
  })
})

onUnmounted(() => {
  unsubscribeDrops?.()
})

function clearItems() {
  droppedItems.value = []
}
</script>

<template>
  <div class="space-y-6 animate-fade-in">
    <!-- 说明 -->
    <div class="p-4 rounded-lg bg-[var(--td-brand-color-light)] text-[var(--td-brand-color)] text-sm leading-relaxed">
      {{ $t('demo.dragDropDesc') }}
      <p class="mt-2 text-xs opacity-70">注：Wails 会在原生拖拽进入该区域时自动添加 <code>.file-drop-target-active</code> 类。</p>
    </div>

    <!-- 拖拽投放区 -->
    <div 
      id="native-drop-zone"
      class="drop-zone h-48 border-2 border-dashed rounded-2xl flex flex-col items-center justify-center transition-all duration-300"
      data-file-drop-target="true"
    >
      <div class="flex flex-col items-center pointer-events-none">
        <CloudUploadIcon 
          size="48" 
          class="drop-icon text-[var(--td-text-color-placeholder)]"
        />
        <p class="mt-4 text-sm font-medium text-[var(--td-text-color-secondary)] drop-text">
          {{ $t('demo.dragFilesHere') }}
        </p>
      </div>
    </div>

    <!-- 结果列表 -->
    <div class="space-y-4">
      <div class="flex items-center justify-between">
        <h3 class="text-sm font-medium text-[var(--td-text-color-primary)]">
          {{ $t('demo.droppedList') }} ({{ droppedItems.length }})
        </h3>
        <t-link theme="primary" size="small" @click="clearItems" v-if="droppedItems.length">
          <template #prefix-icon><DeleteIcon /></template>
          {{ $t('common.clear') }}
        </t-link>
      </div>

      <div class="space-y-2">
        <transition-group name="list">
          <div 
            v-for="(path, idx) in droppedItems" 
            :key="path + idx"
            class="flex items-center gap-3 p-3 rounded-xl bg-[var(--td-bg-color-container)] border border-[var(--td-border-level-1-color)] hover:shadow-md transition-all duration-200"
          >
            <div class="p-2 rounded-lg bg-[var(--td-bg-color-secondarycontainer)] text-[var(--td-brand-color)]">
              <FileIcon size="18" />
            </div>
            <div class="flex-1 min-w-0">
              <p class="text-xs text-[var(--td-text-color-secondary)] truncate font-mono">
                {{ path }}
              </p>
            </div>
            <CheckCircleFilledIcon class="text-[var(--td-success-color)] opacity-60" size="14" />
          </div>
        </transition-group>

        <!-- 空状态 -->
        <div v-if="!droppedItems.length" class="py-12 flex flex-col items-center justify-center text-[var(--td-text-color-placeholder)]">
          <FileIcon size="32" class="opacity-20 mb-2" />
          <p class="text-xs">{{ $t('demo.noDroppedItems') }}</p>
        </div>
      </div>
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

/* Wails 原生激活类 */
.drop-zone {
  border-color: var(--td-component-border);
  background-color: var(--td-bg-color-secondarycontainer);
}

.drop-zone.file-drop-target-active {
  border-color: var(--td-brand-color);
  background-color: var(--td-brand-color-light);
  transform: scale(1.01);
  box-shadow: var(--td-shadow-1);
}

.drop-zone.file-drop-target-active .drop-icon {
  color: var(--td-brand-color);
  transform: scale(1.1);
}

.drop-zone.file-drop-target-active .drop-text {
  color: var(--td-brand-color);
}

.drop-icon, .drop-text {
  transition: all 0.3s ease;
}

/* 列表动画 */
.list-enter-active,
.list-leave-active {
  transition: all 0.3s ease;
}
.list-enter-from {
  opacity: 0;
  transform: translateX(-20px);
}
.list-leave-to {
  opacity: 0;
  transform: scale(0.95);
}
</style>

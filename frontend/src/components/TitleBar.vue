<script setup lang="ts">
import { useWindowControl } from '@/composables'
import { RectangleIcon, MinusIcon, CloseIcon } from 'tdesign-icons-vue-next'

defineProps<{
  title?: string
  noMinimize?: boolean
  noMaximize?: boolean
}>()

const { minimise, toggleMaximise, close } = useWindowControl()
</script>

<template>
  <header
    class="h-[32px] flex items-center justify-between pl-3 shrink-0 bg-[var(--td-bg-color-container)] select-none border-b border-[var(--td-border-level-1-color)]"
    style="--wails-draggable: drag; -webkit-app-region: drag"
  >
    <span class="text-xs text-[var(--td-text-color-secondary)]">{{ title }}</span>

    <div class="flex h-full items-stretch" style="-webkit-app-region: no-drag">
      <div v-if="!noMinimize" class="control-btn" @click="minimise">
        <MinusIcon size="14" />
      </div>
      <div v-if="!noMaximize" class="control-btn" @click="toggleMaximise">
        <RectangleIcon size="12" />
      </div>
      <div class="control-btn hover:bg-[#e81123] hover:text-white" @click="close">
        <CloseIcon size="14" />
      </div>
    </div>
  </header>
</template>

<style scoped>
.control-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 100%;
  cursor: pointer;
  transition: all 0.2s;
  color: var(--td-text-color-primary);
}

.control-btn:hover {
  background-color: var(--td-bg-color-secondarycontainer);
}

.control-btn :deep(svg) {
  display: block;
  margin: auto;
}
</style>

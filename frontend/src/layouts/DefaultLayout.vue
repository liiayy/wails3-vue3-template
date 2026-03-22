<script setup lang="ts">
</script>

<template>
  <div class="h-screen w-full flex flex-col overflow-hidden bg-[var(--td-bg-color-page)]">
    <!-- Wails3 窗口拖拽区域（如果你配置了无边框窗口） -->
    <div class="w-full h-[50px] titlebar-drag-region flex items-center px-4 bg-[var(--td-bg-color-container)] border-b border-[var(--td-border-level-1-color)]" style="--wails-draggable:drag">
      <h1 class="text-sm font-semibold text-[var(--td-text-color-primary)]">MyApp2</h1>
    </div>

    <!-- 主体内容区 -->
    <main class="flex-1 overflow-auto p-4">
      <router-view />
    </main>
  </div>
</template>

<style scoped>
/* 定义拖拽区域 */
.titlebar-drag-region {
  -webkit-app-region: drag;
  user-select: none;
}
</style>

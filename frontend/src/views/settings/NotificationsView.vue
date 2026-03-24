<script setup lang="ts">
import { NotificationIcon } from 'tdesign-icons-vue-next'
</script>

<template>
  <div class="flex flex-col items-center justify-center py-10 opacity-30 grayscale">
    <NotificationIcon size="60" />
    <p class="mt-4 text-lg font-medium">{{ $t('settings.comingSoon') }}</p>
  </div>
</template>

<style scoped>
/* 设置项样式 */
</style>

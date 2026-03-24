<script setup lang="ts">
import { NotificationIcon } from 'tdesign-icons-vue-next'
</script>

<template>
  <div class="flex flex-col items-center justify-center py-10 opacity-30 grayscale animate-fade-in">
    <NotificationIcon size="60" />
    <p class="mt-4 text-lg font-medium">{{ $t('settings.comingSoon') }}</p>
  </div>
</template>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.1s cubic-bezier(0.16, 1, 0.3, 1);
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
</style>

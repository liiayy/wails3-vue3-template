<script setup lang="ts">
import { useRoute } from 'vue-router'
import TitleBar from '../components/TitleBar.vue'

const route = useRoute()
const isStandalone = route.path.startsWith('/standalone')
</script>

<template>
  <div class="h-full flex flex-col bg-[var(--td-bg-color-page)] overflow-hidden">
    <TitleBar v-if="isStandalone" :title="$t('about.title')" no-minimize no-maximize />
    
    <div class="flex-1 flex flex-col items-center justify-center p-10 space-y-4 overflow-auto">
      <div class="w-20 h-20 rounded-full bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center text-white text-3xl font-bold shadow-lg">
        M
      </div>
      <h1 class="text-xl font-bold text-[var(--td-text-color-primary)]">{{ $t('about.title') }}</h1>
      <p class="text-sm text-[var(--td-text-color-secondary)]">
        {{ $t('about.version') }}: v1.0.0-alpha
      </p>
      <t-divider class="w-full" />
      <div class="text-center text-sm text-[var(--td-text-color-placeholder)] space-y-1">
        <p>{{ $t('about.description') }}</p>
        <p>{{ $t('about.copyright') }}</p>
      </div>
    </div>
  </div>
</template>

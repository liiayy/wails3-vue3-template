<script setup lang="ts">
import { useRoute } from 'vue-router'
import TitleBar from '@/components/TitleBar.vue'

const route = useRoute()
const isStandalone = route.path.startsWith('/standalone')
</script>

<template>
  <div class="h-full flex flex-col bg-[var(--td-bg-color-page)] overflow-hidden">
    <TitleBar v-if="isStandalone" :title="$t('about.title')" no-minimize no-maximize />

    <div class="flex-1 overflow-auto bg-[var(--td-bg-color-page)] flex items-center justify-center p-8">
      <t-row justify="center" class="w-full max-w-lg">
        <t-col :span="12" class="text-center space-y-6">
          <!-- Logo 栅格分布 -->
          <div class="flex justify-center">
            <div class="w-24 h-24 rounded-[28%] bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center text-white text-4xl font-bold shadow-2xl rotate-3 hover:rotate-0 transition-transform duration-500">
              M
            </div>
          </div>

          <div class="space-y-2">
            <h1 class="text-2xl font-bold text-[var(--td-text-color-primary)] tracking-tight">
              {{ $t('about.title') }}
            </h1>
            <p class="text-sm font-medium text-[var(--td-brand-color)] bg-[var(--td-brand-color-light)] inline-block px-3 py-1 rounded-full">
               v1.2.0-Production
            </p>
          </div>

          <t-divider border-style="dashed" />

          <div class="space-y-4 text-[var(--td-text-color-secondary)]">
            <p class="text-sm leading-relaxed">{{ $t('about.description') }}</p>
            <div class="pt-4 text-xs text-[var(--td-text-color-placeholder)] font-mono">
              <p>{{ $t('about.copyright') }}</p>
              <p class="mt-1 opacity-60 italic">Built with Wails 3 & TDesign</p>
            </div>
          </div>
        </t-col>
      </t-row>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, watch, onMounted, onUnmounted } from 'vue'
import { merge } from 'lodash-es'
import enConfig from 'tdesign-vue-next/es/locale/en_US'
import zhConfig from 'tdesign-vue-next/es/locale/zh_CN'
import { useSettingsStore } from '@/stores/settings'
import i18n from '@/locales'

const settings = useSettingsStore()

// 1. 国际化配置 (TDesign)
const globalConfig = computed(() => {
  const customConfig = {}
  if (settings.language === 'en-US') {
    return merge({}, enConfig, customConfig)
  }
  return merge({}, zhConfig, customConfig)
})

// 2. 监听语言变化同步给 vue-i18n
watch(
  () => settings.language,
  (lang) => {
    if (i18n.global.locale) {
      ;(i18n.global.locale as any).value = lang
    }
  },
  { immediate: true },
)

// 3. 监听主题变化操作 DOM
const applyTheme = () => {
  const doc = document.documentElement
  let targetTheme = settings.theme

  if (targetTheme === 'auto') {
    const isDark = window.matchMedia('(prefers-color-scheme: dark)').matches
    targetTheme = isDark ? 'dark' : 'light'
  }

  if (targetTheme === 'dark') {
    doc.setAttribute('theme-mode', 'dark')
    doc.classList.add('dark')
  } else {
    doc.removeAttribute('theme-mode')
    doc.classList.remove('dark')
  }
}

watch(() => settings.theme, applyTheme, { immediate: true })

// 监听系统主题变化（仅在 auto 模式下有效）
const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
const handleSystemThemeChange = () => {
  if (settings.theme === 'auto') {
    applyTheme()
  }
}

onMounted(() => {
  mediaQuery.addEventListener('change', handleSystemThemeChange)
})

onUnmounted(() => {
  mediaQuery.removeEventListener('change', handleSystemThemeChange)
})

// 4. 监听缩放变化
watch(
  () => settings.zoom,
  (zoom) => {
    ;(document.body.style as any).zoom = `${zoom}%`
  },
  { immediate: true },
)
</script>

<template>
  <t-config-provider :global-config="globalConfig">
    <router-view />
  </t-config-provider>
</template>

<style scoped></style>

<script setup lang="ts">
import { computed } from 'vue'
import { merge } from 'lodash-es'
import enConfig from 'tdesign-vue-next/es/locale/en_US'
import zhConfig from 'tdesign-vue-next/es/locale/zh_CN'
import { useSettingsStore } from '@/stores/settings'

const settings = useSettingsStore()

const globalConfig = computed(() => {
  const customConfig = {}
  if (settings.language === 'en-US') {
    return merge({}, enConfig, customConfig)
  }
  return merge({}, zhConfig, customConfig)
})
</script>

<template>
  <t-config-provider :global-config="globalConfig">
    <router-view />
  </t-config-provider>
</template>

<style scoped></style>

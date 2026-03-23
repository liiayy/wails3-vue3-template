<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useSettingsStore } from '../stores/settings'
import TitleBar from '../components/TitleBar.vue'

const settings = useSettingsStore()
const route = useRoute()
const isStandalone = route.path.startsWith('/standalone')
</script>

<template>
  <div class="h-full flex flex-col bg-[var(--td-bg-color-page)] overflow-hidden">
    <TitleBar v-if="isStandalone" :title="$t('settings.title')" no-minimize no-maximize />

    <div class="flex-1 overflow-auto py-10 px-6 space-y-8">
      <div class="max-w-xl mx-auto space-y-8">
        <h1 class="text-2xl font-bold text-[var(--td-text-color-primary)]">{{ $t('settings.title') }}</h1>

        <!-- 主题设置 -->
        <t-card :title="$t('settings.themeLabel')" :bordered="false" class="shadow-sm">
          <div class="flex justify-between items-center">
            <span>{{ $t('settings.themeLabel') }}</span>
            <t-radio-group
              variant="default-filled"
              :value="settings.theme"
              @change="(val: any) => settings.updateSetting('theme', val)"
            >
              <t-radio-button value="light">{{ $t('settings.themeLight') }}</t-radio-button>
              <t-radio-button value="dark">{{ $t('settings.themeDark') }}</t-radio-button>
            </t-radio-group>
          </div>
        </t-card>

        <!-- 语言设置 -->
        <t-card :title="$t('settings.languageLabel')" :bordered="false" class="shadow-sm">
          <div class="flex justify-between items-center">
            <span>{{ $t('settings.languageLabel') }}</span>
            <t-select
              :value="settings.language"
              @change="(val: any) => settings.updateSetting('language', val)"
              style="width: 200px"
            >
              <t-option value="zh-CN" label="简体中文" />
              <t-option value="en-US" label="English" />
            </t-select>
          </div>
        </t-card>

        <!-- 侧边栏设置 -->
        <t-card :title="$t('settings.sidebarLabel')" :bordered="false" class="shadow-sm">
          <div class="flex justify-between items-center">
            <span>{{ $t('settings.sidebarLabel') }}</span>
            <t-switch
              :value="settings.isSidebarCollapsed"
              @change="(val: any) => settings.updateSetting('isSidebarCollapsed', val)"
            />
          </div>
        </t-card>
      </div>
    </div>
  </div>
</template>

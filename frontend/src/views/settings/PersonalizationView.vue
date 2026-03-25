<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useSettingsStore } from '../../stores/settings'
import SettingSection from '@/views/settings/components/SettingSection.vue'

const settings = useSettingsStore()
const { t } = useI18n()

// 通用设置项配置
const generalItems = computed(() => [
  {
    key: 'language',
    label: t('settings.displayLanguage'),
    type: 'select',
    value: settings.language,
    options: [
      { value: 'zh-CN', label: '简体中文' },
      { value: 'en-US', label: 'English (United States)' },
    ],
  },
  {
    key: 'isSidebarCollapsed',
    label: t('settings.autoHideSidebar'),
    type: 'switch',
    value: settings.isSidebarCollapsed,
  },
  {
    key: 'isAutostart',
    label: t('settings.isAutostart'),
    description: t('settings.isAutostartDesc'),
    type: 'switch',
    value: settings.isAutostart,
  },
])

function handleSettingChange(key: string, value: any) {
  settings.updateSetting(key, value)
}
</script>

<template>
  <div class="w-full max-w-4xl mx-auto space-y-5 rounded-lg text-[var(--td-text-color-primary)]">
    <!-- 外观设置 -->
    <section class="space-y-2">
      <div class="text-[15px] font-bold px-6">
        {{ $t('settings.themeMode') }}
      </div>

      <div class="bg-[var(--td-bg-color-container)] p-4 rounded-2xl shadow-sm flex gap-8">
        <!-- 明亮模式预览卡片 -->
        <div
          class="group relative cursor-pointer w-[130px] flex flex-col items-center gap-3"
          @click="settings.updateSetting('theme', 'light')"
        >
          <div
            class="w-full aspect-[16/11] rounded-xl border-[3px] transition-all duration-300 p-2 bg-white overflow-hidden flex flex-col gap-1.5"
            :class="
              settings.theme === 'light'
                ? 'border-[var(--td-brand-color)]'
                : 'border-transparent hover:border-gray-300 ring-1 ring-gray-200'
            "
          >
            <div class="h-1.5 w-1/3 bg-gray-100 rounded"></div>
            <div class="h-5 w-full bg-white border border-gray-100 rounded shadow-sm"></div>
            <div class="flex gap-1.5 flex-1">
              <div class="w-1/3 bg-gray-50 rounded"></div>
              <div class="w-2/3 bg-gray-50 rounded"></div>
            </div>
          </div>
          <span class="text-[13px]">
            {{ $t('settings.themeLight') }}
          </span>
        </div>

        <!-- 黑暗模式预览卡片 -->
        <div
          class="group relative cursor-pointer w-[130px] flex flex-col items-center gap-3"
          @click="settings.updateSetting('theme', 'dark')"
        >
          <div
            class="w-full aspect-[16/11] rounded-xl border-[3px] transition-all duration-300 p-2 bg-[#181818] overflow-hidden flex flex-col gap-1.5"
            :class="
              settings.theme === 'dark'
                ? 'border-[var(--td-brand-color)]'
                : 'border-transparent hover:border-gray-500 ring-1 ring-gray-600'
            "
          >
            <div class="h-1.5 w-1/3 bg-gray-800 rounded"></div>
            <div class="h-5 w-full bg-gray-800 border border-gray-700 rounded shadow-sm"></div>
            <div class="flex gap-1.5 flex-1">
              <div class="w-1/3 bg-gray-800 rounded"></div>
              <div class="w-2/3 bg-gray-800 rounded"></div>
            </div>
          </div>
          <span class="text-[13px]">
            {{ $t('settings.themeDark') }}
          </span>
        </div>

        <!-- 跟随系统预览卡片 -->
        <div
          class="group relative cursor-pointer w-[130px] flex flex-col items-center gap-3"
          @click="settings.updateSetting('theme', 'auto')"
        >
          <div
            class="w-full aspect-[16/11] rounded-xl border-[3px] transition-all duration-300 overflow-hidden flex"
            :class="
              settings.theme === 'auto'
                ? 'border-[var(--td-brand-color)]'
                : 'border-transparent hover:border-gray-300 ring-1 ring-gray-400'
            "
          >
            <div class="flex-1 bg-white p-2 flex flex-col gap-1.5 border-r border-gray-200">
              <div class="h-1 w-2/3 bg-gray-100 rounded"></div>
              <div class="h-3 w-full bg-gray-50 border border-gray-100 rounded-sm"></div>
            </div>
            <div class="flex-1 bg-[#181818] p-2 flex flex-col gap-1.5">
              <div class="h-1 w-2/3 bg-gray-800 rounded"></div>
              <div class="h-3 w-full bg-gray-800 border border-gray-700 rounded-sm"></div>
            </div>
          </div>
          <span class="text-[13px]">
            {{ $t('settings.themeAuto') }}
          </span>
        </div>
      </div>
    </section>

    <SettingSection
      :title="$t('settings.general')"
      :items="generalItems"
      @change="handleSettingChange"
    />
  </div>
</template>

<style scoped>
/* 使用TDesign提供的颜色与布局变量来实现深色与浅色的全自适应 */
</style>

<script setup lang="ts">
import { useSettingsStore } from '../../stores/settings'
import { CheckCircleFilledIcon } from 'tdesign-icons-vue-next'

const settings = useSettingsStore()
</script>

<template>
  <div class="w-full max-w-6xl mx-auto px-4 py-2 space-y-6">
    <!-- 主题模式 -->
    <section class="space-y-3">
      <div
        class="text-[11px] font-bold text-[var(--td-text-color-placeholder)] uppercase tracking-wider"
      >
        {{ $t('settings.themeMode') }}
      </div>
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3">
        <!-- 明亮模式预览卡片 -->
        <div
          class="group relative cursor-pointer w-full"
          @click="settings.updateSetting('theme', 'light')"
        >
          <div
            class="aspect-[16/10] rounded-xl border-2 transition-all duration-300 p-2.5 bg-white shadow-sm overflow-hidden flex flex-col gap-1"
            :class="
              settings.theme === 'light'
                ? 'border-[var(--td-brand-color)] ring-2 ring-[var(--td-brand-color-light)]'
                : 'border-transparent bg-gray-50 hover:border-gray-200'
            "
          >
            <!-- 模拟界面 -->
            <div class="h-1.5 w-1/3 bg-gray-100 rounded"></div>
            <div class="h-5 w-full bg-white border border-gray-100 rounded shadow-sm"></div>
            <div class="flex gap-1 flex-1">
              <div class="w-1/3 bg-gray-50 rounded"></div>
              <div class="w-2/3 bg-gray-50 rounded"></div>
            </div>
          </div>
          <div class="mt-2 flex items-center justify-between px-1">
            <span class="font-medium text-[12px] text-[var(--td-text-color-primary)]">
              {{ $t('settings.themeLight') }}
            </span>
            <CheckCircleFilledIcon
              v-if="settings.theme === 'light'"
              class="text-[var(--td-brand-color)]"
              size="16"
            />
            <div v-else class="w-3.5 h-3.5 rounded-full border-2 border-gray-200"></div>
          </div>
        </div>

        <!-- 黑暗模式预览卡片 -->
        <div class="group relative cursor-pointer w-full" @click="settings.updateSetting('theme', 'dark')">
          <div
            class="aspect-[16/10] rounded-xl border-2 transition-all duration-300 p-2.5 bg-[#181818] shadow-sm overflow-hidden flex flex-col gap-1"
            :class="
              settings.theme === 'dark'
                ? 'border-[var(--td-brand-color)] ring-2 ring-[var(--td-brand-color-light)]'
                : 'border-transparent bg-gray-900 hover:border-gray-700'
            "
          >
            <!-- 模拟界面 -->
            <div class="h-1.5 w-1/3 bg-gray-800 rounded"></div>
            <div class="h-5 w-full bg-gray-800 border border-gray-700 rounded shadow-sm"></div>
            <div class="flex gap-1 flex-1">
              <div class="w-1/3 bg-gray-800 rounded"></div>
              <div class="w-2/3 bg-gray-800 rounded"></div>
            </div>
          </div>
          <div class="mt-2 flex items-center justify-between px-1">
            <span class="font-medium text-[12px] text-[var(--td-text-color-primary)]">
              {{ $t('settings.themeDark') }}
            </span>
            <CheckCircleFilledIcon
              v-if="settings.theme === 'dark'"
              class="text-[var(--td-brand-color)]"
              size="16"
            />
            <div v-else class="w-3.5 h-3.5 rounded-full border-2 border-gray-200"></div>
          </div>
        </div>

        <!-- 跟随系统预览卡片 (Split Design) -->
        <div class="group relative cursor-pointer w-full" @click="settings.updateSetting('theme', 'auto')">
          <div
            class="aspect-[16/10] rounded-xl border-2 transition-all duration-300 shadow-sm overflow-hidden flex"
            :class="
              settings.theme === 'auto'
                ? 'border-[var(--td-brand-color)] ring-2 ring-[var(--td-brand-color-light)]'
                : 'border-transparent bg-gray-100 hover:border-gray-200'
            "
          >
            <!-- 左半部分：亮色 -->
            <div class="flex-1 bg-white p-2 flex flex-col gap-1 border-r border-gray-100">
              <div class="h-1 w-2/3 bg-gray-100 rounded"></div>
              <div class="h-4 w-full bg-gray-50 border border-gray-100 rounded-sm"></div>
            </div>
            <!-- 右半部分：暗色 -->
            <div class="flex-1 bg-[#181818] p-2 flex flex-col gap-1">
              <div class="h-1 w-2/3 bg-gray-800 rounded"></div>
              <div class="h-4 w-full bg-gray-800 border border-gray-700 rounded-sm"></div>
            </div>
          </div>
          <div class="mt-2 flex items-center justify-between px-1">
            <span class="font-medium text-[12px] text-[var(--td-text-color-primary)]">
              {{ $t('settings.themeAuto') }}
            </span>
            <CheckCircleFilledIcon
              v-if="settings.theme === 'auto'"
              class="text-[var(--td-brand-color)]"
              size="16"
            />
            <div v-else class="w-3.5 h-3.5 rounded-full border-2 border-gray-200"></div>
          </div>
        </div>
      </div>
    </section>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-4 lg:gap-6">
      <!-- 语言设置简版 -->
      <t-card
        :bordered="false"
        class="bg-[var(--td-bg-color-container)] rounded-xl p-0.5 shadow-sm"
      >
        <div class="space-y-4">
          <div class="text-[10px] font-bold text-[var(--td-text-color-placeholder)] uppercase">
            {{ $t('settings.displayLanguage') }}
          </div>
          <t-select
            :value="settings.language"
            @change="(val: any) => settings.updateSetting('language', val)"
            class="w-full"
            variant="filled"
          >
            <t-option value="zh-CN" label="简体中文 (Chinese)" />
            <t-option value="en-US" label="English (United States)" />
          </t-select>
          <p class="text-[11px] text-[var(--td-text-color-placeholder)] italic">
            {{ $t('settings.langDesc') }}
          </p>
        </div>
      </t-card>

      <!-- 侧边栏设置简版 -->
      <t-card
        :bordered="false"
        class="bg-[var(--td-bg-color-container)] rounded-xl p-0.5 shadow-sm"
      >
        <div class="flex justify-between items-start mb-2">
          <div class="space-y-1">
            <div class="text-[10px] font-bold text-[var(--td-text-color-placeholder)] uppercase">
              {{ $t('settings.sidebarBehavior') }}
            </div>
            <div class="font-bold text-[var(--td-text-color-primary)] text-sm">
              {{ $t('settings.autoHideSidebar') }}
            </div>
          </div>
          <t-switch
            :value="settings.isSidebarCollapsed"
            @change="(val: any) => settings.updateSetting('isSidebarCollapsed', val)"
          />
        </div>
        <p class="text-[11px] text-[var(--td-text-color-secondary)]">
          {{ $t('settings.sidebarDesc') }}
        </p>
      </t-card>

      <!-- 自启动设置 -->
      <t-card
        :bordered="false"
        class="bg-[var(--td-bg-color-container)] rounded-xl p-0.5 shadow-sm"
      >
        <div class="flex justify-between items-start mb-2">
          <div class="space-y-1">
            <div class="text-[10px] font-bold text-[var(--td-text-color-placeholder)] uppercase">
              {{ $t('common.confirm') }}
            </div>
            <div class="font-bold text-[var(--td-text-color-primary)] text-sm">
              {{ $t('settings.isAutostart') }}
            </div>
          </div>
          <t-switch
            :value="settings.isAutostart"
            @change="(val: any) => settings.updateSetting('isAutostart', val)"
          />
        </div>
        <p class="text-[11px] text-[var(--td-text-color-secondary)]">
          {{ $t('settings.isAutostartDesc') }}
        </p>
      </t-card>
    </div>
  </div>
</template>

<style scoped>
:deep(.t-card) {
  transition: transform 0.2s;
}
:deep(.t-card:hover) {
  transform: translateY(-2px);
}

/* 响应式优化：小屏幕时减少内边距 */
@media (max-width: 640px) {
  div[class*="aspect-[16/10]"] {
    aspect-ratio: 16/9;
  }
}
</style>

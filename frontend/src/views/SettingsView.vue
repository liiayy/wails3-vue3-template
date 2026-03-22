<script setup lang="ts">
import { useSettingsStore } from '../stores/settings'

const settings = useSettingsStore()
</script>

<template>
  <div class="max-w-xl mx-auto py-10 px-6 space-y-8">
    <h1 class="text-2xl font-bold text-[var(--td-text-color-primary)]">系统设置</h1>

    <!-- 主题设置 -->
    <t-card title="外观" :bordered="false" class="shadow-sm">
      <div class="flex justify-between items-center">
        <span>界面主题</span>
        <t-radio-group
          variant="default-filled"
          :value="settings.theme"
          @change="(val: any) => settings.updateSetting('theme', val)"
        >
          <t-radio-button value="light">明亮</t-radio-button>
          <t-radio-button value="dark">黑暗</t-radio-button>
        </t-radio-group>
      </div>
    </t-card>

    <!-- 语言设置 -->
    <t-card title="语言与区域" :bordered="false" class="shadow-sm">
      <div class="flex justify-between items-center">
        <span>界面语言</span>
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
    <t-card title="布局" :bordered="false" class="shadow-sm">
      <div class="flex justify-between items-center">
        <span>折叠侧边栏</span>
        <t-switch
          :value="settings.isSidebarCollapsed"
          @change="(val: any) => settings.updateSetting('isSidebarCollapsed', val)"
        />
      </div>
    </t-card>
  </div>
</template>

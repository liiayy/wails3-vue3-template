<script setup lang="ts">
import { ref } from 'vue'
import { Dialogs } from '@wailsio/runtime'
import {
  FileIcon,
  FileCopyIcon,
  FolderOpenIcon,
  SaveIcon,
  DeleteIcon
} from 'tdesign-icons-vue-next'

const selectedFiles = ref<string[]>([])
const selectedFolder = ref('')
const savedPath = ref('')

async function pickFile() {
  const result = await Dialogs.OpenFile({
    Title: '选择单个文件',
    Filters: [
      { DisplayName: '所有文件', Pattern: '*.*' },
      { DisplayName: '文本文件', Pattern: '*.txt;*.md' },
      { DisplayName: '图片文件', Pattern: '*.png;*.jpg;*.gif' }
    ]
  })
  if (result && typeof result === 'string') {
    selectedFiles.value = [result]
  }
}

async function pickMultipleFiles() {
  const results = await Dialogs.OpenFile({
    Title: '选择多个文件',
    AllowsMultipleSelection: true,
    Filters: [{ DisplayName: '资源文件', Pattern: '*.*' }]
  })
  if (results && Array.isArray(results)) {
    selectedFiles.value = results
  }
}

async function pickDirectory() {
  const result = await Dialogs.OpenFile({
    Title: '选择文件夹',
    CanChooseDirectories: true,
    CanChooseFiles: false
  })
  if (result && typeof result === 'string') {
    selectedFolder.value = result
  }
}

async function saveFile() {
  const result = await Dialogs.SaveFile({
    Title: '保存文件演示',
    Filename: 'demo-export.txt',
    Filters: [{ DisplayName: '文本', Pattern: '*.txt' }]
  })
  if (result) {
    savedPath.value = result
  }
}

function clearResults() {
  selectedFiles.value = []
  selectedFolder.value = ''
  savedPath.value = ''
}
</script>

<template>
  <div class="space-y-6">
    <!-- 操作按钮组 (栅格化) -->
    <t-row :gutter="[16, 16]">
      <t-col :xs="6" :md="3">
        <t-button block variant="outline" @click="pickFile">
          <template #icon><FileIcon /></template>
          {{ $t('demo.pickFile') }}
        </t-button>
      </t-col>
      <t-col :xs="6" :md="3">
        <t-button block variant="outline" @click="pickMultipleFiles">
          <template #icon><FileCopyIcon /></template>
          {{ $t('demo.pickMultiple') }}
        </t-button>
      </t-col>
      <t-col :xs="6" :md="3">
        <t-button block variant="outline" @click="pickDirectory">
          <template #icon><FolderOpenIcon /></template>
          {{ $t('demo.pickFolder') }}
        </t-button>
      </t-col>
      <t-col :xs="6" :md="3">
        <t-button block variant="outline" @click="saveFile">
          <template #icon><SaveIcon /></template>
          {{ $t('demo.saveFile') }}
        </t-button>
      </t-col>
    </t-row>

    <!-- 结果回显区 (栅格化) -->
    <div class="space-y-4">
      <div class="flex items-center justify-between">
        <h3 class="text-sm font-medium text-[var(--td-text-color-primary)]">
          {{ $t('demo.results') }}
        </h3>
        <t-link theme="primary" size="small" @click="clearResults" v-if="selectedFiles.length || selectedFolder || savedPath">
          <template #prefix-icon><DeleteIcon /></template>
          {{ $t('common.clear') }}
        </t-link>
      </div>

      <t-row :gutter="[16, 16]">
        <!-- 文件列表 -->
        <t-col v-if="selectedFiles.length" :span="12">
          <t-card :title="$t('demo.selectedFiles')" size="small" class="bg-[var(--td-bg-color-container)]">
            <ul class="text-xs space-y-1 overflow-hidden">
              <li v-for="file in selectedFiles" :key="file" class="truncate text-[var(--td-text-color-secondary)]">
                {{ file }}
              </li>
            </ul>
          </t-card>
        </t-col>

        <!-- 文件夹 -->
        <t-col v-if="selectedFolder" :xs="12" :md="6">
          <t-card :title="$t('demo.selectedFolder')" size="small" class="bg-[var(--td-bg-color-container)]">
            <p class="text-xs truncate text-[var(--td-text-color-secondary)]">{{ selectedFolder }}</p>
          </t-card>
        </t-col>

        <!-- 保存路径 -->
        <t-col v-if="savedPath" :xs="12" :md="6">
          <t-card :title="$t('demo.savedPath')" size="small" class="bg-[var(--td-bg-color-container)]">
            <p class="text-xs truncate text-[var(--td-text-color-secondary)]">{{ savedPath }}</p>
          </t-card>
        </t-col>

        <!-- 无结果占位 -->
        <t-col v-if="!selectedFiles.length && !selectedFolder && !savedPath" :span="12">
          <div class="py-12 border-2 border-dashed border-[var(--td-component-border)] rounded-xl flex flex-col items-center justify-center text-[var(--td-text-color-placeholder)]">
            <FileIcon size="40" class="opacity-20 mb-2" />
            <p class="text-sm">{{ $t('demo.noResults') }}</p>
          </div>
        </t-col>
      </t-row>
    </div>
  </div>
</template>

<style scoped>
/* 文件演示页样式微调 */
</style>

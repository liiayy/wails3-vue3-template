<script setup lang="ts">
import { ref } from 'vue'
import { Clipboard } from '@wailsio/runtime'
import {
  CopyIcon,
  DownloadIcon,
  DeleteIcon,
  CheckCircleFilledIcon,
  ErrorCircleFilledIcon
} from 'tdesign-icons-vue-next'
import { MessagePlugin } from 'tdesign-vue-next'

const inputContent = ref('Hello Wails 3 Clipboard! 🚀')
const clipboardContent = ref('')
const lastOperation = ref<'read' | 'write' | 'clear' | null>(null)

async function writeToClipboard() {
  if (!inputContent.value) {
    MessagePlugin.warning('内容不能为空')
    return
  }
  try {
    await Clipboard.SetText(inputContent.value)
    lastOperation.value = 'write'
    MessagePlugin.success('已写入剪贴板')
  } catch (err) {
    MessagePlugin.error('写入失败: ' + err)
  }
}

async function readFromClipboard() {
  try {
    const val = await Clipboard.Text()
    clipboardContent.value = val
    lastOperation.value = 'read'
    MessagePlugin.success('成功读取剪贴板')
  } catch (err) {
    MessagePlugin.error('读取失败: ' + err)
  }
}

async function clearClipboard() {
  try {
    await Clipboard.SetText('')
    clipboardContent.value = ''
    lastOperation.value = 'clear'
    MessagePlugin.info('剪贴板已清空')
  } catch (err) {
    MessagePlugin.error('操作失败: ' + err)
  }
}
</script>

<template>
  <div class="space-y-6">
    <t-row :gutter="[24, 24]">
      <!-- 写入区域 -->
      <t-col :xs="12" :md="6">
        <t-card :title="$t('demo.clipboardWrite')" size="small" class="h-full">
          <div class="space-y-4">
            <div class="mb-4">
              <t-textarea  v-model="inputContent" :placeholder="$t('demo.clipboardWritePlaceholder')" />
            </div>
            <t-row :gutter="[8, 8]">
              <t-col :span="6">
                <t-button block @click="writeToClipboard">
                  <template #icon><CopyIcon /></template>
                  {{ $t('demo.writeNow') }}
                </t-button>
              </t-col>
              <t-col :span="6">
                <t-button block theme="danger" variant="outline" @click="clearClipboard">
                  <template #icon><DeleteIcon /></template>
                  {{ $t('demo.clearClipboard') }}
                </t-button>
              </t-col>
            </t-row>
          </div>
        </t-card>
      </t-col>

      <!-- 读取区域 -->
      <t-col :xs="12" :md="6">
        <t-card :title="$t('demo.clipboardRead')" size="small" class="h-full">
          <div class="space-y-4">
            <div 
              class="min-h-[100px] p-3 rounded-lg border-2 border-[var(--td-component-border)] bg-[var(--td-bg-color-secondarycontainer)] text-sm text-[var(--td-text-color-secondary)] whitespace-pre-wrap break-all"
            >
              <span v-if="clipboardContent">{{ clipboardContent }}</span>
              <span v-else class="text-[var(--td-text-color-placeholder)] italic">{{ $t('demo.clickReadToFetch') }}</span>
            </div>
            <t-button block theme="primary" variant="outline" @click="readFromClipboard">
              <template #icon><DownloadIcon /></template>
              {{ $t('demo.readNow') }}
            </t-button>
          </div>
        </t-card>
      </t-col>

      <!-- 操作提示 -->
      <t-col :span="12">
        <div v-show="lastOperation" class="flex items-center gap-2 p-3 rounded-lg bg-[var(--td-success-color-1)] text-[var(--td-success-color-6)] text-xs animate-in slide-in-from-bottom-2 duration-300">
          <CheckCircleFilledIcon />
          <span>{{ $t('demo.lastSuccessAction') }}: {{ $t(`demo.op_${lastOperation}`) }}</span>
        </div>
      </t-col>
    </t-row>
  </div>
</template>

<style scoped>
/* 剪贴板演示页样式 */
</style>

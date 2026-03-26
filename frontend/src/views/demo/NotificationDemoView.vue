<script setup lang="ts">
import { ref, onActivated, onDeactivated } from 'vue'
import { Events } from '@wailsio/runtime'
import { NotificationBinding } from '#/myapp2/internal/binding'
import {
  NotificationIcon,
  ChatIcon,
  CheckCircleFilledIcon,
  TimeIcon,
  UserIcon,
  InfoCircleIcon,
} from 'tdesign-icons-vue-next'
import { MessagePlugin } from 'tdesign-vue-next'
import { handleResult } from '@/api/base'

import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const lastResponse = ref<any>(null)
let unsubscribeNotifications: (() => void) | null = null

onActivated(() => {
  // 适配 KeepAlive：当进入此页面（激活）时开启监听
  console.log('[Notification] Component activated, subscribing to events')
  unsubscribeNotifications = Events.On('notification-clicked', (event: any) => {
    lastResponse.value = {
      ...event.data,
      timestamp: new Date().toLocaleTimeString(),
    }
    const action = event.data.ActionIdentifier || t('common.confirm')
    MessagePlugin.info(`${t('demo.lastSuccessAction')}: ${action}`)
  })
})

onDeactivated(() => {
  // 适配 KeepAlive：当切离此页面（停用）时立即取消监听
  // 这是解决 KeepAlive 模式下 EventListener 重复堆叠的最佳实践
  if (unsubscribeNotifications) {
    unsubscribeNotifications()
    unsubscribeNotifications = null
    console.log('[Notification] Component deactivated, unsubscribed')
  }
})

async function sendBasic() {
  try {
    const res = await NotificationBinding.SendBasic(t('demo.sendBasicTitle'), t('demo.sendBasicBody'))
    handleResult(res)
  } catch (err) {
    // 报错已由 handleResult 处理
  }
}

async function sendSubtitle() {
  try {
    const res = await NotificationBinding.SendWithSubtitle(
      t('demo.sendSubtitleTitle'),
      t('demo.sendSubtitleSub'),
      t('demo.sendSubtitleBody')
    )
    handleResult(res)
  } catch (err) {
    // 报错已由 handleResult 处理
  }
}

async function sendInteractive() {
  try {
    const res = await NotificationBinding.SendInteractive(t('demo.sendInteractiveTitle'), t('demo.interactiveBody'))
    handleResult(res)
  } catch (err) {
    // 报错已由 handleResult 处理
  }
}
</script>

<template>
  <div class="space-y-6">
    <!-- 说明 -->
    <div
      class="p-4 rounded-lg bg-[var(--td-brand-color-light)] text-[var(--td-brand-color)] text-sm leading-relaxed"
    >
      {{ $t('demo.notificationDesc') }}
    </div>

    <!-- 通知操作区 (栅格化) -->
    <t-row :gutter="[16, 16]">
      <!-- 基础通知 -->
      <t-col :xs="12" :md="6">
        <t-card :title="$t('demo.sendBasic')" header-bordered class="h-full">
          <div class="space-y-4">
            <p class="text-xs text-[var(--td-text-color-secondary)]">{{ $t('demo.sendBasicDesc') }}</p>
            <t-button block theme="primary" variant="outline" @click="sendBasic">
              <template #prefixIcon><NotificationIcon /></template>
              {{ $t('demo.sendBasic') }}
            </t-button>
          </div>
        </t-card>
      </t-col>

      <!-- 带副标题通知 -->
      <t-col :xs="12" :md="6">
        <t-card :title="$t('demo.sendSubtitle')" header-bordered class="h-full">
          <div class="space-y-4">
            <p class="text-xs text-[var(--td-text-color-secondary)]">
              {{ $t('demo.sendSubtitleDesc') }}
            </p>
            <t-button block theme="primary" variant="outline" @click="sendSubtitle">
              <template #prefixIcon><InfoCircleIcon /></template>
              {{ $t('demo.sendSubtitle') }}
            </t-button>
          </div>
        </t-card>
      </t-col>

      <!-- 交互式通知 -->
      <t-col :span="12">
        <t-card :title="$t('demo.sendInteractive')" header-bordered>
          <div class="space-y-4">
            <div
              class="p-4 rounded-lg bg-[var(--td-bg-color-secondarycontainer)] border border-[var(--td-border-level-1-color)]"
            >
              <div class="flex items-start gap-4">
                <t-avatar size="medium" shape="round"><UserIcon /></t-avatar>
                <div class="flex-1 min-w-0">
                  <p class="text-sm font-bold text-[var(--td-text-color-primary)]">{{ $t('demo.mockMsg') }}</p>
                  <p class="text-xs text-[var(--td-text-color-secondary)] mt-1">
                    {{ $t('demo.interactiveBody') }}
                  </p>
                </div>
              </div>
              <div class="mt-4 flex gap-3">
                <div
                  class="px-3 py-1 rounded border border-[var(--td-brand-color)] text-[var(--td-brand-color)] text-xs font-medium"
                >
                  {{ $t('demo.notifApprove') }}
                </div>
                <div
                  class="px-3 py-1 rounded border border-[var(--td-border-level-1-color)] text-[var(--td-text-color-placeholder)] text-xs"
                >
                  {{ $t('demo.notifReject') }}
                </div>
                <div
                  class="flex-1 text-right text-xs text-[var(--td-text-color-placeholder)] italic pt-1"
                >
                  {{ $t('demo.notifReply') }}...
                </div>
              </div>
            </div>
            <t-button block theme="primary" size="large" @click="sendInteractive">
              <template #prefixIcon><ChatIcon /></template>
              {{ $t('demo.sendInteractive') }}
            </t-button>
          </div>
        </t-card>
      </t-col>
    </t-row>

    <!-- 交互响应日志 -->
    <div class="space-y-4">
      <h3 class="text-sm font-medium text-[var(--td-text-color-primary)] flex items-center gap-2">
        <CheckCircleFilledIcon class="text-[var(--td-brand-color)]" />
        {{ $t('demo.lastResponse') }}
      </h3>

      <div
        class="p-4 rounded-xl border border-[var(--td-border-level-1-color)] bg-[var(--td-bg-color-container)] min-h-[120px] flex flex-col justify-center transition-all duration-300"
      >
        <Transition name="fade-sub" mode="out-in">
          <div v-if="lastResponse" :key="lastResponse.timestamp" class="space-y-3">
            <div
              class="flex items-center justify-between border-b border-[var(--td-border-level-1-color)] pb-2 mb-2"
            >
              <span class="text-xs font-bold text-[var(--td-brand-color)]"
                >Action: {{ lastResponse.ActionIdentifier }}</span
              >
              <span
                class="text-[10px] text-[var(--td-text-color-placeholder)] flex items-center gap-1"
              >
                <TimeIcon size="12" /> {{ lastResponse.timestamp }}
              </span>
            </div>
            <div class="grid grid-cols-2 gap-y-2 text-xs">
              <div class="text-[var(--td-text-color-placeholder)]">Notification ID</div>
              <div class="text-[var(--td-text-color-primary)] font-mono">{{ lastResponse.ID }}</div>

              <div class="text-[var(--td-text-color-placeholder)]">User Text</div>
              <div class="text-[var(--td-brand-color)] font-medium">
                {{ lastResponse.UserText || '(None)' }}
              </div>
            </div>
          </div>
          <div
            v-else
            key="empty"
            class="flex flex-col items-center justify-center text-[var(--td-text-color-placeholder)] py-4"
          >
            <ChatIcon size="32" class="opacity-20 mb-2" />
            <p class="text-xs">{{ $t('demo.noResponse') }}</p>
          </div>
        </Transition>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 状态切换淡入效果 (Sub-Transition) */
.fade-sub-enter-active,
.fade-sub-leave-active {
  transition: all 0.2s ease;
}
.fade-sub-enter-from {
  opacity: 0;
  transform: translateY(5px);
}
.fade-sub-leave-to {
  opacity: 0;
  transform: translateY(-5px);
}
</style>

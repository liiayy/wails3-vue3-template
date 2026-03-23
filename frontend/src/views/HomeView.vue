<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { registerUser, fetchUserProfile } from '../api/user'
import { useSettingsStore } from '../stores/settings'
import { useWailsEvent, useAsyncAction } from '../composables'

const settings = useSettingsStore()
const { t } = useI18n()

// 【Composable 1】useWailsEvent — 自动管理事件订阅/卸载
const { data: currentTime } = useWailsEvent<string>('time')

// 【Composable 2】useAsyncAction — 注册操作
const registerForm = ref({ name: '', email: '' })
const {
  execute: doRegister,
  loading: registerLoading,
  error: registerError,
  data: registeredUser,
} = useAsyncAction(async () => {
  if (!registerForm.value.name || !registerForm.value.email) {
    throw new Error(t('common.required'))
  }
  return await registerUser(registerForm.value.name, registerForm.value.email)
})

// 【Composable 2】useAsyncAction — 查询操作
const queryId = ref<number>(1)
const {
  execute: doQuery,
  loading: queryLoading,
  error: queryError,
  data: queriedUser,
} = useAsyncAction(async () => {
  return await fetchUserProfile(queryId.value)
})
</script>

<template>
  <t-card :bordered="false" class="shadow-sm rounded-lg max-w-2xl mx-auto mt-8">
    <div class="space-y-8 flex flex-col items-center py-6">
      <div class="text-center space-y-2">
        <h2 class="text-2xl font-bold text-[var(--td-text-color-primary)]">
          {{ $t('home.title') }}
        </h2>
        <p class="text-[var(--td-text-color-secondary)]">
          {{ $t('home.subtitle') }}
        </p>
      </div>

      <!-- 状态结果面板 -->
      <div
        class="w-full bg-[var(--td-bg-color-secondarycontainer)] p-4 rounded-md flex justify-between items-center text-[var(--td-brand-color)] font-medium"
      >
        <span v-if="registeredUser">
          ✅ {{ $t('home.registerSuccess') }}: {{ registeredUser.name }} (ID:
          {{ registeredUser.id }})
        </span>
        <span v-else-if="queriedUser">
          🔍 {{ $t('home.found') }}: {{ queriedUser.name }} &lt;{{ queriedUser.email }}&gt;
        </span>
        <span v-else-if="registerError || queryError" class="text-[var(--td-error-color)]">
          ❌ {{ registerError || queryError }}
        </span>
        <span v-else>{{ $t('home.ready') }}</span>

        <div class="flex items-center gap-2">
          <span class="text-xs text-[var(--td-text-color-secondary)] uppercase">{{
            $t('home.persistTheme')
          }}</span>
          <t-radio-group
            variant="default-filled"
            :value="settings.theme"
            @change="(val: any) => settings.updateSetting('theme', val)"
          >
            <t-radio-button value="light">{{ $t('home.themeLight') }}</t-radio-button>
            <t-radio-button value="dark">{{ $t('home.themeDark') }}</t-radio-button>
          </t-radio-group>
        </div>
      </div>

      <div class="w-full grid grid-cols-2 gap-8">
        <!-- Register Section -->
        <div class="space-y-4">
          <h3 class="font-semibold text-lg">{{ $t('home.registerSection') }}</h3>
          <t-input v-model="registerForm.name" :placeholder="$t('home.namePlaceholder')" />
          <t-input v-model="registerForm.email" :placeholder="$t('home.emailPlaceholder')" />
          <t-button block theme="primary" :loading="registerLoading" @click="doRegister()">{{
            $t('home.registerBtn')
          }}</t-button>
        </div>

        <!-- Query Section -->
        <div class="space-y-4">
          <h3 class="font-semibold text-lg">{{ $t('home.lookupSection') }}</h3>
          <t-input-number v-model="queryId" :min="1" placeholder="ID" class="w-full" />
          <t-button block theme="default" :loading="queryLoading" @click="doQuery()">{{
            $t('home.fetchBtn')
          }}</t-button>
        </div>
      </div>

      <t-alert theme="info" class="mt-8 w-full">
        <template #title>{{ $t('home.eventBusTitle') }}</template>
        {{ currentTime || $t('home.eventBusWaiting') }}
      </t-alert>
    </div>
  </t-card>
</template>

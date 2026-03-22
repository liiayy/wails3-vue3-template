<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { registerUser, fetchUserProfile } from '../api/user'
import { useSettingsStore } from '../stores/settings'
import { Events } from '@wailsio/runtime'

const settings = useSettingsStore()
const registerForm = ref({ name: '', email: '' })
const queryId = ref<number>(1)
const resultText = ref("Ready to interact with Go Mega-Structure Backend 🚀")
const currentTime = ref("")

// 订阅后端事件
onMounted(() => {
  Events.On('time', (time: unknown) => {
    currentTime.value = time as string
  })
})
onUnmounted(() => { Events.Off('time') })

// 表单防抖提示
const loading = ref(false)

async function onRegister() {
  if (!registerForm.value.name || !registerForm.value.email) {
    resultText.value = 'Name and Email are required!'
    return
  }
  loading.value = true
  try {
    const user = await registerUser(registerForm.value.name, registerForm.value.email)
    resultText.value = `Success! User [${user?.name}] created with ID: ${user?.id}`
  } catch (err) {
    resultText.value = "Register failed: " + String(err)
  }
  loading.value = false
}

async function onQuery() {
  loading.value = true
  try {
    const user = await fetchUserProfile(queryId.value)
    if (user) {
      resultText.value = `Found ID(${user.id}): ${user.name} <${user.email}>`
    } else {
      resultText.value = `User ID ${queryId.value} not found (Returned null)`
    }
  } catch (err) {
    resultText.value = "Query failed: " + String(err)
  }
  loading.value = false
}
</script>

<template>
  <t-card :bordered="false" class="shadow-sm rounded-lg max-w-2xl mx-auto mt-8">
    <div class="space-y-8 flex flex-col items-center py-6">
      
      <div class="text-center space-y-2">
        <h2 class="text-2xl font-bold text-[var(--td-text-color-primary)]">
          Wails 3 Mega-Structure Demo
        </h2>
        <p class="text-[var(--td-text-color-secondary)]">
          Testing App → Binding → Service → Repository Layers
        </p>
      </div>

      <div class="w-full bg-[var(--td-bg-color-secondarycontainer)] p-4 rounded-md flex justify-between items-center text-[var(--td-brand-color)] font-medium">
        <span>{{ resultText }}</span>
        <div class="flex items-center gap-2">
          <span class="text-xs text-[var(--td-text-color-secondary)] uppercase">持久化主题:</span>
          <t-radio-group 
            variant="default-filled" 
            :value="settings.theme" 
            @change="(val: any) => settings.updateSetting('theme', val)"
          >
            <t-radio-button value="light">明亮</t-radio-button>
            <t-radio-button value="dark">黑暗</t-radio-button>
          </t-radio-group>
        </div>
      </div>

      <div class="w-full grid grid-cols-2 gap-8">
        <!-- Register Section -->
        <div class="space-y-4">
          <h3 class="font-semibold text-lg">1. Register (Write)</h3>
          <t-input v-model="registerForm.name" placeholder="Name" />
          <t-input v-model="registerForm.email" placeholder="Email" />
          <t-button block theme="primary" :loading="loading" @click="onRegister">Register User</t-button>
        </div>

        <!-- Query Section -->
        <div class="space-y-4">
          <h3 class="font-semibold text-lg">2. Lookup (Read)</h3>
          <t-input-number v-model="queryId" :min="1" placeholder="User ID" class="w-full" />
          <t-button block theme="default" :loading="loading" @click="onQuery">Fetch User</t-button>
        </div>
      </div>

      <t-alert theme="info" class="mt-8 w-full">
        <template #title>Wails Event Bus Stream (App Hook):</template>
        {{ currentTime || "Waiting for 'time' event from background goroutine..." }}
      </t-alert>
      
    </div>
  </t-card>
</template>

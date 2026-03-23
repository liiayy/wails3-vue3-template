<script setup lang="ts">
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { useSettingsStore } from '../stores/settings'
import TitleBar from '../components/TitleBar.vue'
import {
  PaletteIcon,
  TranslateIcon,
  ViewListIcon,
  NotificationIcon,
  UserIcon,
  SettingIcon,
  WifiIcon,
  SecuredIcon,
  HistoryIcon,
  CheckCircleFilledIcon,
} from 'tdesign-icons-vue-next'

const settings = useSettingsStore()
const route = useRoute()
const isStandalone = route.path.startsWith('/standalone')

// 当前激活的分类
const activeCategory = ref('personalization')

const categories = [
  { id: 'personalization', label: 'settings.personalization', icon: PaletteIcon },
  { id: 'notifications', label: 'settings.notifications', icon: NotificationIcon },
]
</script>

<template>
  <div class="h-full flex flex-col bg-[var(--td-bg-color-container)] overflow-hidden">
    <!-- 无边框窗口标题栏 -->
    <TitleBar v-if="isStandalone" :title="$t('settings.title')" no-minimize no-maximize />

    <div class="flex-1 flex overflow-hidden">
      <!-- 第二栏：分类导航 (Middle Column) -->
      <aside
        class="w-44 border-r border-[var(--td-border-level-1-color)] bg-[var(--td-bg-color-container)] shrink-0 flex flex-col"
      >
        <div class="p-2">
          <nav class="space-y-1">
            <div
              v-for="cat in categories"
              :key="cat.id"
              @click="activeCategory = cat.id"
              class="flex items-center gap-2.5 px-3 py-2 rounded-lg cursor-pointer transition-all duration-200"
              :class="
                activeCategory === cat.id
                  ? 'bg-[var(--td-brand-color-light)] text-[var(--td-brand-color)] font-medium shadow-sm'
                  : 'text-[var(--td-text-color-secondary)] hover:bg-[var(--td-bg-color-secondarycontainer)]'
              "
            >
              <component :is="cat.icon" size="16" />
              <span class="text-sm">
                {{ $t(cat.label) }}
              </span>
            </div>
          </nav>
        </div>
      </aside>

      <!-- 第三栏：主内容区 (Content Area) -->
      <main class="flex-1 overflow-auto bg-[var(--td-bg-color-page)] relative">
        <div class="max-w-3xl p-6 space-y-6 animate-fade-in">
          <!-- 头部标题 -->
          <div class="space-y-1">
            <h1 class="text-xl font-bold text-[var(--td-text-color-primary)] tracking-tight">
              {{ $t(categories.find((c) => c.id === activeCategory)?.label || '') }}
            </h1>
            <p class="text-[var(--td-text-color-secondary)] text-sm">
              {{ $t('settings.subtitle') }}
            </p>
          </div>

          <!-- 根据分类显示内容 -->
          <div v-if="activeCategory === 'personalization'" class="space-y-6">
            <section class="space-y-3">
              <div
                class="text-[11px] font-bold text-[var(--td-text-color-placeholder)] uppercase tracking-wider"
              >
                {{ $t('settings.themeMode') }}
              </div>
              <div class="grid grid-cols-3 gap-3">
                <!-- 明亮模式预览卡片 -->
                <div
                  class="group relative cursor-pointer"
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
                <div
                  class="group relative cursor-pointer"
                  @click="settings.updateSetting('theme', 'dark')"
                >
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
                    <div
                      class="h-5 w-full bg-gray-800 border border-gray-700 rounded shadow-sm"
                    ></div>
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
                <div
                  class="group relative cursor-pointer"
                  @click="settings.updateSetting('theme', 'auto')"
                >
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

            <div class="grid grid-cols-2 gap-6">
              <!-- 语言设置简版 -->
              <t-card
                :bordered="false"
                class="bg-[var(--td-bg-color-container)] rounded-xl p-0.5 shadow-sm"
              >
                <div class="space-y-4">
                  <div
                    class="text-[10px] font-bold text-[var(--td-text-color-placeholder)] uppercase"
                  >
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
                    <div
                      class="text-[10px] font-bold text-[var(--td-text-color-placeholder)] uppercase"
                    >
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
            </div>
          </div>

          <!-- 其他占位 -->
          <div v-else class="flex flex-col items-center justify-center py-10 opacity-30 grayscale">
            <component :is="categories.find((c) => c.id === activeCategory)?.icon" size="60" />
            <p class="mt-4 text-lg font-medium">{{ $t('settings.comingSoon') }}</p>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

:deep(.t-card) {
  transition: transform 0.2s;
}
:deep(.t-card:hover) {
  transform: translateY(-2px);
}
</style>

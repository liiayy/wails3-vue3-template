import { defineStore } from 'pinia'
import { NotificationBinding, SettingBinding } from '#/myapp2/internal/binding'
import { Events } from '@wailsio/runtime'
import i18n from '@/locales'

interface SettingsState {
  theme: 'light' | 'dark' | 'auto'
  language: string
  isSidebarCollapsed: boolean
}

// 跨窗口同步事件名称
const SYNC_EVENT = 'app:settings-changed'

export const useSettingsStore = defineStore('settings', {
  state: (): SettingsState => ({
    theme: 'auto',
    language: 'zh-CN',
    isSidebarCollapsed: true,
  }),

  actions: {
    async init() {
      try {
        const remoteSettings = await SettingBinding.GetAll()

        if (remoteSettings) {
          if (remoteSettings.theme) this.theme = remoteSettings.theme as any
          if (remoteSettings.language) this.language = remoteSettings.language
          if (remoteSettings.isSidebarCollapsed)
            this.isSidebarCollapsed = remoteSettings.isSidebarCollapsed === 'true'
        }

        this.applyTheme()
        this.applyLanguage()

        // 注册跨窗口同步监听器
        Events.On(SYNC_EVENT, (ev: any) => {
          const { key, value } = ev.data
          if ((this.$state as any)[key] === value) return

          ;(this.$state as any)[key] = value

          if (key === 'theme') this.applyTheme()
          if (key === 'language') this.applyLanguage()
        })

        // 监听系统主题变化
        window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', () => {
          if (this.theme === 'auto') {
            this.applyTheme()
          }
        })
      } catch (err) {
        console.error('[Settings] 初始化失败:', err)
      }
    },

    async updateSetting<K extends keyof SettingsState>(key: K, value: SettingsState[K]) {
      this.$state[key] = value

      try {
        await SettingBinding.Save(key, String(value))

        if (key === 'theme') {
          this.applyTheme()
        }
        if (key === 'language') {
          this.applyLanguage()
        }

        // 广播变更
        Events.Emit(SYNC_EVENT, { key, value })
      } catch (err) {
        console.error(`[Settings] 同步项目 ${key} 失败:`, err)
      }
    },

    applyTheme() {
      const doc = document.documentElement
      let targetTheme = this.theme

      if (targetTheme === 'auto') {
        const isDark = window.matchMedia('(prefers-color-scheme: dark)').matches
        targetTheme = isDark ? 'dark' : 'light'
      }

      if (targetTheme === 'dark') {
        doc.setAttribute('theme-mode', 'dark')
        doc.classList.add('dark')
      } else {
        doc.removeAttribute('theme-mode')
        doc.classList.remove('dark')
      }
    },

    applyLanguage() {
      // 1. 更新前端 Vue-i18n
      if (i18n.global.locale) {
        ;(i18n.global.locale as any).value = this.language
      }
      
      // 2. 同步给后端的国际化服务
      NotificationBinding.SetLanguage(this.language).catch((err) => {
        console.error('[Settings] 同步后端语言失败:', err)
      })
    },
  },
})

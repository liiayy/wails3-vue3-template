import { defineStore } from 'pinia'
import { SettingBinding } from '../../bindings/myapp2/internal/binding'
import { Events } from '@wailsio/runtime'
import i18n from '../locales'

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
    /**
     * 从 Go 后端 (SQLite) 初始化加载所有设置
     * 并开启 Wails 事件监听，实现多窗口实时同步
     */
    async init() {
      try {
        const remoteSettings = await SettingBinding.GetAll()

        if (remoteSettings) {
          if (remoteSettings.theme) this.theme = remoteSettings.theme as any
          if (remoteSettings.language) this.language = remoteSettings.language
          if (remoteSettings.isSidebarCollapsed)
            this.isSidebarCollapsed = remoteSettings.isSidebarCollapsed === 'true'
        }

        console.log('[Settings] 初始化加载完成:', this.$state)
        this.applyTheme()
        this.applyLanguage()

        // 注册跨窗口同步监听器
        Events.On(SYNC_EVENT, (ev: any) => {
          const { key, value } = ev.data
          // 如果本地状态已是最新，说明是自己发出的或重复事件，跳过
          if ((this.$state as any)[key] === value) return

          console.log(`[Settings] 收到跨窗口同步事件: ${key} -> ${value}`)
          ;(this.$state as any)[key] = value

          // 执行部分需要立即生效的副作用
          if (key === 'theme') this.applyTheme()
          if (key === 'language') this.applyLanguage()
        })

        // 监听系统主题变化
        window
          .matchMedia('(prefers-color-scheme: dark)')
          .addEventListener('change', () => {
            if (this.theme === 'auto') {
              this.applyTheme()
            }
          })
      } catch (err) {
        console.error('[Settings] 初始化失败:', err)
      }
    },

    /**
     * 修改设置并同步至后端持久化，同时广播事件给其他窗口
     */
    async updateSetting<K extends keyof SettingsState>(key: K, value: SettingsState[K]) {
      // 1. 更新本地状态
      this.$state[key] = value

      // 2. 执行持久化同步 (DB)
      try {
        await SettingBinding.Save(key, String(value))

        // 3. 应用本地副作用
        if (key === 'theme') {
          this.applyTheme()
        }
        if (key === 'language') {
          this.applyLanguage()
        }

        // 4. 重头戏：通过 Wails 事件总线广播变更，让其他窗口实时更新
        // 注意：@wailsio/runtime 的 Emit 签名可能是 (name, data)
        Events.Emit(SYNC_EVENT, { key, value })
      } catch (err) {
        console.error(`[Settings] 同步项目 ${key} 失败:`, err)
      }
    },

    /**
     * 将主题应用到 DOM (用于 Tailwind/TDesign)
     */
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

    /**
     * 更新 i18n
     */
    applyLanguage() {
      ;(i18n.global.locale as any).value = this.language
    },
  },
})

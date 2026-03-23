import { defineStore } from 'pinia'
import { SettingBinding } from '../../bindings/myapp2/internal/binding'
import i18n from '../locales'

interface SettingsState {
  theme: 'light' | 'dark' | 'auto'
  language: string
  isSidebarCollapsed: boolean
}

export const useSettingsStore = defineStore('settings', {
  state: (): SettingsState => ({
    theme: 'light',
    language: 'zh-CN',
    isSidebarCollapsed: false,
  }),

  actions: {
    /**
     * 从 Go 后端 (SQLite) 初始化加载所有设置
     */
    async init() {
      try {
        const remoteSettings = await SettingBinding.GetAll()
        
        if (remoteSettings) {
          if (remoteSettings.theme) this.theme = remoteSettings.theme as any
          if (remoteSettings.language) this.language = remoteSettings.language
          if (remoteSettings.isSidebarCollapsed) this.isSidebarCollapsed = remoteSettings.isSidebarCollapsed === 'true'
        }
        
        console.log('[Settings] 初始化加载完成:', this.$state)
        this.applyTheme()
        this.applyLanguage()
      } catch (err) {
        console.error('[Settings] 初始化失败:', err)
      }
    },

    /**
     * 修改设置并同步至后端持久化
     */
    async updateSetting<K extends keyof SettingsState>(key: K, value: SettingsState[K]) {
      this.$state[key] = value
      
      // 执行持久化同步
      try {
        await SettingBinding.Save(key, String(value))
        
        if (key === 'theme') {
          this.applyTheme()
        }
        if (key === 'language') {
          this.applyLanguage()
        }
      } catch (err) {
        console.error(`[Settings] 同步项目 ${key} 失败:`, err)
      }
    },

    /**
     * 将主题应用到 DOM (用于 Tailwind/TDesign)
     */
    applyTheme() {
      const doc = document.documentElement
      if (this.theme === 'dark') {
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
      (i18n.global.locale as any).value = this.language
    }
  }
})

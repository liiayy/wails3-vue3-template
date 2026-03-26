import { defineStore } from 'pinia'
import {
  NotificationBinding,
  SettingBinding,
  SystemBinding,
} from '#/myapp2/internal/binding'
import { Events } from '@wailsio/runtime'
import { handleResult } from '@/api/base'

interface SettingsState {
  theme: 'light' | 'dark' | 'auto'
  language: string
  isSidebarCollapsed: boolean
  isAutostart: boolean
  zoom: number
}

// 跨窗口同步事件名称
const SYNC_EVENT = 'app:settings-changed'

export const useSettingsStore = defineStore('settings', {
  state: (): SettingsState => ({
    theme: 'auto',
    language: 'zh-CN',
    isSidebarCollapsed: true,
    isAutostart: false,
    zoom: 100,
  }),

  actions: {
    /**
     * 从 Go 后端 (SQLite) 初始化加载所有设置
     * 并开启 Wails 事件监听，实现多窗口实时同步
     */
    async init() {
      const remoteRes = await SettingBinding.GetAll()
      const autostartRes = await SystemBinding.IsAutostartEnabled()

      // 静默获取，因为初始化不想弹窗
      const remoteSettings = handleResult<Record<string, string>>(remoteRes, true) || {}
      const autostart = handleResult<boolean>(autostartRes, true)

      if (remoteSettings.theme) this.theme = remoteSettings.theme as any
      if (remoteSettings.language) this.language = remoteSettings.language
      if (remoteSettings.isSidebarCollapsed)
        this.isSidebarCollapsed = remoteSettings.isSidebarCollapsed === 'true'
      if (remoteSettings.zoom)
        this.zoom = Number(remoteSettings.zoom) || 100

      this.isAutostart = autostart

        // 注册跨窗口同步监听器
        Events.On(SYNC_EVENT, (ev: any) => {
          const { key, value } = ev.data
          if ((this.$state as any)[key] === value) return
          ;(this.$state as any)[key] = value
        })
    },

    async updateSetting<K extends keyof SettingsState>(key: K, value: SettingsState[K]) {
      this.$state[key] = value

      if (key === 'isAutostart') {
        const res = await SystemBinding.SetAutostart(value as boolean)
        handleResult(res)
      } else {
        const res = await SettingBinding.Save(key, String(value))
        handleResult(res)
      }

      if (key === 'language') {
        // 同步给后端的国际化服务 (逻辑同步，非 UI)
        const res = await NotificationBinding.SetLanguage(value as string)
        handleResult(res)
      }

      // 广播变更
      Events.Emit(SYNC_EVENT, { key, value })
    },

  },
})

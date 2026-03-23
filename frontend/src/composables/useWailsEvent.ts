import { onMounted, onUnmounted, ref, type Ref } from 'vue'
import { Events } from '@wailsio/runtime'

/**
 * useWailsEvent — 自动管理 Wails 事件的订阅与卸载
 *
 * 在组件 mount 时注册事件监听器，在 unmount 时自动销毁，杜绝内存泄漏。
 * 支持泛型 data payload 类型推导。
 *
 * @example
 * ```ts
 * // 基础用法
 * const { data } = useWailsEvent<string>('time')
 *
 * // 自定义回调
 * useWailsEvent<number>('progress', (payload) => {
 *   console.log('当前进度:', payload)
 * })
 * ```
 */
export function useWailsEvent<T = any>(
  eventName: string,
  callback?: (data: T) => void
) {
  const data: Ref<T | null> = ref(null)
  let unsubscribe: (() => void) | null = null

  onMounted(() => {
    unsubscribe = Events.On(eventName, (ev: any) => {
      data.value = ev.data as T
      callback?.(ev.data as T)
    })
  })

  onUnmounted(() => {
    unsubscribe?.()
    unsubscribe = null
  })

  return { data }
}

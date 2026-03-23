import { ref, watch, onUnmounted, type Ref } from 'vue'

/**
 * useDebounce — 防抖值封装
 *
 * 常用于搜索输入框，在用户停止输入后延迟触发后端查询。
 *
 * @example
 * ```ts
 * const keyword = ref('')
 * const debouncedKeyword = useDebounce(keyword, 300)
 *
 * watch(debouncedKeyword, (val) => {
 *   // 仅在用户停止输入 300ms 后才触发
 *   searchApi(val)
 * })
 * ```
 */
export function useDebounce<T>(source: Ref<T>, delayMs = 300): Ref<T> {
  const debounced: Ref<T> = ref(source.value) as Ref<T>
  let timer: ReturnType<typeof setTimeout> | null = null

  const stop = watch(source, (newVal: T) => {
    if (timer) clearTimeout(timer)
    timer = setTimeout(() => {
      debounced.value = newVal
    }, delayMs)
  })

  onUnmounted(() => {
    if (timer) clearTimeout(timer)
    stop()
  })

  return debounced
}

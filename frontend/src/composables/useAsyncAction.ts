import { ref, type Ref } from 'vue'

/**
 * useAsyncAction — 通用异步操作封装
 *
 * 将任意异步函数包裹为带有响应式 loading / error / data 状态的安全操作。
 * 自动捕获异常并填充 error，避免在每个组件中重复编写 try/catch。
 *
 * @example
 * ```ts
 * const { execute, loading, error, data } = useAsyncAction(
 *   (id: number) => fetchUserProfile(id)
 * )
 *
 * // 在模板中
 * <t-button :loading="loading" @click="execute(1)">查询</t-button>
 * <p v-if="error">{{ error }}</p>
 * <p v-if="data">{{ data.name }}</p>
 * ```
 */
export function useAsyncAction<T, Args extends any[] = any[]>(fn: (...args: Args) => Promise<T>) {
  const loading: Ref<boolean> = ref(false)
  const error: Ref<string | null> = ref(null)
  const data: Ref<T | null> = ref(null) as Ref<T | null>

  async function execute(...args: Args): Promise<T | null> {
    loading.value = true
    error.value = null

    try {
      const result = await fn(...args)
      data.value = result
      return result
    } catch (err: any) {
      error.value = err?.message || String(err)
      data.value = null
      return null
    } finally {
      loading.value = false
    }
  }

  /** 重置所有状态 */
  function reset() {
    loading.value = false
    error.value = null
    data.value = null
  }

  return {
    execute,
    loading,
    error,
    data,
    reset,
  }
}

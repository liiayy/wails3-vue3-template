import { MessagePlugin } from 'tdesign-vue-next'

/** 
 * 与后端 internal/binding/response.go 中的 Result 结构对应 
 */
export interface Result<T = any> {
  success: boolean
  data?: T
  error?: AppError
}

/** 
 * 与后端 internal/domain/errors.go 中的 AppError 结构对应 
 */
export interface AppError {
  code: string
  message: string
}

/**
 * 统一处理后端返回的 Result 包装。
 * 如果业务失败，自动弹出错误提示并抛出异常。
 * 
 * @param result 后端返回的原始 Result 对象
 * @param silent 是否静默处理（不弹出 Message 提示）
 */
export function handleResult<T>(result: any, silent = false): T {
  const res = result as Result<T>
  
  if (!res) {
    throw new Error('Backend returned empty response')
  }

  if (res.success) {
    return res.data as T
  }

  // 业务逻辑失败
  const errMsg = res.error?.message || 'Unknown business error'
  const errCode = res.error?.code || 'INTERNAL_ERROR'

  if (!silent) {
    MessagePlugin.error(`[${errCode}] ${errMsg}`)
  }

  const err = new Error(errMsg)
  ;(err as any).code = errCode
  throw err
}

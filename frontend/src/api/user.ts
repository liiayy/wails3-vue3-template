import { UserBinding } from '../../bindings/myapp2/internal/binding'
import type { User, UserListResult } from '../../bindings/myapp2/internal/domain'

export type { User, UserListResult }

/** 注册新用户 */
export async function registerUser(name: string, email: string): Promise<User | null> {
  try {
    return await UserBinding.Register(name, email)
  } catch (err) {
    console.error('[API] registerUser failed:', err)
    throw err
  }
}

/** 查询单个用户 */
export async function fetchUserProfile(id: number): Promise<User | null> {
  try {
    return await UserBinding.GetProfile(id)
  } catch (err) {
    console.error('[API] fetchUserProfile failed:', err)
    throw err
  }
}

/** 分页查询用户列表 */
export async function listUsers(keyword: string, page: number, pageSize: number): Promise<{ items: User[]; total: number }> {
  try {
    const res = await UserBinding.List(keyword, page, pageSize)
    return {
      items: (res?.items || []).filter((u): u is User => u !== null),
      total: res?.total || 0,
    }
  } catch (err) {
    console.error('[API] listUsers failed:', err)
    throw err
  }
}

/** 更新用户 */
export async function updateUser(id: number, name: string, email: string): Promise<User | null> {
  try {
    return await UserBinding.Update(id, name, email)
  } catch (err) {
    console.error('[API] updateUser failed:', err)
    throw err
  }
}

/** 删除用户 */
export async function deleteUser(id: number): Promise<void> {
  try {
    await UserBinding.Delete(id)
  } catch (err) {
    console.error('[API] deleteUser failed:', err)
    throw err
  }
}

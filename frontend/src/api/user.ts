import { UserBinding } from '#/myapp2/internal/binding'
import type { User, UserListResult } from '#/myapp2/internal/domain'
import { handleResult } from './base'

export type { User, UserListResult }

/** 注册新用户 */
export async function registerUser(name: string, email: string): Promise<User> {
  const res = await UserBinding.Register({ name, email })
  return handleResult<User>(res)
}

/** 查询单个用户 */
export async function fetchUserProfile(id: number): Promise<User> {
  const res = await UserBinding.GetProfile(id)
  return handleResult<User>(res)
}

/** 分页查询用户列表 */
export async function listUsers(
  keyword: string,
  page: number,
  pageSize: number,
): Promise<{ items: User[]; total: number }> {
  const res = await UserBinding.List(keyword, page, pageSize)
  const data = handleResult<UserListResult>(res)
  return {
    items: data.items.filter((u): u is User => u !== null),
    total: data.total,
  }
}

/** 更新用户 */
export async function updateUser(id: number, name: string, email: string): Promise<User> {
  const res = await UserBinding.Update({ id, name, email })
  return handleResult<User>(res)
}

/** 删除用户 */
export async function deleteUser(id: number): Promise<void> {
  const res = await UserBinding.Delete(id)
  handleResult<void>(res)
}

/** 导出所有用户 */
export async function exportUsers(): Promise<void> {
  const res = await UserBinding.Export()
  handleResult<void>(res)
}

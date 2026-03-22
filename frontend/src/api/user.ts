import { UserBinding } from '../../bindings/myapp2/internal/binding'
import type { User } from '../../bindings/myapp2/internal/domain/models'

// 获取个人信息
export async function fetchUserProfile(id: number): Promise<User | null> {
  try {
    return await UserBinding.GetProfile(id)
  } catch (error) {
    console.error("[API Error] fetchUserProfile:", error)
    throw error
  }
}

// 注册新用户
export async function registerUser(name: string, email: string): Promise<User | null> {
  try {
    return await UserBinding.Register(name, email)
  } catch (error) {
    console.error("[API Error] registerUser:", error)
    throw error
  }
}

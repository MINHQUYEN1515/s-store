import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { UserEntity } from '@domain/entities/user.entity'
import { dependencies } from '@app/providers/dependency.provider'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<UserEntity | null>(null)
  const loading = ref(false)
  const errorMessage = ref<string | null>(null)

  async function login(email: string, password: string) {
    loading.value = true
    errorMessage.value = null

    try {
      user.value = await dependencies.loginUseCase.execute(email, password)
    } catch (error) {
      errorMessage.value = error instanceof Error ? error.message : 'Đăng nhập thất bại'
      throw error
    } finally {
      loading.value = false
    }
  }

  async function logout() {
    loading.value = true
    errorMessage.value = null

    try {
      await dependencies.logoutUseCase.execute()
      user.value = null
    } catch (error) {
      errorMessage.value = error instanceof Error ? error.message : 'Đăng xuất thất bại'
      throw error
    } finally {
      loading.value = false
    }
  }

  return {
    user,
    loading,
    errorMessage,
    login,
    logout,
  }
})

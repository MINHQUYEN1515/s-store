import { storeToRefs } from 'pinia'
import { useAuthStore } from '@presentation/stores/auth.store'

export function useAuth() {
  const authStore = useAuthStore()
  const { user, loading, errorMessage } = storeToRefs(authStore)

  return {
    user,
    loading,
    errorMessage,
    login: authStore.login,
    logout: authStore.logout,
  }
}

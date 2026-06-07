import { apiClient } from './api-client'
import type { LoginResponseModel } from '@data/models/login-response.model'
import type { UserModel } from '@data/models/user.model'

export const authApi = {
  async login(email: string, password: string): Promise<LoginResponseModel> {
    const response = await apiClient.post<LoginResponseModel>('/auth/login', {
      email,
      password,
    })

    return response.data
  },

  async getProfile(): Promise<UserModel> {
    const response = await apiClient.get<UserModel>('/auth/profile')
    return response.data
  },

  async logout(): Promise<void> {
    await apiClient.post('/auth/logout')
  },
}

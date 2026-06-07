import type { UserRepository } from '@domain/repositories/user.repository'
import type { UserEntity } from '@domain/entities/user.entity'
import { authApi } from '@data/datasources/remote/auth.api'
import { mapUserModelToEntity } from '@data/mappers/user.mapper'
import { tokenStorage } from '@data/datasources/local/token.storage'

export class UserRepositoryImpl implements UserRepository {
  async login(email: string, password: string): Promise<UserEntity> {
    const response = await authApi.login(email, password)

    tokenStorage.setTokens(response.access_token, response.refresh_token)

    return mapUserModelToEntity(response.user)
  }

  async logout(): Promise<void> {
    await authApi.logout()
    tokenStorage.clearTokens()
  }

  async getProfile(): Promise<UserEntity> {
    const response = await authApi.getProfile()
    return mapUserModelToEntity(response)
  }
}

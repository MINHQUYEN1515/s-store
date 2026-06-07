import type { UserEntity } from '@domain/entities/user.entity'

export interface UserRepository {
  login(email: string, password: string): Promise<UserEntity>
  logout(): Promise<void>
  getProfile(): Promise<UserEntity>
}

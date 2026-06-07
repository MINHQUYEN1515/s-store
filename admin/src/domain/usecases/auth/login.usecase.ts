import type { UserRepository } from '@domain/repositories/user.repository'
import type { UserEntity } from '@domain/entities/user.entity'

export class LoginUseCase {
  private readonly userRepository: UserRepository

  constructor(userRepository: UserRepository) {
    this.userRepository = userRepository
  }

  async execute(email: string, password: string): Promise<UserEntity> {
    if (!email) {
      throw new Error('Email không được để trống')
    }

    if (!password) {
      throw new Error('Mật khẩu không được để trống')
    }

    return this.userRepository.login(email, password)
  }
}

import type { UserRepository } from '@domain/repositories/user.repository'

export class LogoutUseCase {
  private readonly userRepository: UserRepository

  constructor(userRepository: UserRepository) {
    this.userRepository = userRepository
  }

  async execute(): Promise<void> {
    return this.userRepository.logout()
  }
}

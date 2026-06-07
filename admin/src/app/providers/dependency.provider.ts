import { UserRepositoryImpl } from '@data/repositories/user.repository.impl'
import { LoginUseCase } from '@domain/usecases/auth/login.usecase'
import { LogoutUseCase } from '@domain/usecases/auth/logout.usecase'

const userRepository = new UserRepositoryImpl()

export const dependencies = {
  userRepository,
  loginUseCase: new LoginUseCase(userRepository),
  logoutUseCase: new LogoutUseCase(userRepository),
}

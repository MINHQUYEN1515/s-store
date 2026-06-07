import type { UserModel } from './user.model'

export interface LoginResponseModel {
  access_token: string
  refresh_token: string
  user: UserModel
}

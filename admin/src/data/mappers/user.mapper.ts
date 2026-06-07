import type { UserModel } from '@data/models/user.model'
import type { UserEntity } from '@domain/entities/user.entity'

export function mapUserModelToEntity(model: UserModel): UserEntity {
  return {
    id: model.id,
    name: model.name,
    email: model.email,
    avatarUrl: model.avatar_url,
  }
}

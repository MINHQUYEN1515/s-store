import { env } from './env'

export const appConfig = {
  name: env.appName,
  apiTimeout: 30000,
  tokenKey: 'access_token',
  refreshTokenKey: 'refresh_token',
}

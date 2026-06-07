import { appConfig } from '@app/config/app.config'

export const tokenStorage = {
  getAccessToken(): string | null {
    return localStorage.getItem(appConfig.tokenKey)
  },

  getRefreshToken(): string | null {
    return localStorage.getItem(appConfig.refreshTokenKey)
  },

  setTokens(accessToken: string, refreshToken: string): void {
    localStorage.setItem(appConfig.tokenKey, accessToken)
    localStorage.setItem(appConfig.refreshTokenKey, refreshToken)
  },

  clearTokens(): void {
    localStorage.removeItem(appConfig.tokenKey)
    localStorage.removeItem(appConfig.refreshTokenKey)
  },
}

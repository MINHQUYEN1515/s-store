import axios from 'axios'
import { env } from '@app/config/env'
import { appConfig } from '@app/config/app.config'
import { tokenStorage } from '@data/datasources/local/token.storage'

export const apiClient = axios.create({
  baseURL: env.apiBaseUrl,
  timeout: appConfig.apiTimeout,
  headers: {
    'Content-Type': 'application/json',
  },
})

apiClient.interceptors.request.use((config) => {
  const token = tokenStorage.getAccessToken()

  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }

  return config
})

apiClient.interceptors.response.use(
  (response) => response,
  (error) => Promise.reject(error),
)

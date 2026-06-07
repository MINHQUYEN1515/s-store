import type { App } from 'vue'
import { router } from '@app/router'

export function setupRouter(app: App): void {
  app.use(router)
}

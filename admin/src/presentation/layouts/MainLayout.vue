<template>
  <div class="main-layout">
    <header class="main-layout__header">
      <div class="main-layout__brand">
        <h1>{{ appName }}</h1>
      </div>
      <nav class="main-layout__nav">
        <slot name="nav" />
        <BaseButton v-if="showLogout" type="button" @click="onLogout">
          Đăng xuất
        </BaseButton>
      </nav>
    </header>
    <main class="main-layout__content">
      <slot />
    </main>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { env } from '@app/config/env'
import { APP_ROUTES } from '@core/constants/app.constants'
import BaseButton from '@presentation/components/base/BaseButton.vue'
import { useAuth } from '@presentation/composables/useAuth'

withDefaults(
  defineProps<{
    showLogout?: boolean
  }>(),
  {
    showLogout: true,
  },
)

const router = useRouter()
const { logout } = useAuth()
const appName = env.appName

async function onLogout() {
  await logout()
  router.push(APP_ROUTES.LOGIN)
}
</script>

<style scoped>
.main-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.main-layout__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem 1.5rem;
  background-color: var(--color-surface);
  border-bottom: 1px solid var(--color-border);
}

.main-layout__brand h1 {
  font-size: 1.125rem;
  font-weight: 600;
}

.main-layout__nav {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.main-layout__content {
  flex: 1;
  padding: 1.5rem;
}
</style>

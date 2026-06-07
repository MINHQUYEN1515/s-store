<template>
  <AuthLayout>
    <div class="login-page">
      <h2 class="login-page__title">Đăng nhập</h2>
      <p class="login-page__subtitle">Chào mừng bạn đến với {{ appName }}</p>

      <form class="login-page__form" @submit.prevent="onSubmit">
        <BaseInput
          v-model="email"
          label="Email"
          type="email"
          placeholder="Nhập email"
        />
        <BaseInput
          v-model="password"
          label="Mật khẩu"
          type="password"
          placeholder="Nhập mật khẩu"
        />

        <p v-if="authStore.errorMessage" class="login-page__error">
          {{ authStore.errorMessage }}
        </p>

        <BaseButton type="submit" :loading="authStore.loading" block>
          {{ authStore.loading ? 'Đang đăng nhập...' : 'Đăng nhập' }}
        </BaseButton>
      </form>
    </div>
  </AuthLayout>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { env } from '@app/config/env'
import { APP_ROUTES } from '@core/constants/app.constants'
import AuthLayout from '@presentation/layouts/AuthLayout.vue'
import BaseButton from '@presentation/components/base/BaseButton.vue'
import BaseInput from '@presentation/components/base/BaseInput.vue'
import { useAuthStore } from '@presentation/stores/auth.store'

const router = useRouter()
const authStore = useAuthStore()
const appName = env.appName

const email = ref('')
const password = ref('')

async function onSubmit() {
  try {
    await authStore.login(email.value, password.value)
    router.push(APP_ROUTES.HOME)
  } catch {
    // Error handled in store
  }
}
</script>

<style scoped>
.login-page__title {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 0.25rem;
}

.login-page__subtitle {
  font-size: 0.875rem;
  color: var(--color-text-muted);
  margin-bottom: 1.5rem;
}

.login-page__form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.login-page__error {
  font-size: 0.875rem;
  color: var(--color-danger);
}
</style>

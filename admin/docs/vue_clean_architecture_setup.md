# Setup Base Code VueJS theo Clean Architecture

Tài liệu này dùng để khởi tạo một dự án VueJS có cấu trúc rõ ràng, dễ mở rộng, dễ test và phù hợp cho dự án thực tế như admin dashboard, CRM, CMS, web quản lý hoặc frontend cho backend API.

## 1. Mục tiêu kiến trúc

Clean Architecture giúp tách code thành nhiều lớp rõ ràng:

```txt
UI không phụ thuộc trực tiếp API
Business logic không phụ thuộc framework
Repository che giấu chi tiết gọi API/local storage
Dễ test use case, repository, store
Dễ đổi API client hoặc framework sau này
```

Luồng xử lý chuẩn:

```txt
View / Component
    ↓
Store / ViewModel
    ↓
UseCase
    ↓
Repository Interface
    ↓
Repository Implementation
    ↓
API / Local Storage / External Service
```

## 2. Tạo project VueJS với Vite

```bash
npm create vite@latest my-vue-clean-app -- --template vue-ts
cd my-vue-clean-app
npm install
```

Chạy project:

```bash
npm run dev
```

## 3. Cài package cần thiết

```bash
npm install vue-router pinia axios zod
npm install -D eslint prettier eslint-config-prettier eslint-plugin-vue typescript
```

Ý nghĩa package:

| Package | Dùng để làm gì |
|---|---|
| vue-router | Điều hướng page |
| pinia | State management |
| axios | Gọi API |
| zod | Validate dữ liệu API/form |
| eslint | Check lỗi code |
| prettier | Format code |
| typescript | Type safety |

## 4. Cấu trúc thư mục đề xuất

```txt
src/
├── app/
│   ├── main.ts
│   ├── App.vue
│   ├── router/
│   │   └── index.ts
│   ├── providers/
│   │   ├── pinia.provider.ts
│   │   └── router.provider.ts
│   └── config/
│       ├── env.ts
│       └── app.config.ts
│
├── core/
│   ├── constants/
│   │   └── app.constants.ts
│   ├── errors/
│   │   ├── app-error.ts
│   │   └── api-error.ts
│   ├── types/
│   │   ├── nullable.ts
│   │   └── result.ts
│   └── utils/
│       ├── date.util.ts
│       └── string.util.ts
│
├── domain/
│   ├── entities/
│   │   └── user.entity.ts
│   ├── repositories/
│   │   └── user.repository.ts
│   └── usecases/
│       └── auth/
│           ├── login.usecase.ts
│           └── logout.usecase.ts
│
├── data/
│   ├── datasources/
│   │   ├── remote/
│   │   │   ├── api-client.ts
│   │   │   └── auth.api.ts
│   │   └── local/
│   │       └── token.storage.ts
│   ├── models/
│   │   ├── user.model.ts
│   │   └── login-response.model.ts
│   ├── mappers/
│   │   └── user.mapper.ts
│   └── repositories/
│       └── user.repository.impl.ts
│
├── presentation/
│   ├── components/
│   │   ├── base/
│   │   │   ├── BaseButton.vue
│   │   │   └── BaseInput.vue
│   │   └── common/
│   │       └── AppLoading.vue
│   ├── layouts/
│   │   ├── AuthLayout.vue
│   │   └── MainLayout.vue
│   ├── pages/
│   │   ├── auth/
│   │   │   └── LoginPage.vue
│   │   └── home/
│   │       └── HomePage.vue
│   ├── stores/
│   │   └── auth.store.ts
│   └── composables/
│       └── useAuth.ts
│
├── shared/
│   ├── assets/
│   ├── styles/
│   │   ├── main.css
│   │   └── variables.css
│   └── plugins/
│       └── axios.plugin.ts
│
└── vite-env.d.ts
```

## 5. Ý nghĩa từng layer

### 5.1 app

Chứa phần khởi động ứng dụng:

```txt
main.ts
router
provider
config app
```

Không viết business logic ở đây.

### 5.2 core

Chứa code dùng chung toàn app:

```txt
constants
errors
types
utils
helper
```

Core không được phụ thuộc vào `data`, `presentation` hoặc Vue component.

### 5.3 domain

Đây là phần quan trọng nhất của Clean Architecture.

Domain chứa:

```txt
Entity
Repository interface
UseCase
Business rule
```

Domain không biết API là gì, Axios là gì, Vue là gì.

### 5.4 data

Data layer chịu trách nhiệm lấy dữ liệu:

```txt
Gọi API
Đọc localStorage
Map API model sang entity
Implement repository interface
```

### 5.5 presentation

Presentation là phần UI:

```txt
Page
Component
Store Pinia
Composable
Layout
```

Presentation gọi use case, không gọi API trực tiếp.

## 6. Cấu hình alias path

File `vite.config.ts`:

```ts
import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
      '@app': fileURLToPath(new URL('./src/app', import.meta.url)),
      '@core': fileURLToPath(new URL('./src/core', import.meta.url)),
      '@domain': fileURLToPath(new URL('./src/domain', import.meta.url)),
      '@data': fileURLToPath(new URL('./src/data', import.meta.url)),
      '@presentation': fileURLToPath(new URL('./src/presentation', import.meta.url)),
      '@shared': fileURLToPath(new URL('./src/shared', import.meta.url)),
    },
  },
})
```

File `tsconfig.app.json`:

```json
{
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": ["src/*"],
      "@app/*": ["src/app/*"],
      "@core/*": ["src/core/*"],
      "@domain/*": ["src/domain/*"],
      "@data/*": ["src/data/*"],
      "@presentation/*": ["src/presentation/*"],
      "@shared/*": ["src/shared/*"]
    }
  }
}
```

## 7. Config env

Tạo file `.env.development`:

```env
VITE_APP_NAME=Vue Clean App
VITE_API_BASE_URL=http://localhost:3000/api
```

Tạo file `.env.production`:

```env
VITE_APP_NAME=Vue Clean App
VITE_API_BASE_URL=https://api.example.com/api
```

File `src/app/config/env.ts`:

```ts
export const env = {
  appName: import.meta.env.VITE_APP_NAME,
  apiBaseUrl: import.meta.env.VITE_API_BASE_URL,
}
```

## 8. API Client

File `src/data/datasources/remote/api-client.ts`:

```ts
import axios from 'axios'
import { env } from '@app/config/env'

export const apiClient = axios.create({
  baseURL: env.apiBaseUrl,
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
  },
})

apiClient.interceptors.request.use((config) => {
  const token = localStorage.getItem('access_token')

  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }

  return config
})

apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    return Promise.reject(error)
  },
)
```

## 9. Entity trong domain

File `src/domain/entities/user.entity.ts`:

```ts
export interface UserEntity {
  id: string
  name: string
  email: string
  avatarUrl?: string
}
```

## 10. Repository interface trong domain

File `src/domain/repositories/user.repository.ts`:

```ts
import type { UserEntity } from '@domain/entities/user.entity'

export interface UserRepository {
  login(email: string, password: string): Promise<UserEntity>
  logout(): Promise<void>
  getProfile(): Promise<UserEntity>
}
```

## 11. Use case

File `src/domain/usecases/auth/login.usecase.ts`:

```ts
import type { UserRepository } from '@domain/repositories/user.repository'
import type { UserEntity } from '@domain/entities/user.entity'

export class LoginUseCase {
  constructor(private readonly userRepository: UserRepository) {}

  async execute(email: string, password: string): Promise<UserEntity> {
    if (!email) {
      throw new Error('Email không được để trống')
    }

    if (!password) {
      throw new Error('Mật khẩu không được để trống')
    }

    return this.userRepository.login(email, password)
  }
}
```

## 12. API datasource

File `src/data/datasources/remote/auth.api.ts`:

```ts
import { apiClient } from './api-client'
import type { LoginResponseModel } from '@data/models/login-response.model'

export const authApi = {
  async login(email: string, password: string): Promise<LoginResponseModel> {
    const response = await apiClient.post<LoginResponseModel>('/auth/login', {
      email,
      password,
    })

    return response.data
  },

  async getProfile() {
    const response = await apiClient.get('/auth/profile')
    return response.data
  },

  async logout(): Promise<void> {
    await apiClient.post('/auth/logout')
  },
}
```

## 13. Model API

File `src/data/models/user.model.ts`:

```ts
export interface UserModel {
  id: string
  name: string
  email: string
  avatar_url?: string
}
```

File `src/data/models/login-response.model.ts`:

```ts
import type { UserModel } from './user.model'

export interface LoginResponseModel {
  access_token: string
  refresh_token: string
  user: UserModel
}
```

## 14. Mapper

File `src/data/mappers/user.mapper.ts`:

```ts
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
```

## 15. Repository implementation

File `src/data/repositories/user.repository.impl.ts`:

```ts
import type { UserRepository } from '@domain/repositories/user.repository'
import type { UserEntity } from '@domain/entities/user.entity'
import { authApi } from '@data/datasources/remote/auth.api'
import { mapUserModelToEntity } from '@data/mappers/user.mapper'

export class UserRepositoryImpl implements UserRepository {
  async login(email: string, password: string): Promise<UserEntity> {
    const response = await authApi.login(email, password)

    localStorage.setItem('access_token', response.access_token)
    localStorage.setItem('refresh_token', response.refresh_token)

    return mapUserModelToEntity(response.user)
  }

  async logout(): Promise<void> {
    await authApi.logout()
    localStorage.removeItem('access_token')
    localStorage.removeItem('refresh_token')
  }

  async getProfile(): Promise<UserEntity> {
    const response = await authApi.getProfile()
    return mapUserModelToEntity(response)
  }
}
```

## 16. Dependency Injection đơn giản

File `src/app/providers/dependency.provider.ts`:

```ts
import { UserRepositoryImpl } from '@data/repositories/user.repository.impl'
import { LoginUseCase } from '@domain/usecases/auth/login.usecase'

const userRepository = new UserRepositoryImpl()

export const dependencies = {
  userRepository,
  loginUseCase: new LoginUseCase(userRepository),
}
```

## 17. Pinia store

File `src/presentation/stores/auth.store.ts`:

```ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { UserEntity } from '@domain/entities/user.entity'
import { dependencies } from '@app/providers/dependency.provider'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<UserEntity | null>(null)
  const loading = ref(false)
  const errorMessage = ref<string | null>(null)

  async function login(email: string, password: string) {
    loading.value = true
    errorMessage.value = null

    try {
      user.value = await dependencies.loginUseCase.execute(email, password)
    } catch (error) {
      errorMessage.value = error instanceof Error ? error.message : 'Đăng nhập thất bại'
      throw error
    } finally {
      loading.value = false
    }
  }

  return {
    user,
    loading,
    errorMessage,
    login,
  }
})
```

## 18. Login page

File `src/presentation/pages/auth/LoginPage.vue`:

```vue
<template>
  <div class="login-page">
    <h1>Đăng nhập</h1>

    <form @submit.prevent="onSubmit">
      <input v-model="email" type="email" placeholder="Email" />
      <input v-model="password" type="password" placeholder="Mật khẩu" />

      <p v-if="authStore.errorMessage" class="error">
        {{ authStore.errorMessage }}
      </p>

      <button type="submit" :disabled="authStore.loading">
        {{ authStore.loading ? 'Đang đăng nhập...' : 'Đăng nhập' }}
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@presentation/stores/auth.store'

const router = useRouter()
const authStore = useAuthStore()

const email = ref('')
const password = ref('')

async function onSubmit() {
  await authStore.login(email.value, password.value)
  router.push('/')
}
</script>

<style scoped>
.login-page {
  max-width: 400px;
  margin: 80px auto;
}

.error {
  color: red;
}
</style>
```

## 19. Router

File `src/app/router/index.ts`:

```ts
import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '@presentation/pages/auth/LoginPage.vue'
import HomePage from '@presentation/pages/home/HomePage.vue'

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginPage,
    },
    {
      path: '/',
      name: 'home',
      component: HomePage,
      meta: {
        requiresAuth: true,
      },
    },
  ],
})

router.beforeEach((to) => {
  const token = localStorage.getItem('access_token')

  if (to.meta.requiresAuth && !token) {
    return '/login'
  }

  return true
})
```

## 20. Main app

File `src/app/main.ts`:

```ts
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import { router } from './router'
import '@shared/styles/main.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
```

File `src/app/App.vue`:

```vue
<template>
  <RouterView />
</template>
```

## 21. Quy tắc import giữa các layer

Nên tuân thủ:

```txt
presentation → domain
presentation → app
presentation → core

data → domain
data → core

app → presentation
app → data
app → domain

core → không phụ thuộc layer khác
domain → không phụ thuộc data/presentation/app
```

Không nên:

```txt
domain import axios
domain import vue
domain import pinia
usecase gọi API trực tiếp
component gọi axios trực tiếp
store xử lý quá nhiều business logic
```

## 22. Convention đặt tên file

```txt
Entity: user.entity.ts
Model API: user.model.ts
Mapper: user.mapper.ts
Repository interface: user.repository.ts
Repository implementation: user.repository.impl.ts
UseCase: login.usecase.ts
Store: auth.store.ts
Composable: useAuth.ts
Component: BaseButton.vue
Page: LoginPage.vue
Layout: MainLayout.vue
```

## 23. Khi thêm một chức năng mới

Ví dụ thêm chức năng `Product`:

```txt
1. Tạo entity: domain/entities/product.entity.ts
2. Tạo repository interface: domain/repositories/product.repository.ts
3. Tạo usecase: domain/usecases/product/get-products.usecase.ts
4. Tạo API datasource: data/datasources/remote/product.api.ts
5. Tạo model: data/models/product.model.ts
6. Tạo mapper: data/mappers/product.mapper.ts
7. Tạo repository implementation: data/repositories/product.repository.impl.ts
8. Inject vào dependency.provider.ts
9. Tạo store: presentation/stores/product.store.ts
10. Tạo page/component UI
```

## 24. Ví dụ flow đăng nhập

```txt
LoginPage.vue
    ↓ gọi
AuthStore.login()
    ↓ gọi
LoginUseCase.execute()
    ↓ gọi
UserRepository.login()
    ↓ implement bởi
UserRepositoryImpl.login()
    ↓ gọi
AuthApi.login()
    ↓ gọi
Axios POST /auth/login
```

## 25. Checklist setup base project

```txt
[ ] Tạo project bằng Vite Vue TypeScript
[ ] Cài vue-router
[ ] Cài pinia
[ ] Cài axios
[ ] Cài zod nếu cần validate
[ ] Tạo cấu trúc app/core/domain/data/presentation/shared
[ ] Cấu hình alias path
[ ] Tạo env development/production
[ ] Tạo api-client axios
[ ] Tạo entity đầu tiên
[ ] Tạo repository interface
[ ] Tạo usecase
[ ] Tạo api datasource
[ ] Tạo model
[ ] Tạo mapper
[ ] Tạo repository implementation
[ ] Tạo dependency provider
[ ] Tạo Pinia store
[ ] Tạo router guard
[ ] Tạo layout/page/component
[ ] Kiểm tra rule import giữa các layer
```

## 26. Gợi ý package cho dự án admin thực tế

```bash
npm install @vueuse/core
npm install dayjs
npm install lodash-es
npm install nprogress
```

Nếu dùng UI library:

```bash
npm install element-plus
```

Hoặc:

```bash
npm install ant-design-vue
```

## 27. Gợi ý scripts trong package.json

```json
{
  "scripts": {
    "dev": "vite --mode development",
    "build": "vue-tsc -b && vite build --mode production",
    "preview": "vite preview",
    "type-check": "vue-tsc -b",
    "lint": "eslint . --fix",
    "format": "prettier --write src/"
  }
}
```

## 28. Kết luận

Base VueJS theo Clean Architecture nên giữ nguyên tắc:

```txt
Component không gọi API trực tiếp
Store không chứa quá nhiều business logic
UseCase xử lý nghiệp vụ
Repository che giấu nguồn dữ liệu
Domain không phụ thuộc framework
Data layer chịu trách nhiệm API/local storage
Presentation chỉ lo UI và trạng thái hiển thị
```

Cấu trúc này ban đầu có thể nhiều file hơn, nhưng khi dự án lớn lên sẽ dễ bảo trì, dễ test, dễ thay đổi API và dễ chia việc cho team.

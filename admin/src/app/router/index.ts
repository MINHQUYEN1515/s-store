import { createRouter, createWebHistory } from 'vue-router'
import { tokenStorage } from '@data/datasources/local/token.storage'
import LoginPage from '@presentation/pages/auth/LoginPage.vue'
import HomePage from '@presentation/pages/home/HomePage.vue'

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginPage,
      meta: { requiresAuth: false },
    },
    {
      path: '/',
      name: 'home',
      component: HomePage,
      meta: { requiresAuth: true },
    },
  ],
})

router.beforeEach((to) => {
  const token = tokenStorage.getAccessToken()

  if (to.meta.requiresAuth && !token) {
    return '/login'
  }

  if (to.path === '/login' && token) {
    return '/'
  }

  return true
})

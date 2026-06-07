<template>
  <MainLayout>
    <div class="home-page">
      <h2 class="home-page__title">Trang chủ</h2>
      <p class="home-page__welcome">
        Xin chào{{ user ? `, ${user.name}` : '' }}! Đây là admin dashboard của S-Store.
      </p>

      <div class="home-page__cards">
        <div class="home-page__card">
          <h3>Clean Architecture</h3>
          <p>Dự án được tổ chức theo các layer: app, core, domain, data, presentation.</p>
        </div>
        <div class="home-page__card">
          <h3>Auth Flow</h3>
          <p>LoginPage → AuthStore → UseCase → Repository → API</p>
        </div>
      </div>
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import MainLayout from '@presentation/layouts/MainLayout.vue'
import { useAuth } from '@presentation/composables/useAuth'

const { user } = useAuth()
</script>

<style scoped>
.home-page__title {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 0.5rem;
}

.home-page__welcome {
  color: var(--color-text-muted);
  margin-bottom: 2rem;
}

.home-page__cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 1rem;
}

.home-page__card {
  padding: 1.25rem;
  background-color: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
}

.home-page__card h3 {
  font-size: 1rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.home-page__card p {
  font-size: 0.875rem;
  color: var(--color-text-muted);
}
</style>

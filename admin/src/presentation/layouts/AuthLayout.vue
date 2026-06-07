<template>
  <div class="auth-layout">
    <div class="auth-layout__card">
      <slot />
    </div>
  </div>
</template>

<style scoped>
.auth-layout {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 1.5rem;
  background: linear-gradient(135deg, #eef2ff 0%, #f9fafb 100%);
}

.auth-layout__card {
  width: 100%;
  max-width: 420px;
  padding: 2rem;
  background-color: var(--color-surface);
  border-radius: 12px;
  box-shadow: var(--shadow-sm), 0 4px 24px rgba(0, 0, 0, 0.06);
}
</style>

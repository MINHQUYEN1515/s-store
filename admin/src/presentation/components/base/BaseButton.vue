<template>
  <button
    :type="type"
    class="base-button"
    :class="{ 'base-button--loading': loading, 'base-button--block': block }"
    :disabled="disabled || loading"
  >
    <slot />
  </button>
</template>

<script setup lang="ts">
withDefaults(
  defineProps<{
    type?: 'button' | 'submit' | 'reset'
    disabled?: boolean
    loading?: boolean
    block?: boolean
  }>(),
  {
    type: 'button',
    disabled: false,
    loading: false,
    block: false,
  },
)
</script>

<style scoped>
.base-button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.625rem 1.25rem;
  font-size: 0.875rem;
  font-weight: 500;
  color: #fff;
  background-color: var(--color-primary);
  border: none;
  border-radius: var(--radius-md);
  transition: background-color 0.2s;
}

.base-button:hover:not(:disabled) {
  background-color: var(--color-primary-hover);
}

.base-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.base-button--block {
  width: 100%;
}

.base-button--loading {
  opacity: 0.8;
}
</style>

<template>
  <div class="base-input">
    <label v-if="label" :for="inputId" class="base-input__label">{{ label }}</label>
    <input
      :id="inputId"
      :type="type"
      :value="modelValue"
      :placeholder="placeholder"
      :disabled="disabled"
      class="base-input__field"
      @input="onInput"
    />
    <p v-if="error" class="base-input__error">{{ error }}</p>
  </div>
</template>

<script setup lang="ts">
import { useId } from 'vue'

withDefaults(
  defineProps<{
    modelValue?: string
    label?: string
    type?: string
    placeholder?: string
    disabled?: boolean
    error?: string
  }>(),
  {
    modelValue: '',
    type: 'text',
    disabled: false,
  },
)

const emit = defineEmits<{
  'update:modelValue': [value: string]
}>()

const inputId = useId()

function onInput(event: Event) {
  const target = event.target as HTMLInputElement
  emit('update:modelValue', target.value)
}
</script>

<style scoped>
.base-input {
  display: flex;
  flex-direction: column;
  gap: 0.375rem;
}

.base-input__label {
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--color-text);
}

.base-input__field {
  width: 100%;
  padding: 0.625rem 0.75rem;
  font-size: 0.875rem;
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  background-color: var(--color-surface);
  transition: border-color 0.2s;
}

.base-input__field:focus {
  outline: none;
  border-color: var(--color-primary);
}

.base-input__field:disabled {
  background-color: var(--color-bg);
  cursor: not-allowed;
}

.base-input__error {
  font-size: 0.75rem;
  color: var(--color-danger);
}
</style>

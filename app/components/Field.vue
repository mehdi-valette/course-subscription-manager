<template>
  <div class="flex align-center relative">
    <input
      ref="inputRef"
      class="w-full border border-dashed border-black p-1"
      :value="props.value"
      @change="handleChange"
      @blur="() => $emit('blur')"
      :disabled="props.status === 'processing'"
    />
    <div class="absolute right-0 self-center">
      <img
        v-if="props.status === 'success'"
        src="~/assets/icons/success.svg"
        class="w-5 h-5"
      />
      <img
        v-if="props.status === 'processing'"
        src="~/assets/icons/processing.svg"
        class="w-5 h-5 spinner"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { InputHTMLAttributes } from "nuxt/dist/app/compat/capi";
import { SynchronizationStatus } from "~/types/synchronisation";

const props = defineProps<{
  value: string;
  status: SynchronizationStatus;
}>();

const emit = defineEmits<{
  change: [value: string];
  blur: [];
}>();

const inputRef = ref<HTMLInputElement | null>(null);
const internalValue = ref("");

defineExpose({ focus });

function focus() {
  if (inputRef.value === null) {
    return;
  }

  inputRef.value.focus();
}

function handleChange(evt: Event) {
  if (evt.target === null) {
    return;
  }

  emit("change", (evt.target as InputHTMLAttributes).value);
}
</script>

<style scoped>
.spinner {
  animation-name: spin;
  animation-duration: 1s;
  animation-iteration-count: infinite;
  animation-timing-function: linear;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>

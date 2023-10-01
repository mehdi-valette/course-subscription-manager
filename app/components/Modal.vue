<template>
  <Teleport to="body">
    <dialog
      v-if="open"
      :class="clsx('absolute w-screen h-screen modal', open && 'modal-open')"
      @keydown="handleKeydown"
      @click="handleClick"
    >
      <div class="modal-box" ref="modalBody" tabindex="0">
        <slot :close="() => (open = false)" />
      </div>
    </dialog>
  </Teleport>
</template>

<script lang="ts" setup>
import clsx from "clsx";

const props = defineProps<{
  open: boolean;
}>();

const open = ref(false);
const modalBody = ref<HTMLDivElement | null>(null);

watch(
  () => props.open,
  () => {
    open.value = true;
  }
);

watch(modalBody, () => {
  if (modalBody.value !== null) {
    modalBody.value.focus();
  }
});

function handleClick(evt: MouseEvent) {
  if ((evt.target as HTMLDialogElement).tagName === "DIALOG") {
    open.value = false;
  }
}

function handleKeydown(evt: KeyboardEvent) {
  if (evt.code === "Escape") {
    open.value = false;
  }
}
</script>

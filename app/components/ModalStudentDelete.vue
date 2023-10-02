<template>
  <Modal :open="props.open">
    <template v-slot="modal">
      Are you sure you want to delete {{ fullName }} ?
      <div class="modal-action">
        <button class="btn btn-sm" @click="modal.close">Cancel</button>
        <button
          class="btn btn-sm btn-warning"
          @click="handleConfirm(modal.close)"
        >
          Delete {{ shortName }}
        </button>
      </div>
    </template>
  </Modal>
</template>

<script lang="ts" setup>
const props = defineProps<{
  open: boolean;
  callback: () => Promise<void>;
  fullName: string;
  shortName: string;
}>();

async function handleConfirm(close: () => void): Promise<void> {
  await props.callback();
  close();
}
</script>

<template>
  <Field
    v-if="editMode"
    :value="mutableData"
    :status="dataStatus"
    @change="handleChange"
    @blur="handleBlur"
    ref="inputElement"
  />
  <button
    v-else
    class="w-full h-full text-left p-1 bg-yellow-500"
    @click="setEditMode"
  >
    {{ data }}
  </button>
</template>

<script lang="ts" setup>
import { SynchronizationStatus } from "~/types/synchronisation";

const props = defineProps<{
  data: string;
}>();

const emit = defineEmits<{
  change: [value: string];
}>();

const inputElement = ref<{ focus: () => void } | null>(null);
const editMode = ref(false);
const mutableData = ref("");
const dataStatus = ref<SynchronizationStatus>("none");

watch([inputElement], () => {
  if (inputElement.value === null) {
    return;
  }

  inputElement.value.focus();
});

watch([props], () => {
  mutableData.value = props.data;
  dataStatus.value = "none";
  editMode.value = false;
});

function handleBlur() {
  if (dataStatus.value !== "processing") {
    editMode.value = false;
  }
}

function handleChange(value: string) {
  if (value === props.data) {
    editMode.value = false;
    return;
  }

  mutableData.value = value;

  dataStatus.value = "processing";
  emit("change", value);
}

function setEditMode() {
  editMode.value = true;
  mutableData.value = props.data;
}
</script>

<template>
  <Field
    v-if="editMode"
    :value="mutableData"
    :status="dataStatus"
    @change="handleChange"
    @blur="handleBlur"
    ref="inputElement"
  />
  <div v-else class="relative">
    <button
      class="w-full h-full text-left p-1 bg-yellow-500"
      @click="setEditMode"
    >
      {{ data !== "" ? data : "&nbsp;" }}
    </button>
    <button
      class="absolute right-0 h-full hover:opacity-100 opacity-50"
      @click="handleCopy"
    >
      <label :class="clsx('swap', dataCopied && 'swap-active')">
        <img
          class="swap-on w-5 h-5"
          src="~/assets/icons/copy-success.svg"
          alt="copy"
        />
        <img
          class="swap-off w-5 h-5"
          src="~/assets/icons/copy.svg"
          alt="copy"
        />
      </label>
    </button>
  </div>
</template>

<script lang="ts" setup>
import { SynchronizationStatus } from "~/types/synchronisation";
import clsx from "clsx";

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
const dataCopied = ref(false);
const dataCopiedTimeout = ref<NodeJS.Timeout | null>(null);

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

function handleCopy() {
  if (dataCopiedTimeout.value !== null) {
    clearTimeout(dataCopiedTimeout.value);
  }

  dataCopied.value = true;
  navigator.clipboard.writeText(props.data);

  dataCopiedTimeout.value = setTimeout(() => (dataCopied.value = false), 2_000);
}

function setEditMode() {
  editMode.value = true;
  mutableData.value = props.data;
}
</script>

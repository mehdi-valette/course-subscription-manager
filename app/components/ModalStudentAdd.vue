<template>
  <Modal :open="props.open">
    <template v-slot="modal">
      <form
        @submit.prevent="() => handleAddStudent(modal.close)"
        class="form-control w-full max-w-xs"
      >
        <label class="label" :for="ids.firstname">
          <span class="label-text">First name</span>
        </label>
        <input
          v-model="newStudent.firstname"
          :id="ids.firstname"
          type="text"
          class="input input-bordered w-full max-w-xs"
        />
        <label class="label" :for="ids.lastname">
          <span class="label-text">Last name</span>
        </label>
        <input
          v-model="newStudent.lastname"
          :id="ids.lastname"
          type="text"
          class="input input-bordered w-full max-w-xs"
        />
        <label class="label" :for="ids.email">
          <span class="label-text">Email</span>
        </label>
        <input
          v-model="newStudent.email"
          :id="ids.email"
          type="text"
          class="input input-bordered w-full max-w-xs"
        />
        <label class="label" :for="ids.phone">
          <span class="label-text">Phone</span>
        </label>
        <input
          v-model="newStudent.phone"
          :id="ids.phone"
          type="text"
          class="input input-bordered w-full max-w-xs"
        />

        <div class="modal-action">
          <button class="btn btn-sm btn-primary" type="submit">save</button>
          <button class="btn btn-sm" @click="modal.close">cancel</button>
        </div>
      </form>
    </template>
  </Modal>
</template>

<script lang="ts" setup>
import { nanoid } from "nanoid/non-secure";

const studentStore = useStudentStore();

const props = defineProps<{
  open: boolean;
}>();

const ids = {
  firstname: nanoid(),
  lastname: nanoid(),
  email: nanoid(),
  phone: nanoid(),
};

const newStudent = reactive({
  firstname: "",
  lastname: "",
  email: "",
  phone: "",
});

async function handleAddStudent(close: () => void) {
  await studentStore.addStudent(structuredClone(toRaw(newStudent)));
  newStudent.firstname = "";
  newStudent.lastname = "";
  newStudent.email = "";
  newStudent.phone = "";
  close();
}
</script>

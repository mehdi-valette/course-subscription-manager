<template>
  <div>
    <button class="btn" @click="handleAddUserModal">click here</button>
    <table class="table">
      <thead>
        <tr>
          <th class="w-10">Delete</th>
          <th class="w-48">First name</th>
          <th class="w-48">Last name</th>
          <th class="w-48">E-mail</th>
          <th class="w-48">Phone number</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="student in studentList ?? []" :key="student.id">
          <td>
            <button
              class="btn btn-sm btn-warning"
              @click="() => deleteStudent(student.id)"
            >
              <img src="~/assets/icons/delete.svg" alt="delete" />
            </button>
          </td>
          <td>
            <DataCell
              :data="student.firstname"
              @change="
                (value) => updateStudent({ id: student.id, firstname: value })
              "
            />
          </td>
          <td>
            <DataCell
              :data="student.lastname"
              @change="
                (value) => updateStudent({ id: student.id, lastname: value })
              "
            />
          </td>
          <td>
            <DataCell
              :data="student.email"
              @change="
                (value) => updateStudent({ id: student.id, email: value })
              "
            />
          </td>
          <td>
            <DataCell
              :data="student.phone"
              @change="
                (value) => updateStudent({ id: student.id, phone: value })
              "
            />
          </td>
        </tr>
      </tbody>
    </table>
    <ModalAddStudent :open="openAddUserModal" />
  </div>
</template>

<script lang="ts" setup>
const { data: studentList, refresh } = await useFetch("/api/student");
const openAddUserModal = ref(false);

async function updateStudent(body: {
  id: number;
  [key: string]: string | number;
}) {
  await $fetch(`/api/student/${body.id}`, { method: "patch", body });
  await refresh();
}

async function deleteStudent(id: number) {
  await $fetch(`api/student/${id}`, { method: "DELETE" });
  await refresh();
}

function handleAddUserModal() {
  openAddUserModal.value = !openAddUserModal.value;
}
</script>

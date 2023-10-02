<template>
  <div>
    <button class="btn" @click="() => (openAddUserModal = !openAddUserModal)">
      click here
    </button>
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
              @click="() => deleteStudentConfirm(student)"
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
    <ModalStudentAdd :open="openAddUserModal" />
    <ModalStudentDelete
      :full-name="studentToBeDeleted.fullName"
      :short-name="studentToBeDeleted.shortName"
      :open="studentToBeDeleted.open"
      :callback="studentToBeDeleted.callback"
    />
  </div>
</template>

<script lang="ts" setup>
const { data: studentList, refresh } = await useFetch("/api/student");

type Student = Exclude<typeof studentList.value, null>[0];

const openAddUserModal = ref(false);
const studentToBeDeleted = reactive({
  shortName: "",
  fullName: "",
  callback: async () => {},
  open: false,
});

async function updateStudent(body: {
  id: number;
  [key: string]: string | number;
}) {
  await $fetch(`/api/student/${body.id}`, { method: "patch", body });
  await refresh();
}

async function deleteStudentConfirm(student: Student): Promise<void> {
  studentToBeDeleted.fullName = `${student.firstname} ${student.lastname}`;
  studentToBeDeleted.shortName = student.firstname;
  studentToBeDeleted.callback = async () => {
    await $fetch(`api/student/${student.id}`, { method: "DELETE" });
    await refresh();
  };
  studentToBeDeleted.open = !studentToBeDeleted.open;
}
</script>

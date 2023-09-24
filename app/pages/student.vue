<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th class="w-48">First name</th>
          <th class="w-48">Last name</th>
          <th class="w-48">E-mail</th>
          <th class="w-48">Phone number</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="student in studentList" :key="student.id">
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
  </div>
</template>

<script lang="ts" setup>
const { data: studentList, refresh } = await useFetch("/api/student");

async function updateStudent(body: {
  id: number;
  [key: string]: string | number;
}) {
  await $fetch(`/api/student/${body.id}`, { method: "PATCH", body });
  await refresh();
}
</script>

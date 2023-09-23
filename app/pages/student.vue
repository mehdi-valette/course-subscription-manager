<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th>First name</th>
          <th>Last name</th>
          <th>E-mail</th>
          <th>Phone number</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="student in studentList" :key="student.id">
          <td>{{ student.firstname }}</td>
          <td>{{ student.lastname }}</td>
          <td>{{ student.email }}</td>
          <td>{{ student.phone }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts" setup>
const { data: studentList } = await useFetch("/api/students");
</script>

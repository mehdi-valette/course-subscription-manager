<template>
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="sessionType in sessionTypeList" :key="sessionType.id">
        <td>{{ sessionType.name }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts" setup>
const { data: sessionTypeList } = useFetch("/api/session-type");
</script>

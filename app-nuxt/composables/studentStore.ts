import { defineStore } from "pinia";
import { student } from "@/server/database/schema";

export type Student = typeof student.$inferSelect;

export const useStudentStore = defineStore("student", () => {
  const _studentList = ref<Student[] | null>(null);

  const studentList = computed(() => _studentList);

  async function initialFetch() {
    if (_studentList.value !== null) {
      return;
    }

    await refreshStudentList();
  }

  async function refreshStudentList() {
    const { data, refresh, error } = await useFetch("/api/student");

    if (data.value !== null) {
      _studentList.value = data.value;
    }
  }

  async function addStudent(student: Omit<Student, "id">) {
    await useFetch(`/api/student`, { method: "post", body: student });
    await refreshStudentList();
  }

  async function deleteStudent(studentId: Pick<Student, "id">) {
    await $fetch(`/api/student/${studentId}`, { method: "DELETE" });
    await refreshStudentList();
  }

  async function updateStudent(
    student: Partial<Student> & Pick<Student, "id">
  ) {
    await $fetch(`/api/student/${student.id}`, {
      method: "patch",
      body: student,
    });

    await refreshStudentList();
  }

  return {
    studentList,
    initialFetch,
    addStudent,
    deleteStudent,
    updateStudent,
  };
});

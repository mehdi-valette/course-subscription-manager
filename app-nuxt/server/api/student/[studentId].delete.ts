import { eq } from "drizzle-orm";
import { db } from "../../database/db";
import { student } from "../../database/schema";

export default defineEventHandler(async (event) => {
  const studentId = Number(getRouterParam(event, "studentId"));

  if (isNaN(studentId)) {
    throw createError({ statusCode: 400 });
  }

  return await db.delete(student).where(eq(student.id, studentId));
});

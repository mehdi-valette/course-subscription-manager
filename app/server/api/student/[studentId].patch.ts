import { db } from "../../database/db";
import { student } from "../../database/schema";
import { createInsertSchema } from "drizzle-zod";
import { eq } from "drizzle-orm";

const bodySchema = createInsertSchema(student);

export default defineEventHandler(async (event) => {
  const rawBody = await readBody(event);
  const studentId = Number(getRouterParam(event, "studentId"));
  const body = bodySchema.safeParse(rawBody);

  if (isNaN(studentId) || !body.success) {
    throw createError({ statusCode: 400 });
  }

  try {
    await db.update(student).set(body.data).where(eq(student.id, studentId));
  } catch (e) {
    throw createError({
      statusCode: 500,
      statusMessage: "Unknown error",
    });
  }

  return {
    hello: "done",
  };
});

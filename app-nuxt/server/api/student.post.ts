import { db } from "../database/db";
import { student } from "../database/schema";
import { createInsertSchema } from "drizzle-zod";

const bodySchema = createInsertSchema(student);

export default defineEventHandler(async (event) => {
  const rawBody = await readBody(event);

  const body = bodySchema.safeParse(rawBody);

  if (!body.success) {
    throw createError({ statusCode: 400 });
  }

  try {
    await db.insert(student).values({
      email: body.data.email,
      firstname: body.data.firstname,
      lastname: body.data.lastname,
      phone: body.data.phone,
    });
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

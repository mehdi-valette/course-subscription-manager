import { db } from "../database/db";

export default defineEventHandler(async (event) => {
  return await db.query.student.findMany();
});

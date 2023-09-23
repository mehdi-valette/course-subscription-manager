import { integer, text, sqliteTable } from "drizzle-orm/sqlite-core";

export const student = sqliteTable("student", {
  id: integer("id", { mode: "number" }).primaryKey(),
  firstname: text("firstname"),
  lastname: text("lastname"),
  phone: text("phone"),
  email: text("email"),
});

export const sessionType = sqliteTable("sessionType", {
  id: integer("id", { mode: "number" }).primaryKey(),
  name: text("name"),
});

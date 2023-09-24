import { integer, text, sqliteTable } from "drizzle-orm/sqlite-core";

export const student = sqliteTable("student", {
  id: integer("id", { mode: "number" }).primaryKey(),
  firstname: text("firstname").default("").notNull(),
  lastname: text("lastname").default("").notNull(),
  phone: text("phone").default("").notNull(),
  email: text("email").default("").notNull(),
});

export const sessionType = sqliteTable("sessionType", {
  id: integer("id", { mode: "number" }).primaryKey(),
  name: text("name"),
});

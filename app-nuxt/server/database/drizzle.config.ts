import type { Config } from "drizzle-kit";

export default {
  schema: "./schema.ts",
  driver: "better-sqlite",
  out: "./drizzle",
  dbCredentials: {
    url: process.env.DB_URL ?? "./database.sqlite",
  },
} satisfies Config;

import { drizzle } from "drizzle-orm/better-sqlite3";
import { migrate } from "drizzle-orm/better-sqlite3/migrator";
import Database from "better-sqlite3";
import * as schema from "./schema";

const betterSqlite = new Database("./server/database/database.sqlite");
export const db = drizzle(betterSqlite, { schema });

migrate(db, { migrationsFolder: "./server/database/drizzle" });

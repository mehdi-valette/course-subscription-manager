CREATE TABLE `sessionType` (
	`id` integer PRIMARY KEY NOT NULL,
	`name` text
);
--> statement-breakpoint
CREATE TABLE `student` (
	`id` integer PRIMARY KEY NOT NULL,
	`firstname` text,
	`lastname` text,
	`phone` text,
	`email` text
);

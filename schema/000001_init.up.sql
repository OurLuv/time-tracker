CREATE SEQUENCE users_id_seq START 1;

CREATE TABLE "users" (
  "id" integer DEFAULT nextval('users_id_seq') PRIMARY KEY,
  "passport_number" varchar,
  "name" varchar,
  "surname" varchar,
  "patronymic" varchar,
  "address" varchar
);

CREATE SEQUENCE tasks_id_seq START 1;

CREATE TABLE "tasks" (
  "id" integer DEFAULT nextval('tasks_id_seq') PRIMARY KEY,
  "title" varchar,
  "started_at" timestamp,
  "finished_at" timestamp,
  "is_finished" boolean,
  "user_id" integer
);
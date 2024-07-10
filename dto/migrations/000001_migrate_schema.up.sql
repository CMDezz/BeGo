CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar,
  "role" varchar,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "title" varchar,
  "body" text,
  "user_id" bigserial,
  "status" varchar,
  "created_at" timestamp DEFAULT (now())
);

COMMENT ON COLUMN "posts"."body" IS 'Content of the post';

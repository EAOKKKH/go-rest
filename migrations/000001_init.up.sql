CREATE TABLE "urls" (
  "id" serial PRIMARY KEY,
  "user_id" integer NOT NULL,
  "alias" varchar UNIQUE NOT NULL,
  "url" varchar NOT NULL
);

CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "login" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE INDEX ON "urls" ("alias");

CREATE INDEX ON "users" ("login");

ALTER TABLE "urls" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
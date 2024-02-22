CREATE TABLE "urls" (
  "id" serial PRIMARY KEY,
  "alias" varchar UNIQUE NOT NULL,
  "url" varchar NOT NULL
);

CREATE INDEX ON "urls" ("alias");
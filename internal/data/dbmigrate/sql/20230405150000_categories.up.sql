CREATE TABLE IF NOT EXISTS "categories" (
    "id" UUID PRIMARY KEY,
    "created_at" TIMESTAMP NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "short_desc" VARCHAR(255) NOT NULL
);

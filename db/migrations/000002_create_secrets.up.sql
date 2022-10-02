BEGIN;

CREATE TABLE "secrets" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "name" varchar NOT NULL,
    "login" varchar,
    "password" varchar,
    "email" varchar,
    "website" varchar,
    "note" text,
    "pkey" text,
    "deleted_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (NOW()),
    "updated_at" timestamptz NOT NULL DEFAULT (NOW())
);

CREATE INDEX ON "secrets" ("user_id");
CREATE INDEX ON "secrets" ("name");

ALTER TABLE "secrets" ADD CONSTRAINT "fk_secrets_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id");

COMMIT;
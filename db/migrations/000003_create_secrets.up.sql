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
    "deleted_at" timestamp without time zone,
    "created_at" timestamp without time zone NOT NULL DEFAULT (NOW()),
    "updated_at" timestamp without time zone NOT NULL DEFAULT (NOW())
);

CREATE INDEX ON "secrets" ("user_id");
CREATE INDEX ON "secrets" ("name");

ALTER TABLE "secrets" ADD CONSTRAINT "fk_secrets_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TRIGGER set_timestamp_secrets BEFORE UPDATE ON secrets FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

COMMIT;
BEGIN;

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "jti" varchar NOT NULL,
    "encrypted" varchar NOT NULL,
    "active" boolean NOT NULL DEFAULT(FALSE),
    "deleted_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (NOW()),
    "updated_at" timestamptz NOT NULL DEFAULT (NOW())
);

CREATE INDEX ON "users" ("email");
CREATE INDEX ON "users" ("jti");

COMMIT;
BEGIN;

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "jti" varchar NOT NULL,
    "encrypted" varchar NOT NULL,
    "active" boolean NOT NULL DEFAULT(FALSE),
    "deleted_at" timestamp without time zone,
    "created_at" timestamp without time zone NOT NULL DEFAULT (NOW()),
    "updated_at" timestamp without time zone NOT NULL DEFAULT (NOW())
);

CREATE INDEX ON "users" ("email");
CREATE INDEX ON "users" ("jti");

CREATE TRIGGER set_timestamp_users BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

COMMIT;
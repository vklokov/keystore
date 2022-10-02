BEGIN;

DROP TRIGGER trigger_set_timestamp ON secrets;

ALTER TABLE secrets DROP CONSTRAINT fk_secrets_users;

DROP TABLE IF EXISTS secrets;

COMMIT;
BEGIN;

DROP TRIGGER trigger_set_timestamp ON users;

DROP TABLE IF EXISTS users;

COMMIT;
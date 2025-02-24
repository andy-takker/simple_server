BEGIN;

DROP INDEX IF EXISTS idx_users_username_not_deleted;

DROP INDEX IF EXISTS idx_users_email_not_deleted;

DROP TABLE IF EXISTS users;

COMMIT;
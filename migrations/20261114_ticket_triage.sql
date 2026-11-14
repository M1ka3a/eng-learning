-- Review against the actual jobs table before applying. This project does not
-- yet have an automated migration runner. This script is not run by the app.
BEGIN;

ALTER TABLE jobs
    ALTER COLUMN ctx TYPE jsonb USING ctx::jsonb,
    ADD COLUMN IF NOT EXISTS result jsonb,
    ADD COLUMN IF NOT EXISTS error_message text;

-- The original local schema is not stored in this repository. The conventional
-- jobs_status_check is replaced when present; inspect any differently named
-- status constraint before applying this migration.
ALTER TABLE jobs DROP CONSTRAINT IF EXISTS jobs_status_check;
ALTER TABLE jobs ADD CONSTRAINT jobs_status_check
    CHECK (status IN ('needs_input', 'pending', 'running', 'succeeded', 'failed'));

COMMIT;

ALTER TABLE users
    ALTER COLUMN updated_at SET NOT NULL,
    ALTER COLUMN updated_at SET DEFAULT CURRENT_TIMESTAMP;

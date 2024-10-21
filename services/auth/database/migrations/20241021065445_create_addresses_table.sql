-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS addresses (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    street TEXT,
    zip VARCHAR(10) NULL NULL,
    area VARCHAR(100),
    district VARCHAR(100),
    division VARCHAR(30),
    is_default BOOLEAN,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS addresses_deleted_at ON addresses(deleted_at);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP INDEX addresses_deleted_at;
DROP TABLE IF EXISTS addresses;

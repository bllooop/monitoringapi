-- +goose Up
-- +goose StatementBegin
CREATE TABLE pingdata (
    id SERIAL PRIMARY KEY,
    ip VARCHAR(45) NOT NULL,
    last_success TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pingdata;
-- +goose StatementEnd

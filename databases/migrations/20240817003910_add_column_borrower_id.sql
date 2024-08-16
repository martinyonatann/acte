-- +goose Up
-- +goose StatementBegin
ALTER TABLE loans ADD borrower_id int NULL AFTER id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE loans DROP COLUMN borrower_id;
-- +goose StatementEnd

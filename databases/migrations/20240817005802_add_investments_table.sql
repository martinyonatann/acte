-- +goose Up
-- +goose StatementBegin
CREATE TABLE `investments` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    loan_state_id INT NOT NULL,
    investor_id INT NOT NULL,
    investor_agreement_letter_path TEXT NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `investments`;
-- +goose StatementEnd

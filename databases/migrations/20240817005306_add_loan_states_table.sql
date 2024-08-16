-- +goose Up
-- +goose StatementBegin
CREATE TABLE `loan_states` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    borrower_id VARCHAR(255) NOT NULL,
    principal_amount DECIMAL(15, 2) NOT NULL,
    rate DECIMAL(5, 2) NOT NULL,
    roi DECIMAL(5, 2) NOT NULL,
    state ENUM('proposed', 'approved', 'invested', 'disbursed') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `loan_states`;
-- +goose StatementEnd

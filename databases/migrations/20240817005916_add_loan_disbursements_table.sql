-- +goose Up
-- +goose StatementBegin
CREATE TABLE loan_disbursements (
    id INT AUTO_INCREMENT PRIMARY KEY,
    loan_state_id INT NOT NULL,
    borrower_agreement_letter_path TEXT NOT NULL,
    employee_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `loan_disbursements`;
-- +goose StatementEnd

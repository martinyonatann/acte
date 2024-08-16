-- +goose Up
-- +goose StatementBegin
CREATE TABLE `loan_approvals` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    loan_state_id INT NOT NULL,
    employee_id INT NOT NULL,
    proof_path TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `loan_approvals`;
-- +goose StatementEnd

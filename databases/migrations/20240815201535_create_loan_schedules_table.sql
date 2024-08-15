-- +goose Up
-- +goose StatementBegin
CREATE TABLE `loan_schedules` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `loan_id` INT NOT NULL,
  `amount` DECIMAL(12, 2) DEFAULT NULL,
  `schedule_date` VARCHAR(10) DEFAULT NULL,
  `is_paid` TINYINT(1) DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `loan_schedules`;
-- +goose StatementEnd

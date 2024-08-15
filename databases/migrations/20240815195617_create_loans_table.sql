-- +goose Up
-- +goose StatementBegin
CREATE TABLE `loans` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `amount` DECIMAL(12, 2) DEFAULT NULL,
  `outstanding` DECIMAL(12, 2) DEFAULT NULL,
  `interest_rate` DECIMAL(5, 2) DEFAULT NULL,
  `term_of_payment_in_week` INT DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `loans`;
-- +goose StatementEnd

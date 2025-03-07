-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user_cards` (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    stripe_customer_id VARCHAR(255) NOT NULL,
    stripe_card_id VARCHAR(255) NOT NULL,
    last4 VARCHAR(4) NOT NULL,
    brand VARCHAR(50) NOT NULL,
    exp_month INT NOT NULL,
    exp_year INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    INDEX `idx_user_cards_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_cards`;
-- +goose StatementEnd

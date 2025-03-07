-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `payments` (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    booking_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    payment_method ENUM('credit_card', 'paypal', 'vnpay', 'momo', 'stripe') DEFAULT 'credit_card',
    status ENUM('pending', 'paid', 'failed', 'refunded') DEFAULT 'pending',
    stripe_payment_intent_id VARCHAR(255) UNIQUE,
    stripe_charge_id VARCHAR(255) UNIQUE,
    stripe_customer_id VARCHAR(255),
    stripe_refund_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (booking_id) REFERENCES bookings(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    INDEX `idx_payments_booking` (`booking_id`),
    INDEX `idx_payments_status` (`status`),
    INDEX `idx_payments_method` (`payment_method`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `payments`;
-- +goose StatementEnd

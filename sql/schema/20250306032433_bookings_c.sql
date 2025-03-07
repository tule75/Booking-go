-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `bookings` (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    property_id CHAR(36) NOT NULL,
    room_id CHAR(36) NULL,
    check_in DATE NOT NULL,
    check_out DATE NOT NULL,
    guests INT NOT NULL,
    status ENUM('pending', 'confirmed', 'cancelled', 'completed') DEFAULT 'pending',
    total_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (property_id) REFERENCES properties(id) ON DELETE CASCADE,
    FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE SET NULL,

    INDEX `idx_bookings_user` (`user_id`),
    INDEX `idx_bookings_property` (`property_id`),
    INDEX `idx_bookings_status` (`status`),
    INDEX `idx_bookings_check_in` (`check_in`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `bookings`;
-- +goose StatementEnd

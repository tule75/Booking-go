-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `rooms` (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    property_id CHAR(36) NOT NULL,
    name VARCHAR(255),
    price DECIMAL(10,2) NOT NULL,
    max_guests INT NOT NULL,
    is_available BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (property_id) REFERENCES properties(id) ON DELETE CASCADE,

    INDEX `idx_rooms_property` (`property_id`),
    INDEX `idx_rooms_price` (`price`),
    INDEX `idx_rooms_available` (`is_available`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `rooms`;
-- +goose StatementEnd

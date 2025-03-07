-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `properties` (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    owner_id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    location VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    amenities JSON,
    rating DECIMAL(3,2) DEFAULT 0.0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE,

    INDEX `idx_properties_location` (`location`),
    INDEX `idx_properties_price` (`price`),
    INDEX `idx_properties_rating` (`rating`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `properties`
-- +goose StatementEnd

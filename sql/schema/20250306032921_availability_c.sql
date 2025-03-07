-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not EXISTS `availability` (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    room_id CHAR(36) NOT NULL,
    date DATE NOT NULL,
    is_available BOOLEAN DEFAULT TRUE,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE,

    INDEX `idx_availability_room` (`room_id`),
    INDEX `idx_availability_date` (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS `availability`;
-- +goose StatementEnd

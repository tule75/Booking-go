-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not EXISTS `notifications` (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    message TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    INDEX `idx_notifications_user` (`user_id`),
    INDEX `idx_notifications_status` (`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop table if exists `notifications`;
-- +goose StatementEnd

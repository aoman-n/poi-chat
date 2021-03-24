
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE messages (
  `id` int UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `user_uid` varchar(255) NOT NULL,
  `room_id` int(11) UNSIGNED NOT NULL,
  `body` text NOT NULL,
  `user_name` varchar(255) NOT NULL,
  `user_avatar_url` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT messages_fk_room_id
    FOREIGN KEY (`room_id`)
    REFERENCES rooms (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS messages;

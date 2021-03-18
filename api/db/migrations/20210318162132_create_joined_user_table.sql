
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE joined_users (
  `id` int UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `room_id` int(11) UNSIGNED NOT NULL,
  `avatar_url` varchar(255) NOT NULL,
  `display_name` varchar(255) NOT NULL,
  `user_id` varchar(255) NOT NULL,
  `x` int(11) NOT NULL,
  `y` int(11) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT joined_users_fk_room_id
    FOREIGN KEY (`room_id`)
    REFERENCES rooms (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS joined_users;

-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE rooms (
  `id` int(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `owner_id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `background_url` varchar(255),
  `background_color` varchar(255),
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT rooms_fk_owner_id
    FOREIGN KEY (`owner_id`)
    REFERENCES users (`id`)
    ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS rooms;

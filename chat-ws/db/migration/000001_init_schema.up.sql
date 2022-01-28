CREATE TABLE `user_detail` (
  `user_id` VARCHAR(255) PRIMARY KEY,
  `username` VARCHAR(255) UNIQUE NOT NULL,
  `email` VARCHAR(255) UNIQUE NOT NULL,
  `phone_no` VARCHAR(255) UNIQUE NOT NULL,
  `created_date` DATETIME NOT NULL DEFAULT (now()),
  `updated_date` DATETIME NOT NULL
);

CREATE TABLE `user_credential` (
  `credential_id` CHAR(36) PRIMARY KEY,
  `user_id` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `role` VARCHAR(255) NOT NULL,
  `status` VARCHAR(255) NOT NULL,
  `created_date` DATETIME NOT NULL DEFAULT (now()),
  `updated_date` DATETIME NOT NULL
);

CREATE TABLE `user_access_info` (
  `credential_id` CHAR(36) PRIMARY KEY,
  `last_seen` DATETIME NOT NULL
);

CREATE TABLE `verify_email_token` (
  `token_id` CHAR(36) PRIMARY KEY,
  `user_id` VARCHAR(255) NOT NULL,
  `created_date` DATETIME NOT NULL DEFAULT (now())
);

CREATE TABLE `reset_password_token` (
  `token_id` CHAR(36) PRIMARY KEY,
  `user_id` VARCHAR(255) NOT NULL,
  `created_date` DATETIME NOT NULL DEFAULT (now())
);

CREATE TABLE `workspace` (
  `workspace_id` CHAR(36) PRIMARY KEY,
  `workspace_name` VARCHAR(255) NOT NULL,
  `creator_id` VARCHAR(255) NOT NULL,
  `created_date` DATETIME NOT NULL DEFAULT (now())
);

CREATE TABLE `workspace_user` (
  `workspace_id` CHAR(36),
  `user_id` VARCHAR(255),
  PRIMARY KEY (`workspace_id`, `user_id`)
);

CREATE TABLE `channel` (
  `channel_id` CHAR(36) PRIMARY KEY,
  `channel_name` VARCHAR(255) NOT NULL,
  `workspace_id` CHAR(36) NOT NULL,
  `creator_id` VARCHAR(255) NOT NULL,
  `is_default` BOOLEAN NOT NULL,
  `created_date` DATETIME NOT NULL DEFAULT (now())
);

CREATE TABLE `channel_user` (
  `channel_id` CHAR(36),
  `user_id` VARCHAR(255),
  `last_accessed` DATETIME NOT NULL DEFAULT (now()),
  PRIMARY KEY (`channel_id`, `user_id`)
);

CREATE TABLE `message` (
  `message_id` CHAR(36) PRIMARY KEY,
  `type` VARCHAR(36) NOT NULL,
  `channel_id` CHAR(36) NOT NULL,
  `creator_id` CHAR(36) NOT NULL,
  `content` LONGTEXT NOT NULL,
  `file_name` VARCHAR(255),
  `file_type` VARCHAR(255),
  `created_date` DATETIME NOT NULL DEFAULT (now())
);

CREATE TABLE `task_board` (
  `board_id` CHAR(36) PRIMARY KEY,
  `board_name` VARCHAR(255) NOT NULL,
  `workspace_id` CHAR(36) NOT NULL
);

CREATE TABLE `task_list` (
  `list_id` CHAR(36) PRIMARY KEY,
  `list_name` VARCHAR(255) NOT NULL,
  `board_id` CHAR(36) NOT NULL,
  `order_index` INT NOT NULL
);

CREATE TABLE `card` (
  `card_id` CHAR(36) PRIMARY KEY,
  `title` VARCHAR(255) NOT NULL,
  `priority` VARCHAR(6) NOT NULL,
  `list_id` CHAR(36) NOT NULL,
  `description` LONGTEXT NOT NULL,
  `creator_id` VARCHAR(255) NOT NULL,
  `created_date` DATETIME NOT NULL DEFAULT (now()),
  `due_date` DATETIME NOT NULL,
  `order_index` INT NOT NULL
);

CREATE TABLE `card_user` (
  `card_id` CHAR(36),
  `user_id` VARCHAR(255),
  PRIMARY KEY (`card_id`, `user_id`)
);

CREATE TABLE `checklist_item` (
  `item_id` CHAR(36) PRIMARY KEY,
  `content` VARCHAR(255) NOT NULL,
  `is_checked` BOOLEAN NOT NULL,
  `card_id` CHAR(36) NOT NULL
);

CREATE TABLE `card_comment` (
  `comment_id` CHAR(36) PRIMARY KEY,
  `type` VARCHAR(36) NOT NULL,
  `card_id` CHAR(36) NOT NULL,
  `creator_id` VARCHAR(255) NOT NULL,
  `content` LONGTEXT NOT NULL,
  `file_name` VARCHAR(255),
  `file_type` VARCHAR(255),
  `created_date` DATETIME NOT NULL DEFAULT (now())
);

CREATE TABLE `ticket` (
  `ticket_id` CHAR(36) PRIMARY KEY,
  `title` VARCHAR(255) NOT NULL,
  `type` VARCHAR(36) NOT NULL,
  `status` VARCHAR(255) NOT NULL,
  `priority` VARCHAR(5) NOT NULL,
  `description` LONGTEXT NOT NULL,
  `card_id` CHAR(36) NOT NULL,
  `creator_id` VARCHAR(255) NOT NULL,
  `created_date` DATETIME NOT NULL DEFAULT (now()),
  `due_date` DATETIME NOT NULL
);

ALTER TABLE `user_credential` ADD FOREIGN KEY (`user_id`) REFERENCES `user_detail` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `user_access_info` ADD FOREIGN KEY (`credential_id`) REFERENCES `user_credential` (`credential_id`) ON DELETE CASCADE;

ALTER TABLE `workspace` ADD FOREIGN KEY (`creator_id`) REFERENCES `user_detail` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `verify_email_token` ADD FOREIGN KEY (`user_id`) REFERENCES `user_credential` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `reset_password_token` ADD FOREIGN KEY (`user_id`) REFERENCES `user_credential` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `workspace_user` ADD FOREIGN KEY (`workspace_id`) REFERENCES `workspace` (`workspace_id`) ON DELETE CASCADE;

ALTER TABLE `workspace_user` ADD FOREIGN KEY (`user_id`) REFERENCES `user_detail` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `channel` ADD FOREIGN KEY (`workspace_id`) REFERENCES `workspace` (`workspace_id`) ON DELETE CASCADE;

ALTER TABLE `channel` ADD FOREIGN KEY (`creator_id`) REFERENCES `user_detail` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `channel_user` ADD FOREIGN KEY (`channel_id`) REFERENCES `channel` (`channel_id`) ON DELETE CASCADE;

ALTER TABLE `channel_user` ADD FOREIGN KEY (`user_id`) REFERENCES `user_detail` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `message` ADD FOREIGN KEY (`creator_id`) REFERENCES `user_detail` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `message` ADD FOREIGN KEY (`channel_id`) REFERENCES `channel` (`channel_id`) ON DELETE CASCADE;

ALTER TABLE `task_board` ADD FOREIGN KEY (`workspace_id`) REFERENCES `workspace` (`workspace_id`) ON DELETE CASCADE;

ALTER TABLE `task_list` ADD FOREIGN KEY (`board_id`) REFERENCES `task_board` (`board_id`) ON DELETE CASCADE;

ALTER TABLE `card` ADD FOREIGN KEY (`list_id`) REFERENCES `task_list` (`list_id`) ON DELETE CASCADE;

ALTER TABLE `card` ADD FOREIGN KEY (`creator_id`) REFERENCES `user_detail` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `card_user` ADD FOREIGN KEY (`card_id`) REFERENCES `card` (`card_id`) ON DELETE CASCADE;

ALTER TABLE `card_user` ADD FOREIGN KEY (`user_id`) REFERENCES `user_detail` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `checklist_item` ADD FOREIGN KEY (`card_id`) REFERENCES `card` (`card_id`) ON DELETE CASCADE;

ALTER TABLE `card_comment` ADD FOREIGN KEY (`card_id`) REFERENCES `card` (`card_id`) ON DELETE CASCADE;

ALTER TABLE `card_comment` ADD FOREIGN KEY (`creator_id`) REFERENCES `user_detail` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `ticket` ADD FOREIGN KEY (`card_id`) REFERENCES `card` (`card_id`) ON DELETE CASCADE;

ALTER TABLE `ticket` ADD FOREIGN KEY (`creator_id`) REFERENCES `user_detail` (`user_id`) ON DELETE CASCADE;

CREATE TRIGGER trigger_user_detail
BEFORE INSERT
ON user_detail
FOR EACH ROW
BEGIN
  IF (NEW.user_id IS NULL) THEN
    SELECT
      MAX(user_id) INTO @max_id
    FROM
      user_detail;

    IF (@max_id IS NULL) THEN
      SET NEW.user_id = 'U000001';
    ELSE
      SET NEW.user_id = CONCAT(SUBSTR(@max_id, 1, 1), LPAD(SUBSTR(@max_id, 2) + 1, 6, '0'));
    END IF;
  END IF;
END;
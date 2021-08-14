CREATE TABLE `user` (
  `user_id` VARCHAR(255) PRIMARY KEY,
  `username` VARCHAR(255) UNIQUE NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `role` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) UNIQUE NOT NULL,
  `phone_no` VARCHAR(255) UNIQUE NOT NULL,
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
  `created_date` DATETIME NOT NULL DEFAULT (now())
);

CREATE TABLE `channel_user` (
  `channel_id` CHAR(36),
  `user_id` VARCHAR(255),
  PRIMARY KEY (`channel_id`, `user_id`)
);

CREATE TRIGGER trigger_user
BEFORE INSERT
ON user
FOR EACH ROW
BEGIN
  IF (NEW.user_id IS NULL) THEN
    SELECT
      MAX(user_id) INTO @max_id
    FROM
      user
    WHERE
      role= NEW.role;

    IF (@max_id IS NULL) THEN
      SET @id =
        CASE NEW.role
        WHEN 'ADMIN' THEN 'A'
        WHEN 'USER' THEN 'U'
        ELSE 'UNKNOWN'
      END;

      SET NEW.user_id = CONCAT(@id, '000001');
    ELSE
      SET NEW.user_id = CONCAT(SUBSTR(@max_id, 1, 1), LPAD(SUBSTR(@max_id, 2) + 1, 6, '0'));
    END IF;
  END IF;
END;

ALTER TABLE `workspace` ADD CONSTRAINT `fk_workspace_user_creator_id` FOREIGN KEY (`creator_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `workspace_user` ADD CONSTRAINT `fk_workspace_user_workspace_workspace_id` FOREIGN KEY (`workspace_id`) REFERENCES `workspace` (`workspace_id`) ON DELETE CASCADE;

ALTER TABLE `workspace_user` ADD CONSTRAINT `fk_workspace_user_workspace_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `channel` ADD CONSTRAINT `fk_channel_workspace_workspace_id` FOREIGN KEY (`workspace_id`) REFERENCES `workspace` (`workspace_id`) ON DELETE CASCADE;

ALTER TABLE `channel` ADD CONSTRAINT `fk_channel_user_user_id` FOREIGN KEY (`creator_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `channel_user` ADD CONSTRAINT `fk_channel_user_channel_channel_id` FOREIGN KEY (`channel_id`) REFERENCES `channel` (`channel_id`) ON DELETE CASCADE;

ALTER TABLE `channel_user` ADD CONSTRAINT `fk_channel_user_user_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE;

CREATE TABLE workspace 
(
	workspace_id VARCHAR(255) PRIMARY KEY,
	workspace_name VARCHAR(255) NOT NULL
);

CREATE TABLE account_credential
(
	user_id VARCHAR(255) PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	role VARCHAR(255) NOT NULL,
	created_date DATETIME NOT NULL
);

CREATE TRIGGER trigger_account_credential
BEFORE INSERT
ON account_credential
FOR EACH ROW
BEGIN
  IF (NEW.user_id IS NULL) THEN
    SELECT
      MAX(user_id) INTO @max_id
    FROM
      account_credential 
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
END

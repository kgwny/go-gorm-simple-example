DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
  user_id VARCHAR(36) PRIMARY KEY,
  user_name VARCHAR(256) NOT NULL,
  age TINYINT UNSIGNED NOT NULL,
  gender VARCHAR(256) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO users (user_id, user_name, age, gender) VALUES (UUID(), 'suzuki ichiro', 30, 'male');
INSERT INTO users (user_id, user_name, age, gender) VALUES (UUID(), 'yamada hanako', 20, 'female');
INSERT INTO users (user_id, user_name, age, gender) VALUES (UUID(), 'yamamoto taro', 10, 'male');

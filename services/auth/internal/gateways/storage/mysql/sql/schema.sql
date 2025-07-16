
CREATE TABLE user (
  id INT AUTO_INCREMENT PRIMARY KEY,
	username  TEXT,
	first_name TEXT,
	last_name  TEXT,
	email     TEXT NOT NULL,
	password   BLOB NOT NULL,
	avatar_url  TEXT,
	last_seen_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	is_active   bool NOT NULL DEFAULT true,
	birth_date  DATE,
	provider_id int,
	oauth_acc_id TEXT UNIQUE
);

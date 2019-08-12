-- +goose Up
CREATE TABLE users (
  id int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  email VARCHAR(255),
  display_name VARCHAR(255),
  photo_url VARCHAR(255),
  ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  utime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE reports (
  id int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id int(11) UNSIGNED NOT NULL,
  done TEXT NOT NULL,
  todo TEXT,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='list of reports';

CREATE TABLE comments (
  id int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  report_id int(11) UNSIGNED NOT NULL,
  user_id int(11) UNSIGNED NOT NULL,
  body TEXT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  deleted_at timestamp DEFAULT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (report_id) REFERENCES reports(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='list of comments';

CREATE TABLE tags (
  id int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  report_id int(11) UNSIGNED NOT NULL,
  tag_name varchar(255) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (report_id) REFERENCES reports(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='list of tags';

-- +goose Down
DROP TABLE users;
DROP TABLE reports;
DROP TABLE comments;
DROP TABLE tags;

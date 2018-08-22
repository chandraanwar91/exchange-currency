
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE currency_exchange (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  `from` char(3) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `to` char(3) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `status` enum('enable','disable','removed') COLLATE utf8_unicode_ci DEFAULT 'enable',
  created_by mediumint(6) unsigned NOT NULL DEFAULT 0,
  modified_by mediumint(6) unsigned NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT '1970-01-01 00:00:01',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE currency_exchange;

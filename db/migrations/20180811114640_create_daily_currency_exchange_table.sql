
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE daily_currency_exchange (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  currency_date date NOT NULL,
  currency_exchange_id char(3) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  exchange_rate decimal(17,7) DEFAULT '0.00',
  created_by mediumint(6) unsigned NOT NULL DEFAULT 0,
  modified_by mediumint(6) unsigned NOT NULL DEFAULT 0,
  created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE daily_currency_exchange;

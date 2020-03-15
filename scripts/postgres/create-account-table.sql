CREATE TABLE account(
   account_id serial PRIMARY KEY,
   username VARCHAR (256) UNIQUE NOT NULL,
   latest_cursor VARCHAR(256) NOT NULL,
   shortcodes VARCHAR(32)[] NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
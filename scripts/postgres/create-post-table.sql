CREATE TYPE media_type AS ENUM ('image', 'video');

CREATE TABLE post(
    post_id serial PRIMARY KEY,
    account_id INT NOT NULL,
    shortcode VARCHAR(32) NOT NULL,
    content TEXT NOT NULL,
    like_count INT NOT NULL,
    media_type media_type NOT NULL,
    taken_at INT NOT NULL,
    media_url VARCHAR(1024) NOT NULL,
    download_path VARCHAR(1024) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
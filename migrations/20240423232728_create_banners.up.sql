CREATE TABLE IF NOT EXISTS users
(
    banner_id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    hash_password VARCHAR(255) NOT NULL,
    role VARCHAR(10) NOT NULL,

    CONSTRAINT chk_role CHECK (role IN ('user', 'admin'))
);

CREATE TABLE IF NOT EXISTS banners (
                         banner_id SERIAL PRIMARY KEY,
                         feature_id INTEGER UNIQUE,
                         tag_id INTEGER UNIQUE,
                         content JSONB NOT NULL,
                         status BOOLEAN NOT NULL DEFAULT FALSE,
                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY (banner_id),
                         FOREIGN KEY (banner_id) REFERENCES banners(banner_id),
                         FOREIGN KEY (tag_id) REFERENCES tags(tag_id)
);

CREATE TABLE IF NOT EXISTS tags (
                      tag_id SERIAL PRIMARY KEY,
                      name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS features (
                          feature_id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS banner_tags (
                             banner_id INT NOT NULL,
                             tag_id INT NOT NULL
);
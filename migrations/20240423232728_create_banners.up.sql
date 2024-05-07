CREATE TABLE IF NOT EXISTS banners (
                         banner_id SERIAL PRIMARY KEY UNIQUE ,
                         feature_id INTEGER UNIQUE,
                         tag_ids INTEGER[] UNIQUE,
                         content JSONB NOT NULL,
                         is_active BOOLEAN NOT NULL DEFAULT FALSE,
                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (banner_id) REFERENCES banner_relations(banner_id),
--                          FOREIGN KEY (tag_ids) REFERENCES banner_relations(tag_id)
                         FOREIGN KEY (feature_id) REFERENCES banner_relations(feature_id)
);


CREATE TABLE IF NOT EXISTS banner_relations (
                             banner_id INT NOT NULL UNIQUE ,
                             tag_id INT NOT NULL,
                             feature_id INT NOT NULL UNIQUE
);

DROP TABLE banner_relations;

DROP TABLE banners;
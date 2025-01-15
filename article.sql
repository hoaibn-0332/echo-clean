-- DO $$
--     BEGIN
--        IF EXISTS (SELECT 1 FROM pg_database WHERE datname = 'article') THEN
--           PERFORM pg_terminate_backend(pid)
--           FROM pg_stat_activity
--           WHERE datname = 'article';
--
--         EXECUTE 'DROP DATABASE article';
--     END IF;
-- END $$;

\c article;

CREATE TABLE IF NOT EXISTS author (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS article (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id INT REFERENCES author(id) ON DELETE SET NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO author (id, name, created_at, updated_at)
VALUES (1, 'John Doe', '2017-05-18 13:50:19','2017-05-18 13:50:19');

INSERT INTO author (id, name, created_at, updated_at)
VALUES (2, 'Bach', '2018-05-18 13:50:19','2018-05-18 13:50:19');

INSERT INTO article (title, content, author_id, created_at, updated_at)
VALUES ('First Article', 'This is the content of the first article.', 1, '2017-05-18 13:50:19','2017-05-18 13:50:19')
    ON CONFLICT (id) DO NOTHING;

INSERT INTO article (title, content, author_id, created_at, updated_at)
VALUES ('Second Article', 'This is the content of the second article.', 2, '2018-05-18 13:50:19','2018-05-18 13:50:19')
ON CONFLICT (id) DO NOTHING;
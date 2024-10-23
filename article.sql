DO $$
    BEGIN
       IF EXISTS (SELECT 1 FROM pg_database WHERE datname = 'article') THEN
          PERFORM pg_terminate_backend(pid)
          FROM pg_stat_activity
          WHERE datname = 'article';

        EXECUTE 'DROP DATABASE article';
    END IF;
END $$;

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

INSERT INTO author (name)
VALUES ('John Doe');

INSERT INTO article (title, content, author_id)
VALUES ('First Article', 'This is the content of the first article.', 1)
    ON CONFLICT (id) DO NOTHING;
-- Log start of initialization
\echo 'Starting database initialization...'

-- Create tables
\echo 'Creating tables...'
CREATE TABLE IF NOT EXISTS technologies (
    name    VARCHAR(255) PRIMARY KEY,
    details VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS blogs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    coverURL VARCHAR(255),
    body TEXT NOT NULL
);

-- Grant privileges
\echo 'Granting privileges...'
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO postgres;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO postgres;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO postgres;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO postgres;

-- Insert data
\echo 'Inserting initial data...'
INSERT INTO blogs (title, coverURL, body) VALUES 
    ('First Blog', 'https://example.com/cover1.jpg', 'This is the first blog post'),
    ('Second Blog', 'https://example.com/cover2.jpg', 'This is the second blog post')
ON CONFLICT DO NOTHING;

INSERT INTO technologies VALUES 
    ('Go', 'An open source programming language that makes it easy to build simple and efficient software.'),
    ('JavaScript', 'A lightweight, interpreted, or just-in-time compiled programming language with first-class functions.'),
    ('PostgreSQL', 'A powerful, open source object-relational database system')
ON CONFLICT DO NOTHING;

\echo 'Database initialization complete!'

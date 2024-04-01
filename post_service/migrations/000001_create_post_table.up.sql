CREATE TABLE posts
(
    title TEXT,
    content TEXT,
    image_url TEXT,
    id UUID,
    owner_id UUID,
    likes INT,
    views INT,
    category TEXT,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP
)
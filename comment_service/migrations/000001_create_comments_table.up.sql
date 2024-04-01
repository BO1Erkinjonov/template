CREATE TABLE comments
(
    id UUID,
    description TEXT,
    post_id UUID,
    owner_id UUID,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP
)
CREATE TABLE comments
(
    id UUID,
    description TEXT,
    post_id UUID,
    owner_id UUID,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO comments (id, description, post_id, owner_id)
VALUES
    ('fbfe3fa8-2613-4f65-9bc6-30c35eee0422', 'This is a comment on the post.', 'ed04c889-dead-4b37-afb2-ecd7eeb12e87', '154d572c-ab7f-46c3-89b7-b941eeaf9da4'),
    ('0654ae8c-9dd9-4047-ab63-c15aa446de3c', 'Another comment here.', 'ed04c889-dead-4b37-afb2-ecd7eeb12e87', '154d572c-ab7f-46c3-89b7-b941eeaf9da4'),
    ('01dd7e42-e349-4a7a-871b-308985e21db2', 'Yet another comment.', '4d2931cb-975a-4d88-9f13-46362cf5b27d', '154d572c-ab7f-46c3-89b7-b941eeaf9da4'),
    ('3bce3069-bf6b-4525-976e-2845ee93f2ff', 'This is a comment by the specified owner on the specified post.', '4d2931cb-975a-4d88-9f13-46362cf5b27d', '154d572c-ab7f-46c3-89b7-b941eeaf9da4'),
    ('5d60e6db-a2c1-4270-a71d-c1d7d2ed50d2', 'Another comment here.', '4d2931cb-975a-4d88-9f13-46362cf5b27d', '154d572c-ab7f-46c3-89b7-b941eeaf9da4'),
    ('38ecfd98-c7cd-4997-86dc-e14f009753d4', 'Yet another comment.', 'fb0a0b74-8b5c-4dfb-8f74-47fcae053cc5', 'e404c46f-3762-47a0-9a81-af86ff81b403'),
    ('818b092c-cb59-49ca-be05-378e88cfab32', 'One more comment.', '3e351629-1df1-4f4e-b705-6da6a388b615', 'e404c46f-3762-47a0-9a81-af86ff81b403');




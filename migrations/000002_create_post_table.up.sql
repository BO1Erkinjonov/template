CREATE TABLE posts
(
    title TEXT,
    content TEXT,
    image_url TEXT,
    id UUID,
    owner_id UUID,
    likes INT DEFAULT 0,
    views INT DEFAULT 0,
    category TEXT,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO posts (title, content, image_url, id, owner_id, category)
VALUES
    ('Title 4', 'Content 4', 'https://example.com/image4.jpg', 'ed04c889-dead-4b37-afb2-ecd7eeb12e87', '154d572c-ab7f-46c3-89b7-b941eeaf9da4', 'Technology'),
    ('Title 5', 'Content 5', 'https://example.com/image5.jpg', '4d2931cb-975a-4d88-9f13-46362cf5b27d', '154d572c-ab7f-46c3-89b7-b941eeaf9da4', 'Science'),
    ('Title 6', 'Content 6', 'https://example.com/image6.jpg', '80780e61-c620-47c5-bf00-c56b904a7b50', '154d572c-ab7f-46c3-89b7-b941eeaf9da4', 'Arts'),
    ('Title 7', 'Content 7', 'https://example.com/image7.jpg', '5d5ee6e0-156f-4bb8-9219-3ed996de89c5', '1ec8df45-f671-4e97-89b9-9c4b269d940c', 'Technology'),
    ('Title 8', 'Content 8', 'https://example.com/image8.jpg', '91062c8b-3c40-4110-8503-c099bdafa3b8', '1ec8df45-f671-4e97-89b9-9c4b269d940c', 'Science'),
    ('Title 9', 'Content 9', 'https://example.com/image9.jpg', '240908a4-fa6a-4550-adf3-0fb438a8ca66', '1ec8df45-f671-4e97-89b9-9c4b269d940c', 'Arts'),
    ('Title 10', 'Content 10', 'https://example.com/image10.jpg', '63479da2-a246-4593-9b05-419e414b9269', 'e404c46f-3762-47a0-9a81-af86ff81b403', 'Technology'),
    ('Title 11', 'Content 11', 'https://example.com/image11.jpg', '91c7efab-f3ec-42d1-98b4-f4fb46314f10', 'e404c46f-3762-47a0-9a81-af86ff81b403', 'Science'),
    ('Title 12', 'Content 12', 'https://example.com/image12.jpg', 'fb0a0b74-8b5c-4dfb-8f74-47fcae053cc5', 'e404c46f-3762-47a0-9a81-af86ff81b403', 'Arts'),
    ('Title 13', 'Content 13', 'https://example.com/image13.jpg', '3e351629-1df1-4f4e-b705-6da6a388b615', 'e404c46f-3762-47a0-9a81-af86ff81b403', 'Technology'),
    ('Title 14', 'Content 14', 'https://example.com/image14.jpg', '75f67a39-65f2-4e62-b535-68adb715a4d3', 'e404c46f-3762-47a0-9a81-af86ff81b403', 'Science'),
    ('Title 17', 'Content 17', 'https://example.com/image17.jpg', 'bc6a3e61-da12-4114-a781-c5578af0f30f', '667354fd-7a14-4706-8db6-9f4d03b44e5b', 'Science'),
    ('Title 18', 'Content 18', 'https://example.com/image18.jpg', '5a6dfbba-54e1-4462-9bf5-41413c7c47fa', '667354fd-7a14-4706-8db6-9f4d03b44e5b', 'Technology'),
    ('Title 19', 'Content 19', 'https://example.com/image19.jpg', '1c736e1c-3b34-466f-9c2a-a6878026dc5d', '667354fd-7a14-4706-8db6-9f4d03b44e5b', 'Arts');


CREATE TABLE users
(
    first_name   VARCHAR(30),
    last_name    VARCHAR(30),
    username     VARCHAR(30),
    role VARCHAR(30),
    password     TEXT,
    email        TEXT,
    id           UUID PRIMARY KEY,
    refreshtoken TEXT,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO users (id, first_name, last_name, username, role, password, email, refreshtoken) VALUES
    ('0fea21f6-bd09-474c-9f17-99f210ae3a9c', 'John', 'Doe', 'johndoe', 'superAdmin', 'password123', 'axrorbek@example.com', 'refresh_token_1'),
    ('154d572c-ab7f-46c3-89b7-b941eeaf9da4', 'Jane', 'Smith', 'janesmith', 'user', 'pass123', 'jane@example.com', 'refresh_token_2'),
    ('1ec8df45-f671-4e97-89b9-9c4b269d940c', 'Michael', 'Johnson', 'michaelj', 'user', 'securepass', 'michael@example.com', 'refresh_token_3'),
    ('bfdc5bb8-9065-4630-8961-ba602b3d8801', 'Emily', 'Brown', 'emilyb', 'admin', '123456', 'emily@example.com', 'refresh_token_4'),
    ('e404c46f-3762-47a0-9a81-af86ff81b403', 'David', 'Martinez', 'davidm', 'user', 'davpass', 'david@example.com', 'refresh_token_5'),
    ('853eccd7-cbbe-44fa-8d9b-f8f1c7ab968e', 'Emma', 'Wilson', 'emmaw', 'user', 'password', 'emma@example.com', 'refresh_token_6'),
    ('667354fd-7a14-4706-8db6-9f4d03b44e5b', 'William', 'Anderson', 'willa', 'admin', 'abc123', 'will@example.com', 'refresh_token_7'),
    ('bee1d6a0-dc0b-4e84-a668-52cb1787b914', 'Olivia', 'Taylor', 'oliviat', 'user', 'qwerty', 'olivia@example.com', 'refresh_token_8'),
    ('dd4b2b15-6708-4ee8-82bb-ba3936b0ecec', 'James', 'Thomas', 'jamest', 'admin', 'passpass', 'james@example.com', 'refresh_token_9'),
    ('aaee8d3a-cfa8-488b-8abe-fa00af1cb388', 'Sophia', 'White', 'sophiaw', 'user', 'letmein', 'sophia@example.com', 'refresh_token_10'),
    ('838eef35-47ea-447a-b5fe-cff3df631f54', 'Benjamin', 'Clark', 'benc', 'admin', 'adminpass', 'ben@example.com', 'refresh_token_11'),
    ('d852aa9e-50b9-4a92-a257-fb568767b55c', 'Mia', 'Hall', 'miah', 'user', 'hello123', 'mia@example.com', 'refresh_token_12'),
    ('9280d273-7a8c-4d0a-895c-459bf76a9285', 'Ethan', 'Lewis', 'ethanl', 'user', 'p@ssw0rd', 'ethan@example.com', 'refresh_token_13'),
    ('39acbe92-a2a7-46c3-8540-86d105aba26a', 'Charlotte', 'Young', 'charlottey', 'admin', 'test123', 'charlotte@example.com', 'refresh_token_14'),
    ('f9765b75-e0ef-4d41-9f23-51bd1f4769da', 'Alexander', 'Harris', 'alexanderh', 'user', 'changeme', 'alexander@example.com', 'refresh_token_15');


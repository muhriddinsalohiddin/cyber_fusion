-- farrux
CREATE TABLE "user" (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    gender BOOLEAN NOT NULL,
    birthday DATE NOT NULL,
    email TEXT NOT NULL,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    bio TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);


-- Saidakbar
CREATE TABLE "post" (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES "user" (id),
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
-- Said
CREATE TABLE "comment" (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES "user" (id),
    post_id UUID NOT NULL REFERENCES "post" (id),
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- Lazizbek
CREATE TABLE "like" (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES "user" (id),
    post_id UUID NOT NULL REFERENCES "post" (id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Asror
CREATE TABLE "notification" (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES "user" (id),
    type TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- MuhammadYusuf
CREATE TABLE "message" (
    id UUID PRIMARY KEY,
    sender_id UUID NOT NULL REFERENCES "user" (id),
    receiver_id UUID NOT NULL REFERENCES "user" (id),
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Alibek
CREATE TABLE "author" (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
); 
-- Umar
CREATE TABLE "book" (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    author UUID NOT NULL REFERENCES "author" (id),
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- Alibek
CREATE TABLE "author" (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
); 
INSERT INTO "post" (id, user_id, title, body, created_at, updated_at)
VALUES
    ('72864b4a-a657-44ef-abc3-a559b723da5f', '4e2ff5de-8bcb-4098-95f7-8ce14671b901', 'First Post', 'This is the body of the first post.', '2024-01-08 14:00:00', NULL),
    ('e98e22db-bd66-4ea8-bd93-a67392efcc03', '4e2ff5de-8bcb-4098-95f7-8ce14671b901', 'Second Post', 'This is the body of the second post.', '2024-01-08 15:30:00', NULL);
        
INSERT INTO "comment" (id, user_id, post_id, body, created_at, updated_at)
VALUES
    ('516d6102-ca06-4282-a40e-b82dc976ab55', '4e2ff5de-8bcb-4098-95f7-8ce14671b901', '72864b4a-a657-44ef-abc3-a559b723da5f', 'Great post! Thanks for sharing.', '2024-01-08 16:00:00', NULL),
    ('0b455e83-73dc-4577-ad01-0162059c4f2c', '4e2ff5de-8bcb-4098-95f7-8ce14671b901', '72864b4a-a657-44ef-abc3-a559b723da5f', 'I enjoyed reading this. Well done!', '2024-01-08 17:30:00', NULL);

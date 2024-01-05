-- farrux
CREATE TABLE "user" (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    gender BOOLEAN NOT NULL,
    birthday DATE NOT NULL,
    email TEXT NOT NULL,
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


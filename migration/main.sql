-- farrux
CREATE TABLE IF NOT EXISTS "user" (
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
CREATE TABLE IF NOT EXISTS "post" (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES "user" (id),
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
-- Said
CREATE TABLE IF NOT EXISTS "comment" (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES "user" (id),
    post_id UUID NOT NULL REFERENCES "post" (id),
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- Lazizbek
CREATE TABLE IF NOT EXISTS "like" (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES "user" (id),
    post_id UUID NOT NULL REFERENCES "post" (id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Asror
CREATE TABLE IF NOT EXISTS "notification" (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES "user" (id),
    type TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- MuhammadYusuf
CREATE TABLE IF NOT EXISTS "message" (
    id UUID PRIMARY KEY,
    sender_id UUID NOT NULL REFERENCES "user" (id),
    receiver_id UUID NOT NULL REFERENCES "user" (id),
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- Alibek
CREATE TABLE IF NOT EXISTS "author" (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
); 
-- Umar
CREATE TABLE IF NOT EXISTS "book" (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    author UUID NOT NULL REFERENCES "author" (id),
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- create extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- generate uuid
SELECT uuid_generate_v4();

-- insert data in user table
INSERT INTO "user" (
id,name,gender,birthday,email,login,password,bio) VALUES
(uuid_generate_v4(),'Farrux',true,'1999-01-01','bekorchijon@mail.ru','farrux','123','I am a programmer'),
(uuid_generate_v4(),'Umid',true,'2000-01-01','galatibek@mail.ru','umidjon','123','I am a doctor');


-- insert data in post table
INSERT INTO "post" (
id,user_id,title,body) VALUES
(uuid_generate_v4(),(SELECT id FROM "user" WHERE name='Farrux'),'First post','This is my first post'),
(uuid_generate_v4(),(SELECT id FROM "user" WHERE name='Umid'),'Second post','This is my second post');

-- insert data in comment table
INSERT INTO "comment" (
id,user_id,post_id,body) VALUES
(uuid_generate_v4(),(SELECT id FROM "user" WHERE name='Farrux'),(SELECT id FROM "post" WHERE title='First post'),'This is my first comment'),
(uuid_generate_v4(),(SELECT id FROM "user" WHERE name='Umid'),(SELECT id FROM "post" WHERE title='Second post'),'This is my second comment');

-- insert data in like table
INSERT INTO "like" (
id,user_id,post_id) VALUES
(uuid_generate_v4(),(SELECT id FROM "user" WHERE name='Farrux'),(SELECT id FROM "post" WHERE title='First post')),
(uuid_generate_v4(),(SELECT id FROM "user" WHERE name='Umid'),(SELECT id FROM "post" WHERE title='Second post'));

-- insert data in notification table
INSERT INTO "notification" (
id,user_id,type,body) VALUES
(uuid_generate_v4(),(SELECT id FROM "user" WHERE name='Farrux'),'like','Farrux liked your post'),
(uuid_generate_v4(),(SELECT id FROM "user" WHERE name='Umid'),'comment','Umid commented your post');

-- insert data in message table
INSERT INTO "message" (
id,sender_id,receiver_id,body) VALUES
(uuid_generate_v4(),(SELECT id FROM "user" WHERE name='Farrux'),(SELECT id FROM "user" WHERE name='Umid'),'Hello Umid'),
(uuid_generate_v4(),(SELECT id FROM "user" WHERE name='Umid'),(SELECT id FROM "user" WHERE name='Farrux'),'Hello Farrux');


-- insert data in author table
INSERT INTO "author" (
id,name) VALUES
(uuid_generate_v4(),'Author 1'),
(uuid_generate_v4(),'Author 2');

-- insert data in book table
INSERT INTO "book" (
id,title,author,description) VALUES
(uuid_generate_v4(),'Book 1',(SELECT id FROM "author" WHERE name='Author 1'),'Description 1'),
(uuid_generate_v4(),'Book 2',(SELECT id FROM "author" WHERE name='Author 2'),'Description 2');


SELECT "author".id, "author".name, "book".id, "book".title, "book".description  FROM "author" JOIN "book" ON "author".id = "book".author;

--  SELECT Table1.column1, Table1.column2, Table2.column3
-- FROM Table1
-- JOIN Table2 ON Table1.id = Table2.id;

SELECT "post".id, "post".title, "post".body, "comment".
body, "like".id FROM "post"
Join "comment" on "post".id = "comment".id; 

-- SELECT "post".id, "post".title, "post".body, "comment".body AS comment_body, "like".id AS like_id
-- FROM "post"
-- JOIN "comment" ON "post".id = "comment".post_id
-- LEFT JOIN "like" ON "post".id = "like".post_id;


-- SELECT "user".id AS user_id, "user".name AS user_name, "post".id  "comment".body, "like".id AS post_id, "post".title, "post".body, 
-- FROM "user"
-- JOIN "post" ON "user".id = "post".user_id;

-- SELECT "user".id AS user_id, "user".name AS user_name,
--        "post".id AS post_id, "post".title, "post".body,
--        "comment".body AS comment_body, "like".id AS like_id
-- FROM "user"
-- JOIN "post" ON "user".id = "post".user_id
-- LEFT JOIN "comment" ON "post".id = "comment".post_id
-- LEFT JOIN "like" ON "post".id = "like".post_id;


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

SELECT array_agg(p),* FROM "user" u join "post" p on u.id = p.user_id GROUP BY c.id;

SELECT u.id, u.name,u.gender,u.birthday,u.email,u.login,u.password,u.bio,u.created_at,u.updated_at, array_agg(p.id) posts
FROM "user" u
INNER JOIN "post" p
ON u.id = p.user_id
GROUP BY  u.id;


SELECT 
  u.id, 
  u.name, 
  u.gender, 
  u.birthday, 
  u.email, 
  u.login, 
  u.password, 
  u.bio, 
  to_char(u.created_at, 'YYYY-MM-DD') as created_at,
  to_char(u.updated_at, 'YYYY-MM-DD') as updated_at,
  COALESCE(
    (
      SELECT 
      json_agg(
        json_build_object(
          'id', m.id,
          'sender_id', m.sender_id,
          'receiver_id', m.receiver_id,
          'body', m.body,
          'created_at', to_char(m.created_at, 'YYYY-MM-DD'),
          'updated_at', to_char(m.updated_at, 'YYYY-MM-DD')
        )
      )
      FROM "message" m
      WHERE u.id = m.sender_id 
    ),
    '[]'
  ) as messages,

  COALESCE(
    (
      SELECT 
      json_agg(
        json_build_object(
          'id', n.id,
          'user_id', n.user_id,
          'type', n.type,
          'body', n.body,
          'created_at', to_char(n.created_at, 'YYYY-MM-DD')
        )
      )
      FROM "notification" n
      WHERE u.id = n.user_id 
    ),
    '[]'
  ) as notifications,

  COALESCE(
    (
      SELECT 
      json_agg(
        json_build_object(
          'id', p.id, 
          'user_id', p.user_id, 
          'title', p.title, 
          'body', p.body, 
          'created_at', to_char(p.created_at, 'YYYY-MM-DD'),
          'updated_at', to_char(p.updated_at, 'YYYY-MM-DD'),

        	'likes', COALESCE(
            (
              SELECT 
            	json_agg(
                json_build_object(
                  'id', l.id,
                  'user_id', l.user_id,
                  'post_id', l.post_id,
                  'created_at', to_char(l.created_at, 'YYYY-MM-DD')
                )
              )
              FROM "like" l
              WHERE p.id = l.post_id
            ),
            '[]'
          ),

          'comments', COALESCE(
            (
              SELECT 
              json_agg(
                json_build_object(
                  'id', c.id,
                  'user_id', c.user_id,
                  'post_id', c.post_id,
                  'body', c.body,
                  'created_at', to_char(c.created_at, 'YYYY-MM-DD'),
                  'updated_at', to_char(c.created_at, 'YYYY-MM-DD')
                )
              )
              FROM "comment" c
              WHERE p.id = c.post_id
            ),
            '[]'
          )
        )
      )
      FROM "post" p
      WHERE u.id = p.user_id 
    ),
    '[]'
  ) as posts 
FROM 
	"user" u
GROUP BY  
	u.id;

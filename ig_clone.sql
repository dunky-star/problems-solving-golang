-- 1. Create the database if it doesn’t exist
CREATE DATABASE IF NOT EXISTS ig_clone;

-- 2. Switch to that database
USE ig_clone;

-- 3. Create the users table with UUID primary key
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username STRING(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

-- Create the photos table (UUID-based)
CREATE TABLE IF NOT EXISTS photos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    image_url STRING(500) NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO photos (image_url, user_id)
VALUES
  ('https://example.com/photo1.jpg', (SELECT id FROM users WHERE username='dunky')),
  ('https://example.com/photo2.jpg', (SELECT id FROM users WHERE username='geoffrey')),
  ('https://example.com/photo3.jpg', (SELECT id FROM users WHERE username='dunky'));
  
  
CREATE TABLE IF NOT EXISTS comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    comment_text STRING(255) NOT NULL,
    photo_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    FOREIGN KEY (photo_id) REFERENCES photos(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO comments (comment_text, photo_id, user_id)
VALUES
  ('Nice shot!', (SELECT id FROM photos WHERE image_url LIKE '%photo1%'), (SELECT id FROM users WHERE username='dunky')),
  ('Nice, niceee!', (SELECT id FROM photos WHERE image_url LIKE '%photo2%'), (SELECT id FROM users WHERE username='geoffrey')),
  ('Awesome!', (SELECT id FROM photos WHERE image_url LIKE '%photo3%'), (SELECT id FROM users WHERE username='dunky'));


CREATE TABLE IF NOT EXISTS likes (
    user_id UUID NOT NULL,
    photo_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (photo_id) REFERENCES photos(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, photo_id)
);

CREATE TABLE IF NOT EXISTS follows (
    follower_id UUID NOT NULL,
    followee_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (followee_id) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (follower_id, followee_id)
);

INSERT INTO follows (follower_id, followee_id)
VALUES
  ((SELECT id FROM users WHERE username='dunky'), (SELECT id FROM users WHERE username='geoffrey')),
  ((SELECT id FROM users WHERE username='geoffrey'), (SELECT id FROM users WHERE username='dunky'));
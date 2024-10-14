ALTER TABLE posts
ADD CONSTRAINT fk_user FOREIGN key (user_id) REFERENCES users (id);
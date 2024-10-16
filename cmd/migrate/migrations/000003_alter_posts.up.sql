ALTER TABLE posts
ADD CONSTRAINT fk_user FOREIGN key (user_id) REFERENCES users (id);

ALTER TABLE posts
ADD COLUMN tags varchar(100) [];

ALTER TABLE posts
ADD COLUMN updated_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT now();
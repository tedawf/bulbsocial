CREATE TABLE IF NOT EXISTS followers (
    user_id bigint NOT NULL,
    follower_id bigint NOT NULL,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    PRIMARY KEY (user_id, follower_id),
    FOREIGN key (user_id) REFERENCES users (id) ON DELETE cascade,
    FOREIGN key (follower_id) REFERENCES users (id) ON DELETE cascade
);
INSERT INTO users (name, nick, email, password)
VALUES
("user1", "user1", "user1@example.com", "123456"),
("user2", "user2", "user2@example.com", "123456"),
("user3", "user3", "user3@example.com", "123456");

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(2, 1),
(1, 3);

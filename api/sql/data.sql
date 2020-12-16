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

INSERT INTO publications (title, content, author_id)
VALUES
("publication1", "publication1", 1),
("publication2", "publication2", 2),
("publication3", "publication3", 3);

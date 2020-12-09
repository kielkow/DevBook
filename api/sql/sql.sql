CREATE DATABASE IF NOT EXISTS devbook; 

USE devbook; 

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers; 

CREATE TABLE users( 
    id        INT auto_increment PRIMARY KEY, 
    name      VARCHAR(50) NOT NULL, 
    nick      VARCHAR(50) NOT NULL UNIQUE, 
    email     VARCHAR(50) NOT NULL UNIQUE, 
    password  VARCHAR(100) NOT NULL, 
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP() 
) ENGINE=INNODB;

CREATE TABLE followers( 
    user_id     INT NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES  users(id)
    ON DELETE CASCADE,

    follower_id INT NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES  users(id)
    ON DELETE CASCADE,

    PRIMARY KEY(user_id, follower_id)
) ENGINE=INNODB;

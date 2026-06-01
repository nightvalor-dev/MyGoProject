CREATE TABLE users (
                       id            SERIAL PRIMARY KEY,
                       username      VARCHAR(50)  UNIQUE NOT NULL,
                       email         VARCHAR(50)  UNIQUE NOT NULL,
                       phone         VARCHAR(20),
                       bio           VARCHAR(100),
                       password_hash VARCHAR(255) NOT NULL,
                       role          VARCHAR(20)  NOT NULL DEFAULT 'user',
                       created_at    TIMESTAMP DEFAULT now(),
                       updated_at    TIMESTAMP DEFAULT now()
);


CREATE TABLE blogs (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(50) NOT NULL,
                       content TEXT NOT NULL,
                       user_id INT NOT NULL REFERENCES users(id),
                       status VARCHAR(10) NOT NULL CHECK (status IN ('drafted', 'published', 'archived')),
                       created_at TIMESTAMP DEFAULT now(),
                       updated_at TIMESTAMP DEFAULT now(),
                       published_at TIMESTAMP DEFAULT now()
);

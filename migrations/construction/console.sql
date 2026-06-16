-- USERS
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       email VARCHAR(50) UNIQUE NOT NULL,
                       phone VARCHAR(20),
                       bio VARCHAR(100),
                       password_hash VARCHAR(255) NOT NULL,
                       role VARCHAR(20) NOT NULL DEFAULT 'user',
                       created_at TIMESTAMP DEFAULT now(),
                       updated_at TIMESTAMP DEFAULT now()
);

-- CATEGORIES
CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            category_name VARCHAR(50) UNIQUE NOT NULL,
                            description VARCHAR(100),
                            created_at TIMESTAMP DEFAULT now()
);

-- BLOGS
CREATE TABLE blogs (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(100) NOT NULL,
                       content TEXT NOT NULL,

                       user_id INT NOT NULL,
                       category_id INT REFERENCES categories(id),

                       status VARCHAR(10) NOT NULL
                           CHECK (status IN ('drafted', 'published', 'archived')),

                       created_at TIMESTAMP DEFAULT now(),
                       updated_at TIMESTAMP DEFAULT now(),
                       published_at TIMESTAMP,

                       CONSTRAINT fk_blog_user
                           FOREIGN KEY (user_id)
                               REFERENCES users(id)
                               ON DELETE CASCADE,

                       CONSTRAINT fk_blog_category
                           FOREIGN KEY (category_id)
                               REFERENCES categories(id)
                               ON DELETE SET NULL
);

-- TAGS
CREATE TABLE tags (
                      id SERIAL PRIMARY KEY,
                      tag_name VARCHAR(50) UNIQUE NOT NULL,
                      created_at TIMESTAMP DEFAULT now()
);

-- BLOG_TAGS (JUNCTION TABLE)
CREATE TABLE blog_tags (
                           blog_id INT NOT NULL,
                           tag_id INT NOT NULL,

                           PRIMARY KEY (blog_id, tag_id),

                           CONSTRAINT fk_blog_tags_blog
                               FOREIGN KEY (blog_id)
                                   REFERENCES blogs(id)
                                   ON DELETE CASCADE,

                           CONSTRAINT fk_blog_tags_tag
                               FOREIGN KEY (tag_id)
                                   REFERENCES tags(id)
                                   ON DELETE CASCADE
);

-- COMMENTS
CREATE TABLE comments (
                          id SERIAL PRIMARY KEY,
                          content TEXT NOT NULL,

                          blog_id INT NOT NULL,
                          user_id INT NOT NULL,

                          created_at TIMESTAMP DEFAULT now(),
                          updated_at TIMESTAMP DEFAULT now(),

                          CONSTRAINT fk_comment_blog
                              FOREIGN KEY (blog_id)
                                  REFERENCES blogs(id)
                                  ON DELETE CASCADE,

                          CONSTRAINT fk_comment_user
                              FOREIGN KEY (user_id)
                                  REFERENCES users(id)
                                  ON DELETE CASCADE
);
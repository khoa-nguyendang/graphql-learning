CREATE TABLE IF NOT EXISTS authors
(
    id serial PRIMARY KEY,
    name varchar(100) NOT NULL,
    email varchar(150) NOT NULL,
    created_at date
);

CREATE TABLE IF NOT EXISTS posts
(
    id serial PRIMARY KEY,
    title varchar(100) NOT NULL,
    content text NOT NULL,
    author_id int,
    created_at date
);

Insert into authors(name, email, created_at) values ('author 1', 'author1@gmail.com', '2024-06-01'), ('author 2', 'author2@gmail.com', '2024-06-01');
Insert into posts(title, content, author_id, created_at) values 
('post 1', 'author1 just post 1st', 1, '2024-06-01'), 
('post 2', 'author1  just post 2nd', 1, '2024-06-01'),
('post 3', 'author1  just post 3rd', 1, '2024-06-01'),
('post 4', 'author2 just post 4th', 2, '2024-06-01'),
('post 5', 'author2  just post 5th', 2, '2024-06-01');
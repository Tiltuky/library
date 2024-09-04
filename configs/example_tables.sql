CREATE TABLE Authors
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE Books
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author_id INTEGER NOT NULL,
    published_at DATE NOT NULL,
    isbn VARCHAR(13) NOT NULL UNIQUE,
    FOREIGN KEY (author_id) REFERENCES Authors(id)
);

CREATE TABLE Users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE RentedBooks
(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    book_id INTEGER NOT NULL,
    rented_at TIMESTAMP NOT NULL,
    returned_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (book_id) REFERENCES Books(id),
    UNIQUE (book_id) -- Гарантирует, что книга не может быть арендована более чем одним пользователем одновременно
);

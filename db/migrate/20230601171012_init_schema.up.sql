CREATE TABLE IF NOT EXISTS author (
    id serial PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS books (
    id serial PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    author_id integer references author(id) NOT NULL,
    UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS takenBook (
    user_id integer references users(id),
    book_id integer references books(id)
)
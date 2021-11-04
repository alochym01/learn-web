CREATE TABLE albums (
	id VARCHAR PRIMARY KEY,
	title TEXT NOT NULL,
	artist TEXT NOT NULL,
	price FLOAT NOT NULL UNIQUE
);
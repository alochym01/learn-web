CREATE TABLE albums (
	id  INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL UNIQUE,
	artist TEXT NOT NULL,
	price FLOAT NOT NULL
);
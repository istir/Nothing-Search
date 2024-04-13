CREATE TABLE files (
	id	INTEGER PRIMARY KEY AUTOINCREMENT,
	name	text,
	path   text,
	thumbnail_exists	INTEGER DEFAULT 0,
	date_modified	INTEGER,
	date_created	INTEGER
);

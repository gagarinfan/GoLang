# GoLang
My first steps with GoLang

# Initialize database

CREATE TABLE IF NOT EXISTS peoples(id INTEGER PRIMARY KEY, name TEXT, surname TEXT, age INTEGER, hobby STRING, active INTEGER);

INSERT INTO peoples(id, name, surname, age, hobby, active)
VALUES(1, 'Jan', 'Kowalski', 24, 'golang', 1);
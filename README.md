# GoLang
My first steps with GoLang

# About application
This application is a simple REST app that stores people in database (here sqlite). People is golang's struct that store:
- id (primary key)
- name
- surname
- age
- hobby
- active (is active? which is int 1 as true or 0 as false)

# Initialize database

Simple SQL to create table if not exist:

```sql
CREATE TABLE IF NOT EXISTS peoples(id INTEGER PRIMARY KEY, name TEXT, surname TEXT, age INTEGER, hobby STRING, active INTEGER);
```
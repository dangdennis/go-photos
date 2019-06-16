# Installation

Make sure you're in the `/backend` directory

1. `go run main.go` to run the dev server

## Setting up postgreSQL

[Calhoun's psql guide](https://www.calhoun.io/using-postgresql-with-go/)

[Create Database](https://www.tutorialspoint.com/postgresql/postgresql_create_database.htm)

To access your database, assuming user created with or without password, run:

- `psql -U dennis googlephotos_dev`

To drop a database
- `dropdb your-db-name`

Experiment:

Creating our users and orders table
```
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT,
  email TEXT NOT NULL
)


CREATE TABLE orders (
  id SERIAL PRIMARY KEY
  user_id INT NOT NULL
  amount INT
  description TEXT
)
```



# golang-postgres-api
A RESTful API built with Golang that uses a postgresql database. The go framework used is gin.

## Running
To test the api, run `git clone` on the repository and then `docker build -t api .` in the root directory follwed by `docker-compose.yml -f docker-compose up`.
Then, use the shell in the database to run this sql:

```
CREATE TABLE users {
  id SERIAL PRIMARY KEY NOT NULL,
  age SERIAL,
  first_name VARCHAR(255),
  last_name VARCHAR(255),
  email VARCHAR(255) UNIQUE NOT NULL
);
```
Finally, rerun the container and everything should work. You are free to use this however you wish (nothing illegal though).

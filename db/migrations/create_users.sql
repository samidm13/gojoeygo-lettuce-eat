CREATE TABLE users (
user_id serial PRIMARY KEY,
first_name VARCHAR (50) NOT NULL,
last_name VARCHAR (50) NOT NULL,
email VARCHAR (355) UNIQUE NOT NULL,
password VARCHAR (355) NOT NULL,
);

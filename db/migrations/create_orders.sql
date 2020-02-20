CREATE TABLE orders (
order_id serial PRIMARY KEY,
token int,
rest_id INTEGER REFERENCES restaurants (rest_id),
order_time TIME,
user_id INTEGER REFERENCES users (user_id),
address VARCHAR (1000) NOT NULL
);

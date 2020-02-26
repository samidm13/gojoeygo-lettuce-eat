CREATE TABLE basket (
basket_id serial PRIMARY KEY,
dish_id INTEGER REFERENCES dishes (dish_id),
dish_name VARCHAR (50) NOT NULL,
dish_price DECIMAL NOT NULL,
token int,
user_id INTEGER REFERENCES users (user_id)
);

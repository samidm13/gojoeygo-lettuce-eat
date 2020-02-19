CREATE TABLE dishes (
dish_id serial PRIMARY KEY,
dish_name VARCHAR (50) NOT NULL,
dish_price MONEY NOT NULL,
description VARCHAR (250),
rest_id INTEGER REFERENCES restaurants (rest_id)
);

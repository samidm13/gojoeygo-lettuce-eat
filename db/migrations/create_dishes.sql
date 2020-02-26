CREATE TABLE dishes (
dish_id serial PRIMARY KEY,
dish_name VARCHAR (50) NOT NULL,
dish_price DECIMAL NOT NULL,
description VARCHAR (250),
rest_id INTEGER REFERENCES restaurants (rest_id)
url VARCHAR (250),
);

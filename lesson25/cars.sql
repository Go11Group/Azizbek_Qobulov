CREATE TABLE car (
    id INT PRIMARY KEY,
    name VARCHAR(100)
);
CREATE TABLE user (
    id INT PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE Cars (
    car_id INT,
    user_id INT,
    FOREIGN KEY (car_id) REFERENCES car(id),
    FOREIGN KEY (user_id) REFERENCES user(id)
);


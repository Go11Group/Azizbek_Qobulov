CREATE TABLE author (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    birthdate DATE
);

CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    genre VARCHAR(50),
    publish_date DATE,
    author_id INT,
    FOREIGN KEY (author_id) REFERENCES author (id)
);

--Malumotlar gugldan oldim
INSERT INTO author (name, birthdate) VALUES
('George Orwell', '1903-06-25'),
('J.K. Rowling', '1965-07-31'),
('J.R.R. Tolkien', '1892-01-03'),
('F. Scott Fitzgerald', '1896-09-24'),
('Jane Austen', '1775-12-16'),
('Mark Twain', '1835-11-30'),
('Ernest Hemingway', '1899-07-21');


INSERT INTO book (title, genre, publish_date, author_id) VALUES
('1984', 'Dystopian', '1949-06-08', 1),
('Animal Farm', 'Political Satire', '1945-08-17', 1),
('Harry Potter and the Philosopher\'s Stone', 'Fantasy', '1997-06-26', 2),
('Harry Potter and the Chamber of Secrets', 'Fantasy', '1998-07-02', 2),
('The Lord of the Rings: The Fellowship of the Ring', 'Fantasy', '1954-07-29', 3),
('The Lord of the Rings: The Two Towers', 'Fantasy', '1954-11-11', 3),
('The Great Gatsby', 'Novel', '1925-04-10', 4),
('Pride and Prejudice', 'Romantic', '1813-01-28', 5),
('Sense and Sensibility', 'Romantic', '1811-10-30', 5),
('Adventures of Huckleberry Finn', 'Adventure', '1884-12-10', 6),
('The Adventures of Tom Sawyer', 'Adventure', '1876-06-10', 6),
('The Old Man and the Sea', 'Novel', '1952-09-01', 7),
('A Farewell to Arms', 'Novel', '1929-09-27', 7),
('The Sun Also Rises', 'Novel', '1926-10-22', 7),
('To Have and Have Not', 'Novel', '1937-10-15', 7);

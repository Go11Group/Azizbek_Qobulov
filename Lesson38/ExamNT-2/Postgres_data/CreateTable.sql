-- Active: 1717086991958@@127.0.0.1@5432@master
CREATE TABLE Users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    birthday VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE Courses (
    course_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);
CREATE TABLE lessons (
    lesson_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID REFERENCES courses(course_id),
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);
CREATE TABLE enrollments (
    enrollment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(user_id),
    course_id UUID REFERENCES courses(course_id),
    enrollment_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

--Insert Users
INSERT INTO Users (name, email, birthday, password)
VALUES 
('John Doe', 'john.doe@example.com', '01/01/1980', 'password1'),
('Jane Smith', 'jane.smith@example.com', '02/02/1985', 'password2'),
('Alice Johnson', 'alice.johnson@example.com', '03/03/1990', 'password3'),
('Bob Brown', 'bob.brown@example.com', '04/04/1975', 'password4'),
('Charlie Davis', 'charlie.davis@example.com', '05/05/1982', 'password5'),
('Diana Evans', 'diana.evans@example.com', '06/06/1987', 'password6'),
('Frank Green', 'frank.green@example.com', '07/07/1993', 'password7'),
('Hannah Harris', 'hannah.harris@example.com', '08/08/1978', 'password8'),
('Ivy Jackson', 'ivy.jackson@example.com', '09/09/1983', 'password9'),
('Jack King', 'jack.king@example.com', '10/10/1988', 'password10');

--Insert Courses
INSERT INTO Courses (title, description)
VALUES 
('Math', 'Introduction to mathematics'),
('Science', 'Fundamentals of science'),
('History', 'Introduction to world history'),
('Geography', 'Introduction to world geography'),
('Art', 'Introduction to art'),
('Music', 'Introduction to music'),
('Philosophy', 'Introduction to philosophy'),
('Religion', 'Introduction to religion'),
('Business', 'Introduction to business'),
('Law', 'Introduction to law');

--insert lessons
INSERT INTO lessons (course_id, title, content)
VALUES
((SELECT course_id FROM courses WHERE title = 'Math'), 'Introduction to Mathematics', 'Learn the basics of mathematics'),
((SELECT course_id FROM courses WHERE title = 'Science'), 'Introduction to Science', 'Learn the basics of science'),
((SELECT course_id FROM courses WHERE title = 'History'), 'Introduction to History', 'Learn the basics of history'),
((SELECT course_id FROM courses WHERE title = 'Geography'), 'Introduction to Geography', 'Learn the basics of geography'),
((SELECT course_id FROM courses WHERE title = 'Art'), 'Introduction to Art', 'Learn the basics of art'),
((SELECT course_id FROM courses WHERE title = 'Music'), 'Introduction to Music', 'Learn the basics of music'),
((SELECT course_id FROM courses WHERE title = 'Philosophy'), 'Introduction to Philosophy', 'Learn the basics of philosophy'),
((SELECT course_id FROM courses WHERE title = 'Religion'), 'Introduction to Religion', 'Learn the basics of religion'),
((SELECT course_id FROM courses WHERE title = 'Business'), 'Introduction to Business', 'Learn the basics of business'),
((SELECT course_id FROM courses WHERE title = 'Law'), 'Introduction to Law', 'Learn the basics of law');

--insert enrollments
INSERT INTO enrollments (user_id, course_id, enrollment_date)
VALUES
((SELECT user_id FROM users WHERE name = 'John Doe'), (SELECT course_id FROM courses WHERE title = 'Math'), '2022-01-01'),
((SELECT user_id FROM users WHERE name = 'Jane Smith'), (SELECT course_id FROM courses WHERE title = 'Science'), '2022-02-02'),
((SELECT user_id FROM users WHERE name = 'Alice Johnson'), (SELECT course_id FROM courses WHERE title = 'History'), '2022-03-03'),
((SELECT user_id FROM users WHERE name = 'Bob Brown'), (SELECT course_id FROM courses WHERE title = 'Geography'), '2022-04-04'),
((SELECT user_id FROM users WHERE name = 'Charlie Davis'), (SELECT course_id FROM courses WHERE title = 'Art'), '2022-05-05'),
((SELECT user_id FROM users WHERE name = 'Diana Evans'), (SELECT course_id FROM courses WHERE title = 'Music'), '2022-06-06'),
((SELECT user_id FROM users WHERE name = 'Frank Green'), (SELECT course_id FROM courses WHERE title = 'Philosophy'), '2022-07-07'),
((SELECT user_id FROM users WHERE name = 'Hannah Harris'), (SELECT course_id FROM courses WHERE title = 'Religion'), '2022-08-08'),
((SELECT user_id FROM users WHERE name = 'Ivy Jackson'), (SELECT course_id FROM courses WHERE title = 'Business'), '2022-09-09'),
((SELECT user_id FROM users WHERE name = 'Jack King'), (SELECT course_id FROM courses WHERE title = 'Law'), '2022-10-10');

TRUNCATE table users, courses, lessons, enrollments
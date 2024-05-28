CREATE TABLE Student (
    student_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(100)
);

CREATE TABLE student_course (
    id int PRIMARY KEY,
    student_id uuid,
    course_id INT,
    FOREIGN KEY (student_id) REFERENCES Student(student_id),
    FOREIGN KEY (course_id) REFERENCES course(id)
);



CREATE TABLE course (
    id INT PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE grade (
    id serial PRIMARY KEY,
    student_course_id INT,
    grade int,
    FOREIGN KEY (student_course_id) REFERENCES student_course(id)
);

INSERT INTO student (id, name) VALUES 
(1, 'Alice'),
(2, 'Bob'),
(3, 'Charlie'),
(4, 'David'),
(5, 'Eve'),
(6, 'Frank'),
(7, 'Grace'),
(8, 'Hank'),
(9, 'Ivy'),
(10, 'Jack');

INSERT INTO course (id, name) VALUES 
(1, 'Math'),
(2, 'Science'),
(3, 'History'),
(4, 'Literature'),
(5, 'Physics'),
(6, 'Chemistry'),
(7, 'Biology'),
(8, 'Geography'),
(9, 'Art'),
(10, 'Music');

INSERT INTO student_course (student_id, course_id) VALUES 
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 6),
(7, 7),
(8, 8),
(9, 9),
(10, 10),
(1, 2),
(2, 3),
(3, 4),
(4, 5),
(5, 6),
(6, 7),
(7, 8),
(8, 9),
(9, 10),
(10, 1);

INSERT INTO grade (student_course_id, grade) VALUES 
(1, 5),
(2, 4),
(3, 3),
(4, 5),
(5, 4),
(6, 3),
(7, 5),
(8, 4),
(9, 3),
(10, 5),
(11, 4),
(12, 3),
(13, 5),
(14, 4),
(15, 3),
(16, 5),
(17, 4),
(18, 3),
(19, 5),
(20, 4);

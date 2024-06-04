CREATE TABLE Student (
    student_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(100)
    age int
);

CREATE TABLE student_course (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    student_id uuid,
    course_id INT,
    FOREIGN KEY (student_id) REFERENCES Student(student_id),
    FOREIGN KEY (course_id) REFERENCES course(id)
);



CREATE TABLE course (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100)
);

CREATE TABLE grade (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    student_course_id INT,
    mark int,
    FOREIGN KEY (student_course_id) REFERENCES student_course(id)
);

-- har bir kursda eng yaxshi o'quvchini topish
SELECT c.name AS course, s.name AS student_name, max_grades.max_grade
FROM (
    SELECT sc.course_id, MAX(g.mark) AS max_grade
    FROM student_course sc
    JOIN grade g ON sc.id = g.student_course_id
    GROUP BY sc.course_id
) AS max_grades
JOIN student_course sc ON max_grades.course_id = sc.course_id
JOIN grade g ON sc.id = g.student_course_id AND g.mark = max_grades.max_grade
JOIN student s ON sc.student_id = s.id
JOIN course c ON sc.course_id = c.id;

-- barcha kurslar orasida o'rtacha ball topish
select c.name as course, round(avg(g.mark),2) as average_mark from course c
join student_course sc on c.id = sc.course_id
join grade g on sc.id = g.student_course_id
group by c.name;

-- barcha kurslar orasida eng yosh talabani topish
select c.name as course, s.name as youngest_student, min(s.age) from course c
join student_course sc on c.id = sc.course_id
join student s on sc.student_id = s.id
join (
    select course_id, min(age) as min_age from student_course sc
    join student s on sc.student_id = s.id
    group by course_id
) as y_per_c on sc.course_id = y_per_c.course_id and s.age = y_per_c.min_age
group by c.name, s.name;

-- eng yuqori o'rtacha ballga ega kursni topish
select course, round(max(avgMark),2) as highest_mark
from (
    select c.name as course, avg(g.mark) as avgMark from course c
    join student_course sc on sc.course_id = c.id
    join grade g on g.student_course_id = sc.id
    group by c.name) as avg_marks
group by course
order by highest_mark desc
limit 1;
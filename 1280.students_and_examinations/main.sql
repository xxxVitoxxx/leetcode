-- the SQL query first perform a cross join between Students and Subjects tables,
-- to create a result set that includes all possible combinations of students and subjects.
-- then left joins this result set with the Examinations table,
-- using both the student_id and subject_name columns.
-- to count the number of times each student attended each exam.

-- the COUNT() function is used to count the number of rows in Examinations table that
-- match the given student and subject.
-- because the Examinations table may contain duplicates, 
-- we use the DISTINCT keyword inside the COUNT() function to only count distinct rows.
-- the COALESCE() function is used to replace NULL values with 0, so that students who 
-- did not attend a particular exams are still included in the result set.

SELECT
    s.student_id,
    s.student_name,
    sub.subject_name,
    COUNT(e.student_id) AS attended_exams
FROM
    Students s
CROSS JOIN
    Subjects sub
LEFT JOIN
    Examinations e
ON
    s.student_id = e.student_id AND sub.subject_name = e.subject_name
GROUP BY
    s.student_id, s.student_name, sub.subject_name
ORDER BY
    s.student_id, sub.subject_name;

-- solution2
SELECT
    s.student_id,
    s.student_name,
    sub.subject_name,
    COALESCE(e.attended_exams, 0) AS attended_exams
FROM
    Students s
CROSS JOIN
    Subjects sub
LEFT JOIN (
    SELECT
        student_id,
        subject_name,
        COUNT(*) AS attended_exams
    FROM
        Examinations
    GROUP BY
        student_id, subject_name
) e USING(student_id, subject_name) -- s.student_id = e.student_id AND sub.subject_name = e.subject_name
ORDER BY
    s.student_id, sub.subject_name;

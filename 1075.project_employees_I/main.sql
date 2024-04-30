-- solution1
SELECT
    p.project_id,
    ROUND(SUM(e.experience_years) / COUNT(p.project_id), 2) AS average_years
FROM
    Project p
JOIN
    Employee e
ON
    p.employee_id = e.employee_id
WHERE
    e.experience_years IS NOT NULL
GROUP BY
    project_id;

-- solution2
SELECT
    p.project_id,
    ROUND(AVG(experience_years), 2) AS average_years
FROM
    Project p
-- INNER JOIN / LEFT JOIN
JOIN
    Employee e
ON
    p.employee_id = e.employee_id
GROUP BY
    project_id;

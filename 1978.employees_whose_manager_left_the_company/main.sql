-- subquery and IN clause
-- solution1
SELECT
    employee_id
FROM
    Employees
WHERE
    manager_id IS NOT NULL
AND
    salary < 30000
AND
    manager_id NOT IN (
        SELECT
            employee_id
        FROM
            Employees
    )
ORDER BY
    employee_id;

-- LEFT JOIN
-- solution2
SELECT
    e1.employee_id
FROM
    Employees e1
LEFT JOIN
    Employees e2
ON
    e1.manager_id = e2.employee_id
WHERE
    e1.salary < 30000
AND
    e1.manager_id IS NOT NULL
AND
    e2.employee_id IS NULL
ORDER BY
    e1.employee_id;

-- self JOIN
-- solution1
SELECT
    mgr.employee_id,
    mgr.name,
    COUNT(emp.employee_id) AS reports_count,
    ROUND(AVG(emp.age)) AS average_age
FROM
    Employees emp
JOIN
    Employees mgr
ON
    emp.reports_to = mgr.employee_id
GROUP BY
    mgr.employee_id
ORDER BY
    mgr.employee_id;

-- correlated subquery
-- solution2
SELECT
    reports_to AS employee_id,
    (
        SELECT
            name
        FROM
            Employees e1
        WHERE
            e.reports_to = e1.employee_id
    ) AS name,
    COUNT(reports_to) AS reports_count,
    ROUND(AVG(age)) AS average_age
FROM
    Employees e
GROUP BY
    reports_to
HAVING
    reports_to IS NOT NULL
ORDER BY
    reports_to;

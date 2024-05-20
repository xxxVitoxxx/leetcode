-- combines two distinct sets of logic using the UNION operator
-- solution1
SELECT
    employee_id,
    department_id
FROM
    Employee
WHERE
    primary_flag = 'Y'
UNION
SELECT
    employee_id,
    department_id
FROM
    Employee
GROUP BY
    employee_id
HAVING
    COUNT(employee_id) = 1;


-- window function
-- solution2
SELECT
    employee_id,
    department_id
FROM (
    SELECT
        *,
        COUNT(employee_id) OVER (
            PARTITION BY
                employee_id
        ) AS cnt
    FROM
        Employee
) t
WHERE
    t.cnt = 1 OR t.primary_flag = 'Y';

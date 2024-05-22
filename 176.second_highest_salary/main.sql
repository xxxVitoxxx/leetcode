-- use subquery and LIMIT clause
-- solution1
SELECT (
    SELECT
        DISTINCT salary
    FROM
        Employee
    ORDER BY
        salary DESC
    LIMIT 1 OFFSET 1
) AS SecondHighestSalary;

-- use subquery and MAX() function
-- solution2
SELECT
    MAX(salary) AS SecondHighestSalary
FROM
    Employee
WHERE
    salary < (
        SELECT
            MAX(salary)
        FROM
            Employee
    );

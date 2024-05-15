-- return the first n rows using correlated subquery
-- solution1
-- we can build a correlated subquery to identify the top N records from more than one category.
-- since the correlated subquery is dependent on the main query,
-- the idea behind this approach is to compare the values between in 
-- the main query and the subquery, so that in the subquery,
-- at most N-1 salaries can be greater than each selected salary from the main query.

-- to do this, we first build the main query. in the main query,
-- we can also join the table Department to the table Employee on departmentId to
-- get the name of the departments and rename the columns as requested by the final output.
-- SELECT
--     d.name AS Department,
--     e1.name AS Employee,
--     e1.salary AS Salary
-- FROM
--     Employee e1
-- JOIN
--     Department d
-- ON
--     e1.departmentId = d.id

-- in the correlated subquery, we select the number of salary from the same table Employee.
-- to compare the salaries between in the main query and the subquery,
-- we make sure the department is the same from both queries,
-- but the salary from the subquery is always bigger than the salary from the main query.
-- SELECT
--     COUNT(DISTINCT e1.salary)
-- FROM
--     Employee e2
-- WHERE
--     e2.salary > e1.salary AND e1.departmentId = e2.departmentId

-- since we need to identify the top three earners in the main query,
-- and the subquery always had larger salaries than the salary from the main query,
-- the maximum count of the larger salaries in the subquery is two.
-- we add this criteria as a filter to the main query.
SELECT
    d.name AS Department,
    e1.name AS Employee,
    e1.salary AS Salary
FROM
    Employee e1
JOIN
    Department d
ON
    e1.departmentId = d.id
WHERE
    3 > (
        SELECT
            COUNT(DISTINCT e2.salary)
        FROM
            Employee e2
        WHERE
            e2.salary > e1.salary AND e2.departmentId = e1.departmentId
    )
ORDER BY
    Department;

-- return the fist n rows using DENSE_RANK()
-- solution2
-- unlike the previous approach that utilized a correlated subquery,
-- in this approach, we sorted the salaries in descending order,
-- ranked employees based on their salaries within the department,
-- and selected only the first 3 employees for the final output.

-- we first create a subquery or CTE to rank the employee.
-- since the definition of a high earner is the employee who
-- has a salary in the top three unique salaries for the department,
-- we can use the function DENSE_RANK() to avoid the scenario that employees
-- from the same department make the same amount of salary.
-- in the step, we can also join the table Department on departmentId to
-- get the name of the departments and rename the columns for the final output.
-- WITH employee_department AS (
--     SELECT
--         d.id,
--         d.name AS Department,
--         salary AS Salary,
--         e.name AS Employee,
--         DENSE_RANK() OVER (
--             PARTITION BY
--                 d.id
--             ORDER BY
--                 salary DESC
--         ) AS rnk
--     FROM
--         Department d
--     JOIN
--         Employee e
--     ON
--         d.id = e.departmentId
-- )

-- now, each employee has a rank based on the salary in a descending order for each department.
-- | id | Department | Salary | Employee | rnk |
-- | -- | ---------- | ------ | -------- | --- |
-- | 1  | IT         | 90000  | Max      | 1   |
-- | 1  | IT         | 85000  | Joe      | 2   |
-- | 1  | IT         | 85000  | Randy    | 2   |
-- | 1  | IT         | 70000  | Will     | 3   |
-- | 1  | IT         | 69000  | Janet    | 4   |
-- | 2  | Sales      | 80000  | Henry    | 1   |
-- | 2  | Sales      | 60000  | Sam      | 2   |

-- with the rank, we can select the earner.
-- we can add the filter to select employee that have a rank smaller than or
-- equal to 3 in the main query.
WITH employee_department AS (
    SELECT
        d.id,
        d.name AS Department,
        e.name AS Employee,
        salary AS Salary,
        DENSE_RANK() OVER (
            PARTITION BY
                d.id
            ORDER BY
                salary DESC
        ) AS rnk
    FROM
        Department d
    JOIN
        Employee e
    ON
        d.id = e.departmentId
)

SELECT
    Department,
    Employee,
    Salary
FROM
    employee_department
WHERE
    rnk <= 3;

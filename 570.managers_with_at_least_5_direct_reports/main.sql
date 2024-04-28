-- using IN clause with subquery
-- solution1
SELECT
    name
FROM
    Employee
WHERE
    id IN (
        SELECT
            managerId
        FROM
            Employee
        GROUP BY
            managerId
        HAVING  COUNT(*) >= 5
    )

-- using GROUP BY and JOIN
-- solution2
-- the subquery selects the manager ids from the Employee and groups them based on the number of reports they have.
-- the HAVING clause filters out the managerId having a count of reports greater than or equal to 5.

-- next, we perform join between the table from this subquery and the Employee table to obtain all the manager names.
SELECT
    name
FROM
    Employee
JOIN (
    SELECT
        managerId
    FROM
        Employee
    GROUP BY
        managerId
    HAVING COUNT(managerId) >= 5
) e
ON
    id = e.managerId;

-- solution3
SELECT
    e2.name
FROM
    Employee e1
JOIN
    Employee e2
ON
    e1.managerId = e2.id
GROUP BY
    e1.managerId
HAVING COUNT(*) >= 5;

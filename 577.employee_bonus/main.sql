-- using LEFT JOIN and WHERE clause
-- solution1
SELECT
    e.name,
    b.bonus
FROM
    Employee e
LEFT JOIN
    Bonus b
ON
    e.empID = b.empID
WHERE
    -- IFNULL(b.bonus, 0) < 1000;
    b.bonus IS NULL OR b.bonus < 1000;

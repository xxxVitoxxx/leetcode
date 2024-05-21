-- use WHERE clause
-- solution1
SELECT
    DISTINCT l1.num AS ConsecutiveNums
FROM
    Logs l1,
    Logs l2,
    Logs l3
WHERE
    l1.id = l2.id-1 AND
    l2.id = l3.id-1 AND
    l1.num = l2.num AND
    l2.num = l3.num;

-- use JOIN
-- solution2
SELECT
    DISTINCT l1.num AS ConsecutiveNums
FROM
    Logs l1
JOIN
    Logs l2
ON
    l1.id = l2.id-1
JOIN
    Logs l3
ON
    l1.id = l3.id-2
WHERE
    l1.num = l2.num AND l2.num = l3.num;

-- use EXISTS and subquery
-- solution3
SELECT
    DISTINCT l1.num AS ConsecutiveNums
FROM
    Logs l1
WHERE EXISTS (
    SELECT
        1
    FROM
        Logs l2
    WHERE
        l1.id+1 = l2.id AND l1.num = l2.num AND EXISTS (
            SELECT
                1
            FROM
                Logs l3
            WHERE
                l2.id+1 = l3.id AND l2.num = l3.num
        )
);

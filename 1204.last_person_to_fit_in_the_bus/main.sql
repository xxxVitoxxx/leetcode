-- solution1
SELECT
    person_name
FROM
    Queue q
WHERE
    1000 >= (
        SELECT
            SUM(weight)
        FROM
            Queue
        WHERE
            q.turn >= turn
    )
ORDER BY
    turn DESC
LIMIT 1;

-- use window function
-- solution2
SELECT
    person_name
FROM (
    SELECT
        person_name,
        turn,
        SUM(weight) OVER (
            ORDER BY
                turn
        ) AS sum
    FROM
        Queue
) t
WHERE
    t.sum <= 1000
ORDER BY
    t.turn DESC
LIMIT 1;

-- subqueries and multi-value use of the IN comparison operator
-- solution1
-- we only need to use COUNT(*) in the ROUND() function as opposed to COUNT(player_id)
-- since player_id and event_date is the primary key of the table
SELECT
    ROUND(COUNT(*) / (
        SELECT
            COUNT(DISTINCT player_id)
        FROM
            Activity
    ), 2) AS fraction
FROM
    Activity
WHERE
    (player_id, DATE_SUB(event_date, INTERVAL 1 DAY)) IN (
        SELECT
            player_id,
            MIN(event_date)
        FROM
            Activity
        GROUP BY
            player_id
    );

-- CTEs and RIGHT JOIN
-- solution2
WITH a1 AS (
    SELECT
        player_id,
        MIN(event_date) as event_date
    FROM
        Activity
    GROUP BY
        player_id
)

SELECT
    ROUND(COUNT(a2.event_date) / COUNT(*), 2) AS fraction
FROM
    Activity a2
RIGHT JOIN
    a1
ON
    a1.player_id = a2.player_id
AND
    DATE_SUB(a2.event_date, INTERVAL 1 DAY) = a1.event_date;

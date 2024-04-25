-- using subquery
-- solution1
SELECT
    w1.id
FROM
    Weather w1
WHERE
    w1.temperature > (
        SELECT
            w2.temperature
        FROM
            Weather w2
        WHERE
            w2.recordDate = DATE_SUB(w1.recordDate, INTERVAL 1 DAY)
    );

-- solution2
SELECT
    w1.id
FROM
    Weather w1, Weather w2
WHERE
    DATEDIFF(w1.recordDate, w2.recordDate) = 1
AND
    w1.temperature > w2.temperature;

-- using JOIN
-- solution3
SELECT
    w1.id
FROM
    Weather w1
JOIN
    Weather w2
ON
    DATEDIFF(w1.recordDate, w2.recordDate) = 1
WHERE
    w1.temperature > w2.temperature;

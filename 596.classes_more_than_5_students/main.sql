-- solution1
SELECT
    class
FROM
    Courses
GROUP BY
    class
HAVING
    COUNT(*) >= 5;

-- solution2
SELECT
    class
FROM (
    SELECT
        class,
        COUNT(*) AS cnt
    FROM
        Courses
    GROUP BY
        class
) t
WHERE
    cnt >= 5;

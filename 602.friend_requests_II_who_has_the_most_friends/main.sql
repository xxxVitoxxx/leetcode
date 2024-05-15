-- combining tables using UNION ALL and finding the top values using ORDER BY + LIMIT
-- solution1
WITH all_ids AS (
    SELECT
        requester_id AS id
    FROM
        RequestAccepted
    UNION ALL
    SELECT
        accepter_id AS id
    FROM
        RequestAccepted
)

SELECT
    id,
    COUNT(*) AS num
FROM
    all_ids
GROUP BY
    id
ORDER BY
    num DESC
LIMIT 1

-- combining tables using UNION ALL and finding top values using RANK() and window function
-- solution2
WITH all_ids AS (
    SELECT
        requester_id AS id
    FROM
        RequestAccepted
    UNION ALL
    SELECT
        accepter_id AS id
    FROM
        RequestAccepted
)

SELECT
    id,
    num
FROM (
    SELECT
        id,
        COUNT(id) AS num,
        RANK() OVER(
            ORDER BY
                COUNT(id) DESC
        ) AS rnk
    FROM
        all_ids
    GROUP BY
        id
) t
WHERE
    rnk = 1;

-- window function
-- solution1
WITH cte AS (
    SELECT
        visited_on,
        SUM(amount) AS daily_amount
    FROM
        Customer
    GROUP BY
        visited_on
)
SELECT
    visited_on,
    -- 因為 cte 已將 visited_on group by ，所以只基於前六行就可以知道前六天到當天的總金額
    SUM(daily_amount) OVER (
        ORDER BY
            visited_on
        ROWS
            BETWEEN 6 PRECEDING AND CURRENT ROW
    ) AS amount,
    ROUND(AVG(daily_amount) OVER (
            ORDER BY
                visited_on
            ROWS
                BETWEEN 6 PRECEDING AND CURRENT ROW
        ),2) AS average_amount
FROM
    cte
LIMIT
    1000 OFFSET 6;

-- solution2
-- subquery result
-- SELECT
--     DISTINCT visited_on,
--     SUM(amount) OVER(
--         ORDER BY
--             visited_on
--         RANGE
--             BETWEEN INTERVAL 6 DAY PRECEDING AND CURRENT ROW
--     ) AS amount,
--     MIN(visited_on) OVER() AS 1st_date
-- FROM
--     Customer
-- | visited_on | amount | 1st_date   |
-- | ---------- | ------ | ---------- |
-- | 2019-01-01 | 100    | 2019-01-01 |
-- | 2019-01-02 | 210    | 2019-01-01 |
-- | 2019-01-03 | 330    | 2019-01-01 |
-- | 2019-01-04 | 460    | 2019-01-01 |
-- | 2019-01-05 | 570    | 2019-01-01 |
-- | 2019-01-06 | 710    | 2019-01-01 |
-- | 2019-01-07 | 860    | 2019-01-01 |
-- | 2019-01-08 | 840    | 2019-01-01 |
-- | 2019-01-09 | 840    | 2019-01-01 |
-- | 2019-01-10 | 1000   | 2019-01-01 |
-- 這邊使用沒有條件的 OVER() 表示要在整個查詢結果集中找到 visited_on 最小值。
-- 這種方法使得 MIN(visited_on) 變成一個常數，對於每一行返回的值都相同，即整個查詢結果中最早的日期
-- 如果不使用 OVER() 的話， MIN(visited_on) 會被視為一個聚合函數，只會回傳一個最小日期，通常與 GROUP BY 一起使用

WITH cte AS (
    SELECT
        DISTINCT visited_on,
        -- 計算當前日期到前六天的所有金額
        SUM(amount) OVER(
            ORDER BY
                visited_on
            RANGE
                BETWEEN INTERVAL 6 DAY PRECEDING AND CURRENT ROW
        ) AS amount,
        MIN(visited_on) OVER() AS first_visited_date
    FROM
        Customer
)

SELECT
    visited_on,
    amount,
    ROUND(amount/7, 2) AS average_amount
FROM
    cte
WHERE
    visited_on >= first_visited_date+6;

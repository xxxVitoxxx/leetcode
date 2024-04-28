-- solution1
-- left join Signups table with Confirmations table. 
-- to display confirmation_rate, we need to find average using aggregation function AVG(),
-- here using if clause we specified that if action is confirmed, we count it as 1 and else
-- if null or timeout we will count it as 0.
SELECT
    s.user_id,
    ROUND(AVG(IF(c.action = 'confirmed', 1, 0)), 2) AS confirmation_rate
FROM
    Signups s
LEFT JOIN
    Confirmations c
ON
    s.user_id = c.user_id
GROUP BY
    s.user_id;

-- solution2
SELECT
    s.user_id,
    ROUND(SUM(CASE WHEN c.action = 'confirmed' THEN 1 ELSE 0 END) / COUNT(s.user_id), 2) AS confirmation_rate
FROM
    Signups s
LEFT JOIN
    Confirmation c
ON
    s.user_id = c.user_id
GROUP BY
    s.user_id;

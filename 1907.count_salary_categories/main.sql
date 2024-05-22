-- use subquery
-- solution1
SELECT
    'Low Salary' AS category,
    COUNT(income) AS accounts_count
FROM
    Accounts
WHERE
    income < 20000
UNION ALL
SELECT
    'Average Salary' AS category,
    COUNT(income) AS accounts_count
FROM
    Accounts
WHERE
    income >= 20000 AND income <= 50000
UNION ALL
SELECT
    'High Salary' AS category,
    COUNT(income) AS accounts_count
FROM
    Accounts
WHERE
    income > 50000;

-- use CASE WHEN
-- solution2
SELECT
    'Low Salary' AS category,
    SUM(CASE WHEN income < 20000 THEN 1 ELSE 0 END) AS accounts_count
FROM
    Accounts
UNION ALL
SELECT
    'Average Salary' AS category,
    SUM(CASE WHEN income >= 20000 AND income <= 50000 THEN 1 ELSE 0 END) AS accounts_count
FROM
    Accounts
UNION ALL
SELECT
    'High Salary' AS category,
    SUM(CASE WHEN income > 50000 THEN 1 ELSE 0 END) AS accounts_count
FROM
    Accounts;

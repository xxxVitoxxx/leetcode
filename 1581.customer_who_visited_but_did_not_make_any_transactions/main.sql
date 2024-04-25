-- remove records using NOT IN
-- solution1
SELECT
    customer_id,
    COUNT(visit_id) AS count_no_trans
FROM
    Visits
WHERE
    visit_id NOT IN (
        SELECT
            visit_id
        FROM
            Transactions
    )
GROUP BY
    customer_id;

-- remove record using LEFT JOIN and IS NULL
-- solution2
SELECT
    customer_id,
    COUNT(*) AS count_no_trans
FROM
    Visits
LEFT JOIN
    Transactions
ON
    Visits.visit_id = Transactions.visit_id
WHERE
    Transactions.visit_id IS NULL
GROUP BY
    customer_id;

-- solution3
SELECT
    customer_id,
    COUNT(Visits.visit_id) AS count_no_trans
FROM
    Visits
LEFT JOIN
    Transactions
ON
    Visits.visit_id = Transactions.visit_id
WHERE
    Transactions.visit_id IS NULL
GROUP BY
    customer_id;

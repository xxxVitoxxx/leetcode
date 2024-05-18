-- filtering from minimum values subquery
-- solution1
SELECT
    product_id,
    year AS first_year,
    quantity,
    price
FROM
    Sales
WHERE
    (product_id, year) IN (
        SELECT
            product_id,
            MIN(year) AS year
        FROM
            Sales
        GROUP BY
            product_id
    );

-- using INNER JOIN
-- solution2
WITH cte AS (
    SELECT
        product_id,
        MIN(year) AS min_year
    FROM
        Sales
    GROUP BY
        product_id
)

SELECT
    s.product_id,
    s.year AS first_year,
    s.quantity,
    s.price
FROM
    Sales s
INNER JOIN
    cte
ON
    s.product_id = cte.product_id AND s.year = cte.min_year;

-- use UNION ALL
-- solution1
-- 1. group the table with the product_id field and find the first changed date
-- over 2019-08-16 by using MIN aggregation function on HAVING clause.
-- 2. group the table with the product_id again, and fine the product_id field and
-- the last change date until 2019-08-16
-- 3. fine the last changed new_price field with the last changed date.
-- 4. union the two tables by using UNION ALL

-- we know there are no duplicate tuples when we union the two separated tables because
-- we get one field using GROUP BY for each query. thus, it would be better to use
-- UNION ALL instead of UNION for performance.
SELECT
    product_id,
    10 AS price
FROM
    Products
GROUP BY
    product_id
HAVING
    MIN(change_date) > '2019-08-16'
UNION ALL
SELECT
    product_id,
    new_price AS price
FROM
    Products
WHERE
    (product_id, change_date) IN (
        SELECT
            product_id,
            MAX(change_date)
        FROM
            Products
        WHERE
            change_date <= '2019-08-16'
        GROUP BY
            product_id
    );

-- use window function
-- solution2
-- 1. filter the table where the value of the change_date field is under the given date 2019-08-16.
-- 2. get the last changed price using FIRST_VALUE for each product_id.
SELECT
    DISTINCT p.product_id,
    IFNULL(price, 10) AS price
FROM
    Products p
LEFT JOIN (
    SELECT
        DISTINCT product_id,
        FIRST_VALUE(new_price) OVER (
            PARTITION BY
                product_id
            ORDER BY
                change_date DESC
        ) AS price
    FROM
        Products
    WHERE
        change_date <= '2019-08-16'
) AS LastChangePrice
USING (product_id);

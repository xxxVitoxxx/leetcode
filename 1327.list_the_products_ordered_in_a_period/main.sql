-- solution1
SELECT
    p.product_name,
    t.unit
FROM
    Products p
INNER JOIN (
    SELECT
        product_id,
        SUM(unit) AS unit
    FROM
        Orders
    WHERE
        DATE_FORMAT(order_date, '%Y-%m') = '2020-02'
    GROUP BY
        product_id
) t
USING(product_id)
WHERE
    t.unit >= 100;

-- solution2
SELECT
    p.product_name,
    SUM(o.unit) AS unit
FROM
    Products p
INNER JOIN
    Orders o
ON
    p.product_id = o.product_id AND DATE_FORMAT(o.order_date, '%Y-%m') = '2020-02'
GROUP BY
    p.product_id
HAVING
    SUM(o.unit) >= 100;

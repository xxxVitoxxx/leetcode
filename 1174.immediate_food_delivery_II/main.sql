-- solution1
SELECT
    ROUND(AVG(order_date = customer_pref_delivery_date) * 100, 2) AS immediate_percentage
FROM
    Delivery
WHERE
    (customer_id, order_date) IN (
        SELECT
            customer_id,
            MIN(order_date)
        FROM
            Delivery
        GROUP BY
            customer_id
    );

-- solution2
WITH first_order AS (
    SELECT
        customer_id,
        MIN(order_date) AS first_order_date
    FROM
        Delivery
    GROUP BY
        customer_id
)

SELECT
    ROUND(SUM(CASE WHEN d.order_date = d.customer_pref_delivery_date THEN 1 ELSE 0 END) / COUNT(d.customer_id) * 100, 2) AS immediate_percentage
FROM
    Delivery d
JOIN
    first_order f
ON
    d.customer_id = f.customer_id AND d.order_date = f.first_order_date;

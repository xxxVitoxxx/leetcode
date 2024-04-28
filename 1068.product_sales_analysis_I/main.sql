SELECT
    p.product_name, s.year, s.price
FROM
    Sales s
JOIN
    Product p
ON
    p.product_id = s.product_id;

-- it's important to note that these two tables are related through
-- the product_id column. therefore, we will join these two tables
-- using this column. this way, we will be able to present information
-- from both tables simultaneously.

-- group and aggregation of strings
-- solution1
-- we group the date by sell_date field, to get the num_sold field,
-- we use COUNT(DISTINCT product) to count the number of unique products sold on each sell date.
-- we can use the function GROUP_CONCAT to combine multiple values from multiple rows into a single string.
-- the following shows the syntax of the GROUP_CONCAT function:
-- GROUP_CONCAT(
--     DISTINCT expression1
--     ORDER BY expression2
--     SEPARATOR sep
-- );

-- the keyword DISTINCT ensure that each name in the column expression1 is included only once
-- in the concatenated string.
-- note that we need to sort unique names in ascending order, which is the default order,
-- so we can omit the parameter expression2. the keyword SEPARATOR specifies that
-- the product names should be separated by sep
SELECT
    sell_date,
    COUNT(DISTINCT product) AS num_sold,
    GROUP_CONCAT(DISTINCT product SEPARATOR ',') AS products
FROM
    Activities
GROUP BY
    sell_date
ORDER BY
    sell_date;

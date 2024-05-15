-- creating filters in subqueries
-- solution1
-- in this approach, we identify the qualified values(e.g. all unique locations)
-- in each subquery and then use the subqueries as the filters to identify qualified pid.
-- this approach translates the requirement directly to query without any functions or calculations.

-- we can start by identifying tiv_2015 that have duplicates in all records
-- SELECT
--     tiv_2015
-- FROM
--     Insurance
-- GROUP BY
--     tiv_2015
-- HAVING
--     COUNT(DISTINCT pid) > 1
-- | tiv_2015 |
-- | -------- |
-- | 10       |

-- next, we can identify the unique location(pair of lat and lon) from all locations.
-- since lat and lon are stored in two separate columns, we can combine these two columns
-- using the function CONCAT() for easier calculation.
-- SELECT
--     CONCAT(lat, lon) AS lat_lon
-- FROM
--     Insurance
-- GROUP BY
--     CONCAT(lat, lon)
-- HAVING
--     COUNT(DISTINCT pid) = 1
-- | lat_lon  |
-- | -------- |
-- | 1010     |
-- | 4040     |

-- with the qualified values from both conditions, we can now select the policyholders(pid)
-- that possess both values from these two subqueries in the main query.
-- this approach uses JOIN to find the matching records, but you can also use other functions
-- such as IN or NOT IN instead. to get the final output, we also want to calculate the sum of tiv_2016
-- from the qualified record, round the results to two decimal places,
-- and rename the column as requested in the main query.

SELECT
    ROUND(SUM(tiv_2016), 2) AS tiv_2016
FROM
    Insurance i
JOIN (
    SELECT
        tiv_2015
    FROM
        Insurance
    GROUP BY
        tiv_2015
    HAVING
        COUNT(DISTINCT pid) > 1
) t0
ON
    i.tiv_2015 = t0.tiv_2015
JOIN (
    SELECT
        CONCAT(lat, lon) AS lat_lon
    FROM
        Insurance
    GROUP BY
        CONCAT(lat, lon)
    HAVING
        COUNT(CONCAT(lat, lon)) = 1
) t1
ON
    CONCAT(i.lat, i.lon) = t1.lat_lon;

-- creating filters using window function
-- solution2
-- this approach calculates, for each policyholder(pid), how many times their values of
-- tiv_2015 and location show up in all records.

-- to calculate the aggregate counts for each record, we can leverage the window function,
-- pass the level that we need to calculate by and save the results in separate columns.
-- these two columns will be used later as filters to decide whether we want to keep the record.
-- SELECT
--     *,
--     COUNT(*) OVER (
--         PARTITION BY
--             tiv_2015
--     ) AS tiv_2015_cnt,
--     COUNT(*) OVER (
--         PARTITION BY
--             CONCAT(lat, lon)
--     ) AS loc_cnt
-- FROM
--     Insurance

-- below is the output from this step. each pid now has the number of
-- how many times their value of tiv_2015 and location are shared by pids in the table.
-- | pid | tiv_2015 | tiv_2016 | lat | lon | tiv_2015_cnt | loc_cnt |
-- | --- | -------- | -------- | --- | --- | ------------ | ------- |
-- | 1   | 10       | 5        | 10  | 10  | 3            | 1       |
-- | 3   | 10       | 30       | 20  | 20  | 3            | 2       |
-- | 2   | 20       | 20       | 20  | 20  | 1            | 2       |
-- | 4   | 10       | 40       | 40  | 40  | 3            | 1       |

-- based on these two columns(tiv_2015_cnt and loc_cnt), we can apply the filter to keep the pids
-- that have the same tiv_2015 and are not located in the same city with another pid.
-- we can put the previous step in either a subquery or a CTE. In the main query, we can also calculate
-- the sum of tiv_2016, round the result to two decimal places, and rename the column as requested.

SELECT
    ROUND(SUM(tiv_2016), 2) AS tiv_2016
FROM (
    SELECT
        *,
        COUNT(*) OVER (
            PARTITION BY
                tiv_2015
        ) AS tiv_2015_cnt,
        COUNT(*) OVER (
            PARTITION BY
                CONCAT(lat, lon)
        ) AS loc_cnt
    FROM
        Insurance
) t
WHERE
    tiv_2015_cnt > 1 AND loc_cnt = 1;

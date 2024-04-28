-- calling the original table twice and calculate as two columns
-- solution1
-- for the solution, we calling the original table twice, once as the table that stores the start timestamp,
-- and once as the table that stores the end timestamp.
-- to create the table alias, we give the original table Activity two different names,
-- and filter each table by the activity_type.
-- we also make sure the two tables are joined on the machine_id and process_id,
-- so the output will have the start timestamp and end timestamp stored in two different columns for each machine and process.
SELECT
    a1.machine_id,
    ROUND(AVG(a2.timestamp - a1.timestamp), 3) AS processing_time
FROM
    Activity a1,
    Activity a2
WHERE
    a1.machine_id = a2.machine_id
AND
    a1.process_id = a2.process_id
AND
    a1.activity_type = 'start'
AND
    a2.activity_type = 'end'
GROUP BY
    machine_id;

-- transform values with case when and then calculate
-- solution1
-- if we set all the 'start' timestamp to its negative value, we can get the time difference by using SUM(),
-- since (-start) + end is equal to end - start, which is the time difference.
-- to do this, we use CASE WHEN to multiply all start timestamp by -1, so the aggregated total of time stamp
-- becomes the time to complete a process for each machine.
SELECT
    machine_id,
    ROUND(SUM(CASE WHEN activity_type = 'start' THEN timestamp * -1 ELSE timestamp END)
    / COUNT(DISTINCT process_id), 3) AS processing_time
FROM
    Activity
GROUP BY
    machine_id;

-- JOIN the same table and calculate the difference between start timestamp and end timestamp,
-- group the values by machine_id and round average value of differences.
-- solution3
SELECT
    a1.machine_id,
    ROUND(AVG(a2.timestamp - a1.timestamp), 3) AS processing_time
FROM
    Activity a1
JOIN
    Activity a2
ON
    a2.timestamp > a1.timestamp
AND
    a1.machine_id = a2.machine_id
AND
    a1.process_id = a2.process_id
GROUP BY
    a1.machine_id;

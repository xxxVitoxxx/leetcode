-- using flow control statement CASE
-- solution1
-- for students with odd id, the new id is id+1 after switch unless it is the last seat.
-- and for students with even id, the new id is id-1.
-- in order to known how many seats in total, we can use a subquery.
SELECT
    (CASE
        WHEN MOD(id, 2) != 0 AND counts != id THEN id + 1
        WHEN MOD(id, 2) != 0 AND counts = id THEN id
        ELSE id - 1
    END) AS id,
    student
FROM
    Seat,
    (SELECT
        COUNT(*) AS counts
    FROM
        Seat
    ) AS seat_counts
ORDER BY
    id;

-- using bit manipulation and COALESCE() function
-- solution2
-- SELECT id, (id+1)^1 - 1, student FROM Seat;
-- use bit manipulation expression (id+1)^1 -1 can calculate the new id after switch.
-- | id | (id+1)^1-1 | student |
-- |----|------------|---------|
-- | 1  | 2          | Abbot   |
-- | 2  | 1          | Doris   |
-- | 3  | 4          | Emerson |
-- | 4  | 3          | Green   |
-- | 5  | 6          | Jeames  |

-- we can make a temp table and join seat with this table like below.
-- SELECT
--     *
-- FROM
--     Seat s1
-- LEFT JOIN
--     Seat s2
-- ON
--     (s1.id+1)^1 - 1 = s2.id
-- ORDER BY
--     s1.id;
-- | id | student | id | student |
-- |----|---------|----|---------|
-- | 1  | Abbot   | 2  | Doris   |
-- | 2  | Doris   | 1  | Abbot   |
-- | 3  | Emerson | 4  | Green   |
-- | 4  | Green   | 3  | Emerson |
-- | 5  | Jeames  |null| null    |
SELECT
    s1.id,
    COALESCE(s2.student, s1.student) AS student
FROM
    Seat s1
LEFT JOIN
    Seat s2
ON
    ((s1.id+1)^1) - 1 = s2.id
ORDER BY
    id;

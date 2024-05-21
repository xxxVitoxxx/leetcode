-- note: in math, three segments can form a triangle only
-- if the sum of the any two segments is larger then the third one.

-- solution
SELECT
    x,
    y,
    z,
    CASE WHEN x+y > z AND x+z > y AND y+z > x THEN 'Yes'
    ELSE 'No' END AS triangle
    -- IF(x+y > z AND x+z > y AND y+z > x, 'Yes', 'No') AS triangle
FROM
    Triangle;

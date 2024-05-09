-- solution1
(SELECT
    name AS results
FROM
    Users
INNER JOIN
    MovieRating
USING(user_id)
GROUP BY
    user_id
ORDER BY
    COUNT(user_id) DESC,
    name
LIMIT 1)

UNION ALL

(SELECT
    title AS results
FROM
    Movies
INNER JOIN
    MovieRating
USING(movie_id)
WHERE
    DATE_FORMAT(created_at, '%Y-%m') = '2020-02'
GROUP BY
    title
ORDER BY
    AVG(rating) DESC,
    title
LIMIT 1)

-- solution2
SELECT
    DISTINCT FIRST_VALUE(u.name) OVER(
        ORDER BY
            COUNT(r.movie_id) DESC,
            u.name
    ) AS results
FROM
    Users u
LEFT JOIN
    MovieRating r
USING(user_id)
GROUP BY
    u.user_id
UNION ALL
SELECT
    DISTINCT FIRST_VALUE(m.title) OVER(
        ORDER BY
            AVG(r.rating) DESC,
            m.title
    ) AS results
FROM
    Movies m
JOIN
    MovieRating r
USING(movie_id)
WHERE
    DATE_FORMAT(r.created_at, '%Y-%m') = '2020-02'
GROUP BY
    m.movie_id;

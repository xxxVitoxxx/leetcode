SELECT
    tweet_id
FROM
    Tweets
WHERE
    CHAR_LENGTH(content) > 15;
-- the best function to use to count the number of characters used in a string is CHAR_LENGTH(str),
-- which returns the length of the string.

-- the other common function, LENGTH(str), also works for this question since the column count
-- contains only English characters and no special characters.
-- otherwise, LENGTH() might return a different result because this function returns the length of
-- the string in bytes and certain characters contain more than 1 byte.

-- using character '¥' as an example: CHAR_LENGTH() return 1 as the result,
-- while LENGTH() return 2 because this string contains 2 bytes.

-- use LIKE
-- solution1
SELECT
    patient_id,
    patient_name,
    conditions
FROM
    Patients
WHERE
    conditions LIKE 'DIAB1%' OR conditions LIKE '% DIAB1%';

-- use regular expression word boundaries
-- solution2

-- the REGEXP operator in SQL is used to match text patterns. by using the pattern "\\bDIAB1.*" ,
-- we specify that we are looking for strings where "DIAB1" appears right after a word boundary(
-- indicating the start of a word) and is following by any sequence of characters.
-- this ensures that we only match condition codes that precisely indicate Type DIAB1.
SELECT
    patient_id,
    patient_name,
    conditions
FROM
    Patients
WHERE
    conditions REGEXP '\\bDIAB1.*';

SELECT
    (
        SELECT DISTINCT
            CASE
                WHEN e.salary::text ~ '^[0-9]+(\.[0-9]+)?$' THEN e.salary::numeric
                ELSE NULL
            END
        FROM Employee e
        ORDER BY 1 DESC
        LIMIT 1
        OFFSET 1
    ) AS "SecondHighestSalary";
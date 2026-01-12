SELECT DISTINCT num as ConsecutiveNums
FROM (
    SELECT 
        num,
        LAG(num, 1) OVER (ORDER BY id) AS prev, -- get 'num' from previous row with 1 offset
        LAG(num, 2) OVER (ORDER BY id) AS prev2 -- get 'num' from previous row with 2 offset
    FROM Logs
) t
WHERE num = prev AND num = prev2; -- filter numbers with equal values
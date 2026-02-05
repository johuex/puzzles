CREATE OR REPLACE FUNCTION NthHighestSalary(N INT) 
RETURNS TABLE (Salary INT) AS $$
BEGIN
  RETURN QUERY (

    -- my query start
    WITH ranked_salaries AS (
        SELECT DISTINCT e.salary,
               DENSE_RANK() OVER (ORDER BY e.salary DESC) as rank_num
        FROM Employee e
    )
    SELECT 
        CASE 
            WHEN (SELECT COUNT(DISTINCT e1.salary) FROM Employee e1) < N 
            THEN NULL
            ELSE (SELECT rs.salary 
                  FROM ranked_salaries rs
                  WHERE rs.rank_num = N)
        END as salary
    -- my query end
  );
END;
$$ LANGUAGE plpgsql;
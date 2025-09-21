-- my own solution
select 
    distinct c.visited_on,
    round(sum(c.amount) over(order by c.visited_on range between interval '6 day' preceding and current row), 2) as amount,
    round(sum(c.amount) over(order by c.visited_on range between interval '6 day' preceding and current row) / 7.0, 2) as average_amount
from Customer c
order by c.visited_on
offset 6 rows;

-- another solution
WITH CTE AS(
SELECT visited_on, SUM(amount) as total_amount
FROM Customer
GROUP BY visited_on
ORDER BY visited_on)

SELECT visited_on, SUM(total_amount) OVER (ROWS BETWEEN 6 PRECEDING AND CURRENT ROW) as amount, ROUND(avg(total_amount) OVER (ROWS BETWEEN 6 PRECEDING AND CURRENT ROW),2) as average_amount
FROM CTE
OFFSET 6
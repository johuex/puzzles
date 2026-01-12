-- my own solution
select s.product_id, s.year as first_year, s.quantity, s.price
from sales s
where (s.product_id, s.year) in (
	select s.product_id, min(s.year)
	from sales s
	group by s.product_id
);

-- another solution from top
WITH p_1st_yr AS (
    SELECT 
        *,
        min(year) OVER(PARTITION BY product_id) AS "first_year"
    FROM Sales
)

SELECT 
    product_id,
    first_year,
    quantity, 
    price
FROM p_1st_yr
WHERE first_year = year
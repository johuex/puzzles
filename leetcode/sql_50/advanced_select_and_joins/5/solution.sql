-- my solution
-- Write your PostgreSQL query statement below
with unique_product as (
    select distinct p.product_id
    from Products p
),
max_change_day as (
     SELECT p.product_id,
           MAX(p.change_date)
    FROM products p
    WHERE p.change_date <= '2019-08-16'
    group by p.product_id
)

select 
    up.product_id,
    case
        when mcd.product_id is null then 10
        else p.new_price
    end as price
from unique_product up
left join max_change_day mcd on up.product_id = mcd.product_id
left join products p on p.product_id = mcd.product_id and mcd.max = p.change_date


-- more accurate solution ( a little bit fast by leetcode measure)
with max_change_day as (
     SELECT p.product_id,
           MAX(p.change_date) filter (WHERE p.change_date <= '2019-08-16') as date
    FROM products p
    group by p.product_id
)

SELECT 
		pd.product_id,
		COALESCE(p.new_price, 10) AS price
FROM max_change_day pd
LEFT JOIN products p
ON p.product_id = pd.product_id AND pd.date = p.change_date
ORDER BY pd.product_id
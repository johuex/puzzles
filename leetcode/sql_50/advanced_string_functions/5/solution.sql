-- my solution
-- Write your PostgreSQL query statement below
select a.sell_date, count(distinct a.product) as num_sold, array_to_string(array_agg(distinct a.product), ',') AS products
from Activities a
group by a.sell_date
order by 1;

-- my another solution
select a.sell_date, count(distinct a.product) as num_sold, string_agg(distinct a.product, ',') AS products
from Activities a
group by a.sell_date
order by 1;

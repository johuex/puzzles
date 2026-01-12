-- my solution
with cte as (
    select p.product_id, sum(o.unit)
    from Products p
    join Orders o on p.product_id = o.product_id
    where o.order_date between '2020-02-01' and '2020-02-29'
    group by p.product_id
    having sum(o.unit) >= 100
)

select p.product_name, c.sum as unit
from Products p
join cte c on c.product_id = p.product_id

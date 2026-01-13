-- my solution
select u.user_id "buyer_id", any_value(u.join_date) "join_date", count(o.order_id) filter (where extract(year from o.order_date) = 2019) "orders_in_2019"
from Users u
left join Orders o on o.buyer_id = u.user_id
group by u.user_id;


-- fastest solution
select u.user_id "buyer_id", any_value(u.join_date) "join_date", count(o.order_id) "orders_in_2019" -- any_value is faster than min + don't spend time on extraction year in FILTER
from Users u
-- we filter on join, less rows will be scanned later
left join Orders o on o.buyer_id = u.user_id and o.order_date between '2019-01-01' and '2019-12-31'
group by 1;
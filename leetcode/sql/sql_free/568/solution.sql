select o.customer_number
from Orders o
group by o.customer_number
order by count(*) desc
limit 1;
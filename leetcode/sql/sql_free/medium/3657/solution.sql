-- my solution
select ct.customer_id
from customer_transactions ct
group by ct.customer_id
having count(ct.transaction_id) >= 3 and
    max(ct.transaction_date)::date - min(ct.transaction_date)::date >= 30 and
    (count(ct.transaction_id) filter (where ct.transaction_type = 'refund'))::numeric / count(ct.transaction_id) < 0.2
order by 1 asc;
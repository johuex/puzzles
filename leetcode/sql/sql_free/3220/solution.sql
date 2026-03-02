select
    transaction_date,
    coalesce (sum(amount) filter (where amount % 2 = 1), 0) "odd_sum",
    coalesce (sum(amount) filter (where amount % 2 = 0), 0) "even_sum"
from transactions
group by transaction_date
order by 1 asc;
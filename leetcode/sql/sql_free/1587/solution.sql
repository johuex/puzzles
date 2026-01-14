-- my solution
select any_value(u.name) "name", sum(t.amount) "balance"
from Users u
left join Transactions t on u.account = t.account
group by u.account
having sum(t.amount) > 10000;



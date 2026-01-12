-- my own solution
select max(num) as num
from (
    select num
    from MyNumbers n
    group by num
    having count(num) = 1
)
-- my own solution
-- Write your PostgreSQL query statement below
select agr.id, sum(agr.num) as num 
from (
    select ra1.requester_id as id, count(ra1.requester_id) as num
    from RequestAccepted ra1
    group by ra1.requester_id

    union all

    select ra2.accepter_id as id, count(ra2.accepter_id) as num
    from RequestAccepted ra2
    group by ra2.accepter_id
) as agr
group by 1
order by 2 DESC
limit 1;

-- another solution (mor cost, but more fast in actual)
select ids.id, count(*) as num
from (
    select requester_id as id
    from RequestAccepted
    union all
    select accepter_id as id
    from RequestAccepted
) as ids
group by ids.id
order by 2 desc
limit 1;
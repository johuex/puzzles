-- my solution
select distinct a.*
from stadium as a, stadium as b, stadium as c
where a.people >= 100 and b.people >= 100 and c.people >= 100
    and
    (
        (a.id - b.id = 1 AND b.id - c.id = 1) or
        (c.id - b.id = 1 AND b.id - a.id = 1) or
        (b.id - a.id = 1 AND a.id - c.id = 1)
    )
order by visit_date;

-- faster solution
with s1 as (
    -- mark days with more or equal 100 visits as sequence of equal number
    select *, id - row_number() over(order by id) as grp
    from Stadium
    where people >= 100
),
s2 as (
    -- found groups where more or equal 3 items
    select grp
    from s1
    group by grp
    having count(*) >= 3
)

select id, visit_date, people
from s1
where grp in (
    select grp from s2
);
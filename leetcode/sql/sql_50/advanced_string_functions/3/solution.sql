-- my ugly solution
delete
from Person p
where p.id not in (
    select a.id
    from (
        select p.id, rank() over (partition by p.email order by p.id asc)
        from Person p
    ) as a
    where a.rank = 1
);


-- another solution
delete
from person p1
using person p2
where p1.email = p2.email and p1.id > p2.id
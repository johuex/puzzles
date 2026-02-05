-- my solution
with banned as (
    select u.users_id
    from Users u
    where u.banned = 'Yes'
)


select t.request_at "Day", round(count(t.id) filter (where t.status LIKE 'cancelled%')::numeric / count(t.id), 2) "Cancellation Rate"
from Trips t
where
    t.request_at >= '2013-10-01' and t.request_at <= '2013-10-03'
    and t.client_id not in (select * from banned) and t.driver_id not in (select * from banned)
group by t.request_at;


-- another solution
select t.request_at as Day, 
    ROUND(
        SUM(
            CASE
                WHEN t.status in ('cancelled_by_client', 'cancelled_by_driver') then 1 else 0 end)*1.0 / count(*) , 2)  AS "Cancellation rate" 
from Trips  t 
join Users u on t.client_id = u.users_id and u.banned = 'No'
join Users v on t.driver_id = v.users_id and v.banned = 'No'
where t.request_at  between '2013-10-01' and '2013-10-03'
group by t.request_at ;

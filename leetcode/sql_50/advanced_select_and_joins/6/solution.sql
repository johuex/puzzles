-- my solution
with queue_with_sum as (
    select
        q.turn,
        q.person_id,
        q.person_name,
        q.weight,
        sum(q.weight) over (order by q.turn) as running_total
    from queue q
),
last_pass as (
    select max(turn) as "turn"
    from queue_with_sum
    where running_total <= 1000
)
select q.person_name 
from Queue q
join last_pass lp on q.turn = lp.turn


-- another solution
WITH  Prefix_sum AS (
    SELECT person_id, person_name, 
    SUM(weight) OVER (ORDER BY turn ASC) AS presum
    FROM Queue
) 

SELECT person_name 
FROM Prefix_sum
WHERE presum <= 1000 
ORDER BY presum DESC
LIMIT 1;